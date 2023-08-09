package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Omkeshs/grpc_ecom/order/spec"

	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func Inithandler(logger *zap.Logger, router *mux.Router, eps spec.EndPoints) {
	router.Methods(http.MethodPost).Path(spec.PlaceOrderPath).Handler(kitHttp.NewServer(
		eps.PlaceOrder,
		decodePlaceOrderRequest,
		JSONEncodeAPIResponse,
	))
}

// JSONEncodeAPIResponse ...
func JSONEncodeAPIResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}
