package handler

import (
	models "cars_with_sql/api/model"
	"cars_with_sql/pkg/check"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CustomerLogin godoc
// @Router       /customer/login [POST]
// @Summary      Customer login
// @Description  Customer login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.CustomerLoginRequest true "login"
// @Success      201 {object} models.CustomerLoginResponse
// @Failure      400 {object} models.Responsee
// @Failure      404 {object} models.Responsee
// @Failure      500 {object} models.Responsee
func (h *Handler) CustomerLogin(c *gin.Context) {

	loginReq := models.CustomerLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponseLog(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	fmt.Println("loginReq:", loginReq)

	if err := check.ValidatePassword(loginReq.Password); err != nil {

		handleResponseLog(c, h.Log, "error whilr decoding request body", http.StatusBadRequest, err.Error())

		return

	}

	loginResp, err := h.Services.Auth().CustomerLogin(c.Request.Context(), loginReq)

	if err != nil {
		handleResponseLog(c, h.Log, "unauthorized", http.StatusUnauthorized, err)
		return
	}

	handleResponseLog(c, h.Log, "Succes", http.StatusOK, loginResp)

}

// CustomerRegister godoc
// @Router       /customer/register [POST]
// @Summary      Customer register
// @Description  Customer register
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        register body models.CustomerRegisterRequest true "register"
// @Success      200 {object} models.Responsee
// @Failure      400 {object} models.Responsee
// @Failure      404 {object} models.Responsee
// @Failure      500 {object} models.Responsee
func (h *Handler) CustomerRegister(c *gin.Context) {
	loginReq := models.CustomerRegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponseLog(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}

	if err := check.Validategmail(loginReq.Mail); err != nil {

		handleResponseLog(c, h.Log, "error while putting mail", http.StatusBadRequest, err)

		fmt.Println("loginReq: ", loginReq)

		return

	}

	err := h.Services.Auth().CustomerRegister(c.Request.Context(), loginReq)

	if err != nil {
		fmt.Println(err, "")
		handleResponseLog(c, h.Log, "", http.StatusInternalServerError, err)
		return
	}

	handleResponseLog(c, h.Log, "Otp sent successfull", http.StatusOK, "")

}

// // SingupCustomerRegister godoc
// // @Router       /customer/singup [POST]
// // @Summary      Customer sign up
// // @Description  Customer sign up
// // @Tags         auth
// // @Accept       json
// // @Produce      json
// // @Param        signup body models.CustomerLoginRequest true "signup"
// // @Success      200 {object} models.Responsee
// // @Failure      400 {object} models.Responsee
// // @Failure      500 {object} models.Responsee
// func (h *Handler) SingupCustomerRegister(c *gin.Context) {
// 	loginReq := models.CustomerLoginRequest{}

// 	if err := c.ShouldBindJSON(&loginReq); err != nil {
// 		handleResponseLog(c, h.Log, "error while biding body", http.StatusBadRequest, err)

// 		return
// 	}
// 	fmt.Println("loginReq:", loginReq)

// 	if err := check.ValidatePassword(loginReq.Password); err != nil {
// 		handleResponseLog(c, h.Log, "error while putting mail", http.StatusBadRequest, err)

// 		return
// 	}

// 	handleResponseLog(c, h.Log, "Otp sent successfull", http.StatusOK, "")

// }
