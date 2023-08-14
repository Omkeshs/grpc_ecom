package spec

type ProductResponse map[int32]Product

type ProductRequest struct {
	IDs []int32 `json:"ids"`
}

type UpdateProduct struct {
	ID       int32 `json:"id"`
	Quantity int32 `json:"quantity"`
}

type UpdateProductRequest []UpdateProduct
