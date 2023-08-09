package spec

import "github.com/go-kit/kit/endpoint"

type EndPoints struct {
	PlaceOrder endpoint.Endpoint
	ListOrder  endpoint.Endpoint
}
