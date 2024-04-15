package handler

import (
	models "cars_with_sql/api/model"
	"cars_with_sql/pkg/check"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router 		/car [POST]
// @Summary 	create a car
// @Description This api is creates a new car and returns its id
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param		car body models.Car true "car"
// @Success		200  {object}  models.Car
// @Failure		400  {object}  models.Responsee
// @Failure		404  {object}  models.Responsee
// @Failure		500  {object}  models.Responsee
func (h Handler) CreateCars(c *gin.Context) {

	car := models.Car{}

	if err := c.ShouldBindJSON(&car); err != nil {
		handleResponseLog(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())

	}
	if err := check.ValidateCarYear(car.Year); err != nil {
		handleResponseLog(c, h.Log, "error while validating car year, year: "+strconv.Itoa(car.Year), http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Services.Car().Create(context.Background(), car)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating car", http.StatusBadRequest, err.Error())
	}

	handleResponseLog(c, h.Log, "Created successfully", http.StatusOK, id)

}

// GetAllCars godoc
// @Security ApiKeyAuth
// @Router 		/car [GET]
// @Summary 	Get all car
// @Description This api is get a cars
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param		car query string true "cars"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.Car
// @Failure		400  {object}  models.Responsee
// @Failure		404  {object}  models.Responsee
// @Failure		500  {object}  models.Responsee
func (h Handler) GetAllCarss(c *gin.Context) {
	var (
		request = models.GetAllCarsRequest{}
	)

	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)

	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)

	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}
	request.Page = page
	request.Limit = limit

	cars, err := h.Services.Car().GetAllCars(context.Background(), request)
	if err != nil {
		handleResponseLog(c, h.Log, "error while gettign cars", http.StatusBadRequest, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "", http.StatusOK, cars)
}

// @Router 		/car [DELETE]
// @Summary 	delete a car
// @Description This api is delete a car
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param		car body models.Car true "car"
// @Success		200  {object}  models.Car
// @Failure		400  {object}  models.Responsee
// @Failure		404  {object}  models.Responsee
// @Failure		500  {object}  models.Responsee
func (h Handler) DeleteCar(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.Services.Car().DeleteCar(context.Background(), id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while deleting car", http.StatusInternalServerError, err.Error())
		return
	}

}

// @Router 		/car/{id} [GET]
// @Summary 	getsa car
// @Description getncar by ID
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param		id path string true "car"
// @Success		200  {object}  models.Car
// @Failure		400  {object}  models.Responsee
// @Failure		404  {object}  models.Responsee
// @Failure		500  {object}  models.Responsee
func (h Handler) GetbyidCar(c *gin.Context) {
	var (
		request = models.GetAllCarsRequest{}
	)
	id := c.Param("id")

	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
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
	cars, err := h.Services.Car().GetByidcar(context.Background(), id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while gettign cars", http.StatusBadRequest, err.Error())

		return
	}

	handleResponseLog(c, h.Log, "", http.StatusOK, cars)
}

// //////////////////////////////////////////

// @Router 		/car [PUT]
// @Summary 	update a car
// @Description This api is update a car and returns it's id
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param		car body models.UpdateCarRequest true "car"
// @Success		200  {object}  models.Car
// @Failure		400  {object}  models.Responsee
// @Failure		404  {object}  models.Responsee
// @Failure		500  {object}  models.Responsee
func (h Handler) UpdateCar(c *gin.Context) {

	car := models.UpdateCarRequest{}
	if err := c.ShouldBindJSON(&car); err != nil {
		handleResponseLog(c, h.Log, "error while validating car year ,year"+strconv.Itoa(int(car.Year)), http.StatusBadRequest, err.Error())
		return
	}
	car.ID = c.Param("id")

	err := uuid.Validate(car.ID)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating car id,id: "+car.ID, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Services.Car().UpdateCar(context.Background(), car)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating car ", http.StatusBadRequest, err.Error())
	}
	handleResponseLog(c, h.Log, "Updated successfully", http.StatusOK, id)

}
