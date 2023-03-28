package models

type Order struct {
	Id         string         `json:"id"`
	Name       string         `json:"name"`
	Price      float64        `json:"product_price"`
	TotalPrice float64        `json:"total_price"`
	Quantity   int32          `json:"quantity"`
	User       ReturnUser     `json:"user"`
	Courier    ReturnCourier  `json:"courier"`
	Customer   ReturnCustomer `json:"customer"`
	Product    ReturnProduct  `json:"product"`
	CreatedAt  string         `json:"created_at"`
	UpdatedAt  string         `json:"updated_at"`
}

type OrderPrimaryKey struct {
	Id string `json:"id"`
}

type CreateOrder struct {
	Name       string `json:"name"`
	Quantity   int32  `json:"quantity"`
	UserId     string `json:"user_id"`
	CustomerId string `json:"customer_id"`
	ProductId  string `json:"product_id"`
	CourierId  string `json:"courier_id"`
}

type UpdateOrder struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Quantity   int32  `json:"quantity"`
	UserId     string `json:"user_id"`
	CustomerId string `json:"customer_id"`
	ProductId  string `json:"product_id"`
	CourierId  string `json:"courier_id"`
}

type GetListOrderRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListOrderResponse struct {
	Count  int      `json:"count"`
	Orders []*Order `json:"orders"`
}
