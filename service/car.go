package service

import (
	models "cars_with_sql/api/model"
	"cars_with_sql/pkg/logger"
	"cars_with_sql/storage"
	"context"
	"fmt"
)

type carService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewCarService(storage storage.IStorage, logger logger.ILogger) carService {
	return carService{
		storage: storage,
		logger:  logger,
	}
}

func (u carService) Create(ctx context.Context, car models.Car) (string, error) {

	pKey, err := u.storage.Car().Create(ctx, car)
	if err != nil {
		fmt.Println("error in service layer while creating car ", err.Error())

		return "", err
	}

	return pKey, nil

}

func (s carService) GetAllCars(ctx context.Context, car models.GetAllCarsRequest) ([]models.Car, error) {
	cars, err := s.storage.Car().GetAllCarS(ctx, car)
	if err != nil {
		fmt.Println("error in service layer while getting allcars: ", err.Error())
		return nil, err
	}
	return cars.Cars, nil
}

func (s carService) GetByidcar(ctx context.Context, id string) (models.Car, error) {
	cars, err := s.storage.Car().GetByidcar(ctx, id)
	if err != nil {
		fmt.Println("error in service layer while getting id cars: ", err.Error())
		return models.Car{}, err
	}
	return cars, nil
}

func (u carService) UpdateCar(ctx context.Context, car models.UpdateCarRequest) (string, error) {

	pKey, err := u.storage.Car().UpdateCar(ctx, car)
	if err != nil {
		fmt.Println("error in service layer while updating car ", err.Error())

		return "", err
	}

	return pKey, nil

}

func (u carService) DeleteCar(ctx context.Context, id string) (string, error) {

	pKey, err := u.storage.Car().Deletecar(ctx, id)
	if err != nil {
		fmt.Println("error in service layer while updating car ", err.Error())

		return "", err
	}

	return pKey, nil

}
