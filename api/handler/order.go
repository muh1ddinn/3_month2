package handler

import (
	models "cars_with_sql/api/model"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateOrder godoc
// @Security ApiKeyAuth
// @Router		/order [POST]
// @Summary		create an order
// @Description This api creates a new order and returns its id
// @Tags		order
// @Accept		json
// @Produce		json
// @Param		order body models.CreateOrder true "order"
// @Success		200  {string}  string
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response

func (h Handler) CreateOrder(c *gin.Context) {

	cus := models.CreateOrder{}

	if err := c.ShouldBindJSON(&cus); err != nil {
		handleResponseLog(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return

	}

	id, err := h.Services.Order().Create(context.Background(), cus)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating order", http.StatusBadRequest, err.Error())

		return
	}
	handleResponseLog(c, h.Log, "Created successfully", http.StatusOK, id)

}

// GetAllOrders godoc
// @Security ApiKeyAuth
// @Router		/order [GET]
// @Summary		get all orders
// @Description This api gets all orders and returns their info
// @Tags		order
// @Accept		json
// @Produce		json
// @Param		order query string true "orders"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.GetAllOrdersResponse
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response

func (h Handler) GetAllOrder(c *gin.Context) {
	var (
		request = models.GetAllOrdersRequest{}
	)

	order, err := h.Services.Order().GetAllOrder(context.Background(), request)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting order", http.StatusBadRequest, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "", http.StatusOK, order)
}

// GetOrderByID godoc
// @Security ApiKeyAuth
// @Router		/order/{id} [GET]
// @Summary		get an order by its id
// @Description This api gets a order by its id and returns its info
// @Tags		order
// @Accept		json
// @Produce		json
// @Param		id path string true "order"
// @Success		200  {object}  models.GetOrderResponse
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response

func (h Handler) GetByID(c *gin.Context) {

	id := c.Param("id")

	cus, err := h.Services.Order().GetbyID(context.Background(), id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting  getbyorder", http.StatusBadRequest, err.Error())

		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, cus)
}

// func (h Handler) Deleteorder(c *gin.Context) {

// 	id := c.Param("id")
// 	fmt.Println("id: ", id)

// 	err := uuid.Validate(id)
// 	if err != nil {
// 		handleResponse(c, "error while validating id", http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	err = h.Store.Order().DeleteOrder(id)
// 	if err != nil {

// 		handleResponse(c, "error while deleting car", http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	handleResponse(c, "", http.StatusOK, id)
// }

// UpdateOrder godoc
// @Security ApiKeyAuth
// @Router		/order/{id} [PUT]
// @Summary		update an order
// @Description This api updates a order by its id and returns its id
// @Tags		order
// @Accept		json
// @Produce		json
// @Param		id path string true "order id"
// @Param		order body models.UpdateOrder true "order"
// @Success		200  {string}  string
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response

func (h Handler) UpdateOrder(c *gin.Context) {

	cus := models.CreateOrder{}

	if err := c.ShouldBindJSON(&cus); err != nil {
		handleResponseLog(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return

	}
	cus.Id = c.Param("id")

	err := uuid.Validate(cus.Id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating car id,id: "+cus.Id, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Services.Order().UpdateOrder(context.Background(), cus)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating car ", http.StatusBadRequest, err.Error())
	}
	handleResponseLog(c, h.Log, "Updated successfully", http.StatusOK, id)
}
