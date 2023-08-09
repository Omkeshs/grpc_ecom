package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Omkeshs/grpc_ecom/order/bl"
	"github.com/Omkeshs/grpc_ecom/order/db"
	"github.com/Omkeshs/grpc_ecom/order/dl"
	"github.com/Omkeshs/grpc_ecom/order/endpoints"
	"github.com/Omkeshs/grpc_ecom/order/handler"
	logger "github.com/Omkeshs/grpc_ecom/order/pkg/log"

	// "golang.org/x/vuln/client"

	"github.com/Omkeshs/grpc_ecom/product/grpc/product/pb"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

func main() {
	// logger
	logger := logger.FileLogger("order.log")

	//env
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Log(zapcore.ErrorLevel, "failed to read config file")
		os.Exit(1)
	}

	// DB initialization
	dbConn, _ := db.InitDB(logger)
	defer dbConn.Close()

	// init router
	router := mux.NewRouter()

	// init repo/dl
	orderDL := dl.NewDL(logger, dbConn)

	// grpc.InitClientServer()
	// grpc.InitClientServer()

	// client.InitClientServer()
	pb.NewProductServiceClient(clientConnection)

	// initbl
	ordersvc := bl.NewBL(logger, orderDL)

	// init endpoints
	orderEPs := endpoints.NewOrderEndpoint(logger, ordersvc)

	// inithandler
	handler.Inithandler(logger, router, orderEPs)

	// listen and serve
	restPort := viper.GetString("ORDER_SVC_REST_PORT")
	// grpcPort := viper.GetString("GRPC_PORT")

	errs := make(chan error, 1)
	go func() {
		errs <- http.ListenAndServe(":"+restPort, router)
	}()

	// go func() {
	// 	errs <- grpc.InitGRPCServer(logger, productsvc, grpcPort)
	// }()

	fmt.Println(<-errs)
}
