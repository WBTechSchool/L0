package models

type OrderRequest struct {
	CustomerID        int64  `json:"customer_id"`
	Locale            string `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	DeliveryService   string `json:"delivery_service"`
	TrackNumber       string `json:"track_number"`
	Entry             string `json:"entry"`
	ShardKey          string `json:"shard_key"`
	OffShard          string `json:"off_shard"`
	SmID              int64  `json:"sm_id"`
}

type CreateOrderRequest struct {
	Order      OrderRequest
	Delivery   Delivery
	Payment    Payment
	OrderItems []OrderItem
}

type OrderResponse struct {
	OrderID           int64  `json:"order_id"`
	CustomerID        int64  `json:"customer_id"`
	Locale            string `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	DeliveryService   string `json:"delivery_service"`
	TrackNumber       string `json:"track_number"`
	Entry             string `json:"entry"`
	ShardKey          string `json:"shard_key"`
	OffShard          string `json:"off_shard"`
	SmID              int64  `json:"sm_id"`

	Delivery Delivery    `json:"delivery"`
	Payments Payment     `json:"payments"`
	Items    []OrderItem `json:"items"`
}
