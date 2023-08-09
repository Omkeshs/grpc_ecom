package endpoints

import (
	"context"

	bl "github.com/Omkeshs/grpc_ecom/product/bl"
	"github.com/Omkeshs/grpc_ecom/product/pkg/middleware"
	spec "github.com/Omkeshs/grpc_ecom/product/spec"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

func NewProductEndpoint(logger *zap.Logger, bl bl.ProductBL) spec.EndPoints {
	var listEP endpoint.Endpoint
	{
		listEP = makeListEP(bl)
		listEP = middleware.AddEndpointMiddleware(logger, listEP, "listEP", 0)
	}

	var updateEP endpoint.Endpoint
	{
		updateEP = makeUpdateEP(bl)
		updateEP = middleware.AddEndpointMiddleware(logger, updateEP, "updateEP", 0)
	}

	return spec.EndPoints{
		ListProductEP:   listEP,
		UpdateProductEP: updateEP,
	}
}

func makeListEP(svc bl.ProductBL) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		loginRequest := request.(spec.ProductRequest)
		return svc.ListProduct(ctx, loginRequest)
	}
}

func makeUpdateEP(svc bl.ProductBL) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		updateRequest := request.(spec.UpdateProductRequest)
		return nil, svc.UpdateProduct(ctx, updateRequest)
	}
}
