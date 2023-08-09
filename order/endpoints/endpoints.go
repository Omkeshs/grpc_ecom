package endpoints

import (
	"context"

	"github.com/Omkeshs/grpc_ecom/order/bl"
	"github.com/Omkeshs/grpc_ecom/order/pkg/middleware"
	"github.com/Omkeshs/grpc_ecom/order/spec"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

func NewOrderEndpoint(logger *zap.Logger, bl bl.OrderBL) spec.EndPoints {
	var placeOrderEP endpoint.Endpoint
	{
		placeOrderEP = makeplaceOrderEP(bl)
		placeOrderEP = middleware.AddEndpointMiddleware(logger, placeOrderEP, "placeOrderEP", 0)
	}

	return spec.EndPoints{
		PlaceOrder: placeOrderEP,
	}
}

func makeplaceOrderEP(svc bl.OrderBL) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		placeOrderRequest := request.(spec.PlaceOrderRequest)
		return svc.PlaceOrder(ctx, placeOrderRequest)
	}
}
