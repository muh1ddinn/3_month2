package handler

import (
	models "cars_with_sql/api/model"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Createcus godoc
// @Security ApiKeyAuth
// @Router      /customer [POST]
// @Summary     Create a customer
// @Description Create a new customer
// @Tags        customer
// @Accept      json
// @Produce 	json
// @Param 		customer body models.Customers true "customer"
// @Success 	200  {object}  string
// @Failure		400  {object}  models.Responsee
// @Failure		404  {object}  models.Responsee
// @Failure		500  {object}  models.Responsee
func (h Handler) Createcus(c *gin.Context) {
	cus := models.Customers{}

	if err := c.ShouldBindJSON(&cus); err != nil {
		return
	}

	id, err := h.Services.Customer().CreateCus(context.Background(), cus)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating car", http.StatusBadRequest, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "Created successfully", http.StatusOK, id)
}

// GetAllCustomers godoc
// @Security ApiKeyAuth
// @Router 			/customer [GET]
// @Summary 		Get all customers
// @Description		Retrieves information about all customers.
// @Tags 			customer
// @Accept 			json
// @Produce 		json
// @Param 			search query string true "customer"
// @Param 			page query uint64 false "page"
// @Param 			limit query uint64 false "limit"
// @Success 		200 {object} models.GetAllCustomersResponse
// @Failure 		400 {object} models.Responsee
// @Failure 		500 {object} models.Responsee
func (h Handler) Getallcus(c *gin.Context) {

	var request = models.GetAllCustomerRequest{}

	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)

	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing page", http.StatusBadRequest, err.Error())

	}

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit

	customer, err := h.Services.Customer().GetAllCus(context.Background(), request)
	if err != nil {

		handleResponseLog(c, h.Log, "error while getting customer", http.StatusBadRequest, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "", http.StatusOK, customer)
}

// GetCustomerById godoc
// @Security ApiKeyAuth
// @Router		/customer/{id} [GET]
// @Summary		get a customer by its id
// @Description This api gets a customer by its id and returns its info
// @Tags		customer
// @Accept		json
// @Produce		json
// @Param		id path string true "customer"
// @Success		200  {object}  models.Customers
// @Failure		400  {object}  models.Responsee
// @Failure		404  {object}  models.Responsee
// @Failure		500  {object}  models.Responsee
func (h Handler) GetByIDCus(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		handleResponseLog(c, h.Log, "missing car ID", http.StatusBadRequest, id)
		return
	}

	customer, err := h.Services.Customer().GetByID(context.Background(), id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting customer by ID", http.StatusBadRequest, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Customer was successfully gotten by Id", http.StatusOK, customer)
}

// DeleteCustomer godoc
// @Security ApiKeyAuth
// @Router		/customer/{id} [DELETE]
// @Summary		delete a customer by its id
// @Description This api deletes a customer by its id and returns error or nil
// @Tags		customer
// @Accept		json
// @Produce		json
// @Param		id path string true "customer ID"
// @Success		200  {object}  nil
// @Failure		400  {object}  models.Responsee
// @Failure		404  {object}  models.Responsee
// @Failure		500  {object}  models.Responsee
func (h Handler) Deletecus(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating id ", http.StatusBadRequest, err.Error())
		return
	}

	id, err = h.Services.Customer().DeleteCar(context.Background(), id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while deleting customer", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "all good ", http.StatusOK, id)

}

// GetCustomerCars godoc
// @Security ApiKeyAuth
// @Router		/customer/cars [GET]
// @Summary		get customer's cars
// @Description This api gets customer cars and returns their info
// @Tags		customer
// @Accept		json
// @Produce		json
// @Param 		customerID query string true "Customer ID"
// @Param 		carName query string false "Car Name"
// @Success		200  {object}  models.GetAllCustomerCarsResponse
// @Failure		400  {object}  models.Responsee
// @Failure		404  {object}  models.Responsee
// @Failure		500  {object}  models.Responsee
func (h Handler) GetAllCustomerCars(c *gin.Context) {
	var (
		request = models.GetAllCustomerCarsRequest{}
	)

	request.Search = c.Query("search")
	request.Id = c.Param("id")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing page", http.StatusBadRequest, err.Error())
	}
	limit, err := ParsePageQueryParam(c)

	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("page:", page)
	fmt.Println("limit:", limit)

	request.Page = page
	request.Limit = limit
	Order, err := h.Services.Customer().GetAllCustomerCars(context.Background(), request)
	if err != nil {
		handleResponseLog(c, h.Log, "eror while getting Customers", http.StatusBadRequest, err.Error())

		return
	}

	handleResponseLog(c, h.Log, "", http.StatusOK, Order)
}

/*
// LoginCustomer godoc
// @Router		/customer/login [POST]
// @Summary		customer login
// @Description This api logs in customer account and returns message
// @Tags		customer
// @Accept		json
// @Produce		json
// @Param		customer body models.LoginCustomer true "customer"
// @Success		200  {object}  string
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) LoginCustomer(c *gin.Context) {
	login := models.Changepasswor{}

	if err := c.ShouldBindJSON(&login); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidatePassword(login.OldPassword); err != nil {
		handleResponseLog(c, h.Log, "error while validating password", http.StatusBadRequest, err.Error())
		return
	}

	msg, err := h.Services.Customer().Login(c.Request.Context(), login)
	if err != nil {
		handleResponseLog(c, h.Log, "error while logging", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "", http.StatusOK, msg)
}


// LoginCustomer godoc
// @Security ApiKeyAuth
// @Router		/customer/ [PATCH]
// @Summary		customer change password
// @Description This api changes customer password by its login and password and returns message
// @Tags		customer
// @Accept		json
// @Produce		json
// @Param		customer body models.ChangePassword true "Change Customer Password"
// @Success		200  {object}  string
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) ChangePasswordCustomer(c *gin.Context) {
	pass := models.Changepasswor{}

	if err := c.ShouldBindJSON(&pass); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidatePassword(pass.OldPassword); err != nil {
		handleResponseLog(c, h.Log, "error while validating new password", http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass.OldPassword), bcrypt.DefaultCost)
	if err != nil {
		handleResponseLog(c, h.Log, "error while hashing new password", http.StatusInternalServerError, err.Error())
		return
	}
	pass.OldPassword = string(hashedPassword)

	msg, err := h.Services.Customer().ChangePassword(context.Background(), pass)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating customer", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Customer password was successfully updated", http.StatusOK, msg)
}



// UpdateCustomer godoc
// @Security ApiKeyAuth
// @Router		/customer/{id} [PUT]
// @Summary		update a customer
// @Description This api updates a customer by its id and returns its id
// @Tags		customer
// @Accept		json
// @Produce		json
// @Param 		id path string true "Customer ID"
// @Param		customer body models.UpdateCustomer true "customer"
// @Success		200  {object}  string
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) Updacustomer(c *gin.Context) {

	customer := models.Customers{}
	if err := c.ShouldBindJSON(&customer); err != nil {
		return
	}
	customer.Id = c.Param("id")

	err := uuid.Validate(customer.Id)
	if err != nil {
		return
	}
	id, err := h.Services.Customer().UpdateCustomer(context.Background(), customer)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating car ", http.StatusBadRequest, err.Error())
	}
	handleResponseLog(c, h.Log, "Updated successfully", http.StatusOK, id)

}

*/
