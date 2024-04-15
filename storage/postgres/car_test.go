package postgres

import (
	models "cars_with_sql/api/model"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreaterCar(t *testing.T) {
	carRepo := Newwcar(db)

	reqCar := models.Car{
		Name:       "k5",
		Year:       2023,
		Brand:      "sedan",
		Model:      "kia",
		HorsePower: 180,
		Colour:     "black",
		EngineCap:  1.6,
	}

	id, err := carRepo.Create(context.Background(), reqCar)

	if assert.NoError(t, err) {
		CreatedCar, err := carRepo.GetByidcar(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCar.Name, CreatedCar.Name)
			assert.Equal(t, reqCar.Year, CreatedCar.Year)
			assert.Equal(t, reqCar.Brand, CreatedCar.Brand)
			assert.Equal(t, reqCar.Model, CreatedCar.Model)
			assert.Equal(t, reqCar.HorsePower, CreatedCar.HorsePower)
			assert.Equal(t, reqCar.Colour, CreatedCar.Colour)
			assert.Equal(t, reqCar.EngineCap, CreatedCar.EngineCap)

		} else {
			return

		}

		fmt.Println("created car", CreatedCar)
	}

}
func TestGetByID(t *testing.T) {
	carRepo := Newwcar(db)

	id := "13d31485-c446-4275-b2d5-748f383d2b48"
	expectedName := "nexii2_2"
	expectedYear := 2023

	car, err := carRepo.GetByidcar(context.Background(), id)
	if assert.NoError(t, err) {
		assert.Equal(t, expectedName, car.Name)
		assert.Equal(t, expectedYear, car.Year)
	} else {
		t.Errorf("Unexpected error getting car by ID: %v", err)
	}

	nonExistentID := ""
	_, err = carRepo.GetByidcar(context.Background(), nonExistentID)
	if err == nil {
		t.Errorf("Expected an error for non-existent car ID, got none")
	}
	fmt.Printf("GetbyID_car%v:", car)

}

func TestGetAllcars(t *testing.T) {
	carRepo := Newwcar(db)

	test := []struct {
		Name    string
		Input   models.GetAllCarsRequest
		Count   int
		WantErr bool
	}{
		{
			Name: "tesr with valid parameters",
			Input: models.GetAllCarsRequest{
				Limit:  10,
				Search: "k6",
			},
			Count:   1,
			WantErr: false,
		},
	}

	for _, test := range test {
		t.Run(test.Name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			resp, err := carRepo.GetAllCarS(ctx, test.Input)
			if test.WantErr && err == nil {
				t.Errorf("%s: expected an error,but got nil", test.Name)

				return
			}
			if !test.WantErr && err != nil {
				t.Errorf("%s:unexpected error: %v", test.Name, err)
				return
			}
			if !test.WantErr && len(resp.Cars) != test.Count {
				t.Fatalf("%s: unexpected number of cars, got %d, want %d", test.Name, len(resp.Cars), test.Count)
			}

		})
	}

}
func TestUpdateCar(t *testing.T) {
	carRepo := Newwcar(db)

	reqCar := models.UpdateCarRequest{

		Name:       "k5",
		Year:       2023,
		Brand:      "sedan",
		Model:      "kia",
		HorsePower: 180,
		Colour:     "black",
		EngineCap:  1.6,
		ID:         "e5f72a5b-ad3c-49c5-9aeb-910635ea9840",
	}

	id, err := carRepo.UpdateCar(context.Background(), reqCar)

	if assert.NoError(t, err) {
		update, err := carRepo.GetByidcar(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCar.Name, update.Name)
			assert.Equal(t, reqCar.Year, update.Year)
			assert.Equal(t, reqCar.Brand, update.Brand)
			assert.Equal(t, reqCar.Model, update.Model)
			assert.Equal(t, reqCar.HorsePower, update.HorsePower)
			assert.Equal(t, reqCar.Colour, update.Colour)
			assert.Equal(t, reqCar.EngineCap, update.EngineCap)

		} else {
			return

		}

		fmt.Println("update car", update, id)
	}

}

// func TestDeleteCar(t *testing.T) {
// 	carRepo := Newwcar(db)

// 	id := "13d31485-c446-4275-b2d5-748f383d2b48"

// 	_, err := carRepo.Deletecar(context.Background(), id)
// 	if err != nil {
// 		t.Fatalf("Deletecar failed: %v", err)
// 	}

// 	_, err = carRepo.GetByidcar(context.Background(), id)
// 	if err == nil {
// 		t.Errorf("Expected car to be deleted, but it still exists")
// 	}

// 	fmt.Printf("Car with ID %s successfully deleted", id)
// }

// func TestDeleteCar(t *testing.T) {
// 	carRepo := Newwcar(db)

// 	id := "21ebe125-5519-4f3b-819a-a93beea5b4cc"

// 	_, err := carRepo.Deletecar(context.Background(), id)
// 	if err != nil {
// 		t.Fatalf("Deletecar failed: %v", err)
// 	}

// 	_, err = carRepo.GetByidcar(context.Background(), id)
// 	if err == nil {
// 		t.Errorf("Expected car to be deleted, but it still exists")
// 	}

//		fmt.Printf("Car with ID %s successfully deleted", id)
//	}
func TestDeleteCar(t *testing.T) {
	// Assuming you have a db connection initialized somewhere

	carRepo := Newwcar(db)
	carID := "27d54c78-8091-4ba0-860b-ee63cc8934c0"

	_, err := carRepo.Deletecar(context.Background(), carID) // Method name should be DeleteCar, not Deletecar

	if err != nil { // Check if error is not nil
		t.Errorf("Error deleting car with ID %s: %v", carID, err)
	}
}
