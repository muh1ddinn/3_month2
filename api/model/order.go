package models

type GetOrder struct {
	Id         string    `json:"id"`
	Customers  Customers `json:"cudtomer"`
	Car        Car       `json:"car"`
	FromDate   string    `json:"from_date"`
	ToDate     string    `json:"to_date"`
	Status     string    `json:"status"`
	Paid       bool      `json:"payment_status"`
	Created_at string    `json:"created_at"`
	Updated_at string    `json:"updated_at"`
	Amount     int       `json:"amount"`
}

type CreateOrder struct {
	Id         string  `json:"id"`
	CustomerId string  `json:"customer_id"`
	CarId      string  `json:"cars_id"`
	FromDate   string  `json:"from_date"`
	ToDate     string  `json:"to_date"`
	Status     string  `json:"status"`
	Paid       bool    `json:"payment_status"`
	Amount     float64 `json:"amount"`
}

type GetAllOrdersResponse struct {
	Orders []GetOrder `json:"orders"`
	Count  int64      `json:"count of orders"`
}

type GetAllOrdersRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}


