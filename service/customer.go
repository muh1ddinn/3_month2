package service

import (
	models "cars_with_sql/api/model"
	"cars_with_sql/pkg/logger"
	"cars_with_sql/pkg/password"
	"cars_with_sql/storage"
	"context"
	"fmt"
)

type customerSer struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewCustomerSer(storage storage.IStorage, logger logger.ILogger) customerSer {

	return customerSer{
		storage: storage,
		logger:  logger,
	}
}
func (u customerSer) CreateCus(ctx context.Context, customer models.Customers) (string, error) {

	pKey, err := u.storage.Customer().CreateCus(ctx, customer)
	if err != nil {
		u.logger.Error("ERROR in service layer while creating customer", logger.Error(err))
		return "", err
	}

	return pKey, nil
}

func (s customerSer) GetByID(ctx context.Context, id string) (models.Customers, error) {
	customer, err := s.storage.Customer().GetByIDCustomer(ctx, id)
	if err != nil {
		s.logger.Error("failed to get customer by ID", logger.Error(err))
		return customer, err
	}
	
	return customer, nil
}

func (s customerSer) GetAllCus(ctx context.Context, customers models.GetAllCustomerRequest) ([]models.Customers, error) {
	customer, err := s.storage.Customer().GetAllCustomers(ctx, customers)
	if err != nil {
		s.logger.Error("error in service layer while getting allcars: ", logger.Error(err))
		return nil, err
	}
	return customer.Customers, nil
}

func (s customerSer) UpdateCustomer(ctx context.Context, customers models.Customers) (string, error) {
	customerid, err := s.storage.Customer().UpdateCustomer(ctx, customers)
	if err != nil {
		s.logger.Error("error in service layer while getting allcars: ", logger.Error(err))
		return customerid, err
	}
	return customerid, nil
}

func (s customerSer) DeleteCar(ctx context.Context, id string) (string, error) {
	customerid, err := s.storage.Customer().DeleteCustomer(ctx, id)
	if err != nil {
		s.logger.Error("error in service layer while getting allcars: ", logger.Error(err))
		return customerid, err
	}
	return customerid, nil
}

func (s customerSer) GetAllCustomerCars(ctx context.Context, customers models.GetAllCustomerCarsRequest) (models.GetAllCustomerCarsResponse, error) {
	customer, err := s.storage.Customer().GetAllCustomerCars(ctx, customers)
	if err != nil {
		s.logger.Error("error in service layer while getting allcars: ", logger.Error(err))
		return customer, err
	}
	return customer, nil
}

func (s customerSer) Login(ctx context.Context, req models.Changepasswor) (string, error) {

	hashedPswd, err := s.storage.Customer().GetPassword(ctx, req.Phone)
	if err != nil {
		s.logger.Error("error while getting customer password", logger.Error(err))
		return "", err
	}

	err = password.CompareHashAndPassword(hashedPswd, req.OldPassword)
	if err != nil {
		s.logger.Error("incorrect password", logger.Error(err))
		return "", err
	}
	return "Login successfully", nil
}

func (s customerSer) ChangePassword(ctx context.Context, pass models.Changepasswor) (string, error) {
	msg, err := s.storage.Customer().ChangePassword(ctx, pass)
	if err != nil {
		fmt.Println(err)
		s.logger.Error("failed to change customer password", logger.Error(err))
		return "", err
	}
	return msg, nil
}
