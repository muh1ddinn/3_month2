package service

import (
	"cars_with_sql/pkg/logger"
	"cars_with_sql/storage"
)

type IServiceMangaer interface {
	Car() carService
	Customer() customerSer
	Order() orderSer
	Auth() authService
}

type Service struct {
	carService  carService
	customerSer customerSer
	orderSer    orderSer
	authService authService
	logger      logger.ILogger
}

func New(storage storage.IStorage, log logger.ILogger) Service {
	return Service{
		carService:  NewCarService(storage, log),
		customerSer: NewCustomerSer(storage, log),
		orderSer:    NewOrderservice(storage, log),
		authService: NewAuthService(storage, log),
		logger:      log,
	}
}

func (s Service) Car() carService {
	return s.carService
}

func (s Service) Customer() customerSer {
	return s.customerSer
}

func (s Service) Order() orderSer {
	return s.orderSer

}

func (s Service) Auth() authService {
	return s.authService

}
