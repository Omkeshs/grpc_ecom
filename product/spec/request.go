package spec

type ProductResponse map[int32]Product

type ProductRequest struct {
	IDs []int32 `json:"ids"`
}

type UpdateProductRequest struct {
	ID       int32
	Quantity int32 `json:"quantity"`
}
