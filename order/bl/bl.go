package bl

import (
	"context"

	"github.com/Omkeshs/grpc_ecom/order/spec"

	"go.uber.org/zap"
)

type OrderDL interface {
}

type bl struct {
	logger *zap.Logger
	dl     OrderDL
}

type OrderBL interface {
	PlaceOrder(ctx context.Context, req spec.PlaceOrderRequest) (spec.Order, error)
	// ListOrder()
}

func NewBL(log *zap.Logger, dl OrderDL) *bl {
	return &bl{
		logger: log,
		dl:     dl,
	}
}

func (svc bl) PlaceOrder(ctx context.Context, req spec.PlaceOrderRequest) (spec.Order, error) {
	svc.logger.Sugar().Info(" \n \n <b>Place Order BL</b> \n")
	return spec.Order{
		ID:       1,
		Quantity: 1,
		Status:   "test",
	}, nil
}
