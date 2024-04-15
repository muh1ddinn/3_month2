package postgres

import (
	models "cars_with_sql/api/model"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreaterCustomer(t *testing.T) {
	customerRepo := Newcustomer(db, log)

	reqcus := models.Customers{
		First_name: "muhiddin",
		Last_name:  "rmedov",
		Gmail:      "tblf@gmail.com ",
		Phone:      "(998) 9083764 ",
		Is_blocked: true,
	}

	id, err := customerRepo.CreateCus(context.Background(), reqcus)

	if assert.NoError(t, err) {
		CreatedCar, err := customerRepo.GetByIDCustomer(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqcus.First_name, CreatedCar.First_name)
			assert.Equal(t, reqcus.Last_name, CreatedCar.Last_name)
			assert.Equal(t, reqcus.Gmail, CreatedCar.Gmail)
			assert.Equal(t, reqcus.Phone, CreatedCar.Phone)
			assert.Equal(t, reqcus.Is_blocked, CreatedCar.Is_blocked)

		} else {
			return

		}

		fmt.Println("created car", CreatedCar)
	}

}
func TestGetByIDCus(t *testing.T) {
	customerRepo := Newcustomer(db, log)

	id := "0a14208d-6d60-4b73-a22c-e020680bb806"
	expectedName := "Zacharia"
	expectedGmail := "zhotson7@artisteer.com"

	customer, err := customerRepo.GetByIDCustomer(context.Background(), id)
	if assert.NoError(t, err) {
		assert.Equal(t, expectedName, customer.First_name)
		assert.Equal(t, expectedGmail, customer.Gmail)
	} else {
		t.Errorf("Unexpected error getting car by ID: %v", err)
	}
}

func TestGetAllCustomer(t *testing.T) {
	customerRepo := Newcustomer(db, log)

	test := []struct {
		Name    string
		Input   models.GetAllCustomerRequest
		Count   int
		WantErr bool
	}{
		{
			Name: "tesr with valid parameters",
			Input: models.GetAllCustomerRequest{
				Limit:  1,
				Search: "Zacharia",
			},
			Count:   1,
			WantErr: true,
		},
	}

	for _, test := range test {
		t.Run(test.Name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			resp, err := customerRepo.GetAllCustomers(ctx, test.Input)
			if test.WantErr && err == nil {
				fmt.Println("----------------------", test.Name)

				t.Errorf("%s: expected an error,but got nil", test.Name)
				return
			}
			if !test.WantErr && err != nil {
				t.Errorf("%s:unexpected error: %v", test.Name, err)
				return
			}
			if !test.WantErr && len(resp.Customers) != test.Count {
				t.Fatalf("%s: unexpected number of cars, got %d, want %d", test.Name, len(resp.Customers), test.Count)
			}

		})
	}

}
func TestUpdateCustomer(t *testing.T) {
	customerRep := Newcustomer(db, log)

	reqcus := models.Customers{
		First_name: "muhiddin",
		Last_name:  "rmedov",
		Gmail:      "tkt@gmail.com ",
		Phone:      "(998) 9693554 ",
		Is_blocked: true,
		Id:         "21ebe125-5519-4f3b-819a-a93beea5b4cc",
	}

	id, err := customerRep.UpdateCustomer(context.Background(), reqcus)

	fmt.Println(id)
	if assert.NoError(t, err) {
		update, err := customerRep.GetByIDCustomer(context.Background(), id)
		fmt.Print(id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqcus.Last_name, update.Last_name)
			assert.Equal(t, reqcus.Gmail, update.Gmail)
			assert.Equal(t, reqcus.Phone, update.Phone)
			assert.Equal(t, reqcus.Is_blocked, update.Is_blocked)

		} else {
			return

		}

		fmt.Println("update car", update, id)
	}

}
func TestDeleteCustomer(t *testing.T) {
	customerRep := Newcustomer(db, log)

	id := "23af1923-afee-4941-adf6-37e0a2c7789d"

	_, err := customerRep.DeleteCustomer(context.Background(), id)
	if err != nil {
		t.Fatalf("Deletecar failed: %v", err)
	}

	_, err = customerRep.GetByIDCustomer(context.Background(), id)
	if err == nil {
		t.Errorf("Expected car to be deleted, but it still exists")
	}

	fmt.Printf("Car with ID %s successfully deleted", id)
}
