package bl

import (
	"context"
	"errors"

	spec "github.com/Omkeshs/grpc_ecom/product/spec"

	"go.uber.org/zap"
)

type ProductDL interface {
}

type bl struct {
	logger *zap.Logger
	dl     ProductDL
}

// ProdMap - Product Map ...
var ProdMap = &spec.ProductResponse{
	1: {
		ID:       1,
		Name:     "IPhone12",
		Price:    50000,
		Category: 1,
		Quantity: 100,
	},
	2: {
		ID:       2,
		Name:     "IPhone13",
		Price:    60000,
		Category: 1,
		Quantity: 200,
	},
	3: {
		ID:       3,
		Name:     "IPhone14",
		Price:    70000,
		Category: 1,
		Quantity: 200,
	},
	101: {
		ID:       101,
		Name:     "IPad",
		Price:    30000,
		Category: 2,
		Quantity: 50,
	},
	102: {
		ID:       102,
		Name:     "IPad 3rd gen",
		Price:    330000,
		Category: 2,
		Quantity: 50,
	},
	201: {
		ID:       201,
		Name:     "Mobile case",
		Price:    100,
		Category: 3,
		Quantity: 50,
	},
}

type ProductBL interface {
	ListProduct(ctx context.Context, req spec.ProductRequest) (*spec.ProductResponse, error)
	UpdateProduct(ctx context.Context, req spec.UpdateProductRequest) error
}

func NewBL(log *zap.Logger, dl ProductDL) *bl {
	return &bl{
		logger: log,
		dl:     dl,
	}
}

func (svc *bl) ListProduct(ctx context.Context, req spec.ProductRequest) (*spec.ProductResponse, error) {
	svc.logger.Sugar().Info("layer", "bl")
	return ProdMap, nil
}

func (svc *bl) UpdateProduct(ctx context.Context, req spec.UpdateProductRequest) error {
	products := *ProdMap
	for _, requestProduct := range req {
		if _, ok := products[requestProduct.ID]; !ok {
			//check product is exist
			return errors.New("product not fount")
		}
		product := products[requestProduct.ID]
		product.Quantity = requestProduct.Quantity
		products[requestProduct.ID] = product
	}

	ProdMap = &products

	return nil
}
