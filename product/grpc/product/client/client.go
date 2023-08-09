package client

import (
	"github.com/Omkeshs/grpc_ecom/product/grpc/product/pb"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// InitClientServer ...
func InitClientServer(log *zap.Logger, port string) pb.ProductServiceClient {
	clientConnection, err := grpc.Dial("localhost:"+port, grpc.WithInsecure()) // WithInSecure()
	if err != nil {
		log.Sugar().Panicf("unable to init product grpc-client err :%s", err.Error())
	}

	defer clientConnection.Close()
	conn := pb.NewProductServiceClient(clientConnection)
	return conn
}
