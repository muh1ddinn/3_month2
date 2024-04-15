package models

type Customers struct {
	Id         string     `json:"id"`
	First_name string     `json:"first_name"`
	Last_name  string     `json:"Last_name"`
	Gmail      string     `json:"gmail"`
	Phone      string     `json:"phone"`
	Is_blocked bool       `json:"is_blocked"`
	Created_at string     `json:"created_at"`
	Updated_at string     `json:"updated_at"`
	Deleted_at int        `json:"deleted_at"`
	Order      []GetOrder `json:"order"`
	Password   string     `json:"password"`
	Login      string     `json:"login"`
}

type GetAllCustomersResponse struct {
	Customers []Customers `json:"customers"`
	Count     int64       `json:"count"`
}

type GetAllCustomerRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllCustomer struct {
	Id         string   `json:"id"`
	First_name string   `json:"first_name"`
	Last_name  string   `json:"last_name"`
	Gmail      string   `json:"gmail"`
	Phone      string   `json:"phone"`
	Is_blocked bool     `json:"isblocked"`
	Created_at string   `json:"createdAt"`
	Updated_at string   `json:"updatedAt"`
	OrderCount int      `json:"ordercount"`
	CarsCount  int      `json:"carscount"`
	Order      GetOrder `json:"order"`
	Car        Car      `json:"car"`
}

type GetAllCustomerCars struct {
	CustomerID string  `json:"id"`
	CarName    string  `json:"name"`
	CreatedAt  string  `json:"creatAt"`
	Price      float32 `json:"amount"`
}

type GetAllCustomerCarsRequest struct {
	Search string `json:"search"`
	Id     string `json:"id"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllCustomerCarsResponse struct {
	Customer []GetAllCustomerCars `json:"orders"`
	Count    int16                `json:"count"`
}

type Changepasswor struct {
	Id          string `json:"id"`
	Phone       string `json:"phone"`
	OldPassword string `json:"password"`
	Login       string `json:"login"`
	NewPassword string `json:"newpassword"`
}

type Checklogin struct {
	Id    string `json:"id"`
	Phone string `json:"phone"`
	Gmail string `json:"gmail"`
}
