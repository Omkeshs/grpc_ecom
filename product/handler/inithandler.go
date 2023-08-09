package handler

import (
	"context"
	"encoding/json"
	"net/http"

	spec "github.com/Omkeshs/grpc_ecom/product/spec"

	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Inithandler ...
func Inithandler(logger *zap.Logger, router *mux.Router, eps spec.EndPoints) {

	router.Methods(http.MethodGet).Path(spec.ListProductPath).Handler(kitHttp.NewServer(
		eps.ListProductEP,
		decodeListRequest,
		JSONEncodeAPIResponse,
	))

	router.Methods(http.MethodPut).Path(spec.UpdateProductPath).Handler(kitHttp.NewServer(
		eps.UpdateProductEP,
		decodeUpdateRequest,
		JSONEncodeAPIResponse,
	))
}

// JSONEncodeAPIResponse ...
func JSONEncodeAPIResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}
