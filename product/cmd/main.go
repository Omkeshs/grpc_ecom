package main

import (
	"fmt"
	"net/http"
	"os"

	bl "github.com/Omkeshs/grpc_ecom/product/bl"
	db "github.com/Omkeshs/grpc_ecom/product/db"
	dl "github.com/Omkeshs/grpc_ecom/product/dl"
	endpoints "github.com/Omkeshs/grpc_ecom/product/endpoints"
	grpc "github.com/Omkeshs/grpc_ecom/product/grpc"
	handler "github.com/Omkeshs/grpc_ecom/product/handler"
	log "github.com/Omkeshs/grpc_ecom/product/pkg/log"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

func main() {

	// logger
	logger := log.FileLogger("product.log")

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

	//init repo/dl
	productDL := dl.NewDL(logger, dbConn)

	//initbl
	productsvc := bl.NewBL(logger, productDL)

	//init endpoints
	productEPs := endpoints.NewProductEndpoint(logger, productsvc)

	//inithandler
	handler.Inithandler(logger, router, productEPs)

	//listen and serve
	restPort := viper.GetString("PRODUCT_REST_PORT")
	grpcServerPort := viper.GetString("GRPC_SERVER_PORT")

	errs := make(chan error, 2)
	go func() {
		errs <- http.ListenAndServe(":"+restPort, router)
	}()

	go func() {
		errs <- grpc.InitGRPCServer(logger, productsvc, grpcServerPort)
	}()

	fmt.Println(<-errs)

}
