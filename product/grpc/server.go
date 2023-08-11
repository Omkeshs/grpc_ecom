package grpc

import (
	"context"
	"fmt"
	"net"

	product "github.com/Omkeshs/grpc_ecom/grpc/product"
	// "https://github.com/Omkeshs/grpc_ecom/tree/main/grpc/product"
	bl "github.com/Omkeshs/grpc_ecom/product/bl"
	spec "github.com/Omkeshs/grpc_ecom/product/spec"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	logger    *zap.Logger
	productbl bl.ProductBL
	product.UnimplementedProductServiceServer
}

func NewProductServer(log *zap.Logger, svc bl.ProductBL) *server {
	return &server{
		logger:    log,
		productbl: svc,
	}
}

func InitGRPCServer(log *zap.Logger, svc bl.ProductBL, port string) error {
	listen, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Sugar().Error("internal error in listen grpc sever", err.Error())
		return err
	}

	productServer := NewProductServer(log, svc)

	s := grpc.NewServer()

	product.RegisterProductServiceServer(s, productServer)

	reflection.Register(s)

	log.Log(zapcore.InfoLevel, "Successfully initialize grpc server")

	errs := make(chan error)
	go func() {
		errs <- s.Serve(listen)
	}()
	return <-errs
}

func (svc *server) ListProduct(ctx context.Context, in *pb.ListRequest) (*pb.ListProductResponse, error) {

	fmt.Println(" -- ListProduct")
	req := spec.ProductRequest{}
	products, err := svc.productbl.ListProduct(ctx, req)
	if err != nil {
		svc.logger.Sugar().Debug("layer", "grpc-server", "method", "listproduct", "err", err.Error())
		return nil, err
	}

	resp := pb.ListProductResponse{
		Products: make(map[int32]*pb.Product),
	}

	for _, product := range *products {
		p := pb.Product{
			Id:       product.ID,
			Quantity: product.Quantity,
			Name:     product.Name,
			Price:    int32(product.Price),
			Category: int32(product.Category),
		}
		resp.Products[product.ID] = &p
	}

	return &resp, nil
}
