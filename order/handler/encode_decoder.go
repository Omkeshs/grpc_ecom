package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Omkeshs/grpc_ecom/order/spec"
)

func decodePlaceOrderRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	request := spec.PlaceOrderRequest{}

	var err error
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&request)
	if err != nil {
		return nil, fmt.Errorf("invalid argument")
	}
	return request, nil

}
