package storage

import (
	models "cars_with_sql/api/model"
	"context"
	"time"
)

type IStorage interface {
	CloseDB()
	Car() ICarstorage
	Customer() ICustomerStorage
	Order() Iorderstorage
}

type ICarstorage interface {
	GetAllCarS(context.Context, models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
	UpdateCar(context.Context, models.UpdateCarRequest) (string, error)
	Deletecar(context.Context, string) (string, error)
	Create(context.Context, models.Car) (string, error)
	GetByidcar(context.Context, string) (models.Car, error)
}

type ICustomerStorage interface {
	CreateCus(context.Context, models.Customers) (string, error)
	GetAllCustomers(context.Context, models.GetAllCustomerRequest) (models.GetAllCustomersResponse, error)
	UpdateCustomer(context.Context, models.Customers) (string, error)
	DeleteCustomer(context.Context, string) (string, error)
	GetByIDCustomer(context.Context, string) (models.Customers, error)
	GetAllCustomerCars(context.Context, models.GetAllCustomerCarsRequest) (models.GetAllCustomerCarsResponse, error)
	Login(context.Context, models.Changepasswor) (string, error)
	GetPassword(ctx context.Context, phone string) (string, error)
	ChangePassword(ctx context.Context, pass models.Changepasswor) (string, error)
	GetByLogin(context.Context, string) (models.Customers, error)
	Checklogin(context.Context, string) (models.CustomerRegisterRequest, error)
}
type Iorderstorage interface {
	Create(context.Context, models.CreateOrder) (string, error)
	GetAll(context.Context, models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error)
	GetByID(context.Context, string) (models.GetOrder, error)
	UpdateOrders(context.Context, models.CreateOrder) (string, error)

	//DeleteOrder(string) error

}

type IREdisStorage interface {
	SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Get(ctx context.Context, key string) interface{}
	Del(ctx context.Context, key string) error
}
