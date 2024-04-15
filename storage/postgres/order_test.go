package postgres

import (
	models "cars_with_sql/api/model"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreaterOder(t *testing.T) {
	OrderRepo := NewOrder(db)

	reqorder := models.CreateOrder{

		CarId:      "92c7ef9d-2b28-4dfd-bdc5-dd1a542ba390",
		CustomerId: "55cac2fb-545b-48ec-8401-0c4f7c235c77",
		FromDate:   "2024-03-25",
		ToDate:     "2024-04-04",
		Status:     "active",
		Paid:       true,
		Amount:     340039.0,
	}

	id, err := OrderRepo.Create(context.Background(), reqorder)

	if assert.NoError(t, err) {
		CreateOrder, err := OrderRepo.GetByID(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqorder.FromDate, CreateOrder.FromDate)
			//assert.Equal(t, reqorder.Amount, CreateOrder.Amount)

		} else {
			return

		}

		fmt.Println("created car", CreateOrder)
	}

}

func TestGetByIDOrder(t *testing.T) {
	neworder := NewOrder(db)

	id := "213fed53-ce65-4119-9b6a-3c5bb327bb5c"
	expectedamount := 0
	expectedfromdat := "2024-03-19"

	order, err := neworder.GetByID(context.Background(), id)
	if assert.NoError(t, err) {
		assert.Equal(t, expectedamount, order.Amount)
		assert.Equal(t, expectedfromdat, order.FromDate)
	} else {
		t.Errorf("Unexpected error getting car by ID: %v", err)
	}

	nonExistentID := ""
	_, err = neworder.GetByID(context.Background(), nonExistentID)
	if err == nil {
		t.Errorf("Expected an error for non-existent car ID, got none")
	}
	fmt.Printf("GetbyID_car%v:", order)

}

func TestGetAllCar(t *testing.T) {
	orderRepo := NewOrder(db)

	testCases := []struct {
		name     string
		req      models.GetAllOrdersRequest
		expected int
	}{
		{"Get 1st page with limit 3", models.GetAllOrdersRequest{Search: "", Page: 15, Limit: 2}, 1},
		{"Search for 'kia' cars", models.GetAllOrdersRequest{Search: "kia"}, 15},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			order, err := orderRepo.GetAll(context.Background(), tc.req)
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, len(order.Orders))
		})
	}
}
