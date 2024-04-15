package service

import (
	models "cars_with_sql/api/model"
	"cars_with_sql/pkg/logger"
	"cars_with_sql/storage"
	"context"
	"fmt"
)

type orderSer struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewOrderservice(storage storage.IStorage, logger logger.ILogger) orderSer {
	return orderSer{
		storage: storage,
		logger:  logger,
	}
}

func (u orderSer) Create(ctx context.Context, order models.CreateOrder) (string, error) {

	pkey, err := u.storage.Order().Create(ctx, order)
	if err != nil {
		fmt.Println("eror in servince layer while creating order", err.Error())
		return "", err

	}

	return pkey, nil

}

func (u orderSer) GetAllOrder(ctx context.Context, req models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error) {
	order, err := u.storage.Order().GetAll(ctx, req)
	if err != nil {
		fmt.Println("error in service layer while getting allcars: ", err.Error())
		return order, err
	}
	return order, nil
}

func (u orderSer) GetbyID(ctx context.Context, id string) (models.GetOrder, error) {
	order, err := u.storage.Order().GetByID(ctx, id)
	if err != nil {
		fmt.Println("error in service layer while getting allcars: ", err.Error())
		return order, err
	}
	return order, nil
}

func (u orderSer) UpdateOrder(ctx context.Context, order models.CreateOrder) (string, error) {

	pkey, err := u.storage.Order().UpdateOrders(ctx, order)
	if err != nil {
		fmt.Println("eror in servince layer while creating order", err.Error())
		return "", err

	}

	return pkey, nil

}
