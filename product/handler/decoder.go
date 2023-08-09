package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	spec "github.com/Omkeshs/grpc_ecom/product/spec"
	utils "github.com/Omkeshs/grpc_ecom/product/utils"
)

func decodeListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	request := spec.ProductRequest{}

	var err error
	request.IDs, err = utils.GetQueryIDs(r, "ids")
	if err != nil {
		return nil, fmt.Errorf("invalid ids in query param : %s", err.Error())
	}

	return request, nil
}

func decodeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := spec.UpdateProductRequest{}
	var err error
	req.ID, err = utils.GetPathID(r, "productID")
	if err != nil {
		return nil, fmt.Errorf("invalid productID: %s", err.Error())
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("failed to decode body :%s", err.Error())
	}

	return req, nil
}
