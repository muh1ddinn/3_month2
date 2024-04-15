package main

import (
	"cars_with_sql/api"
	"cars_with_sql/config"
	"cars_with_sql/pkg/logger"
	"cars_with_sql/service"

	"cars_with_sql/storage/postgres"
	"context"
	"fmt"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.ServiceName)

	store, err := postgres.New(context.Background(), cfg, log)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.CloseDB()

	service := service.New(store, log)

	c := api.New(service, log)

	fmt.Println("programm is running on localhost:9090...")
	c.Run(":9090")
}
