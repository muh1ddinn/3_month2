package api

import (
	"cars_with_sql/api/handler"
	"cars_with_sql/pkg/logger"
	"cars_with_sql/service"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "cars_with_sql/api/docs"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(services service.IServiceMangaer, log logger.ILogger) *gin.Engine {
	h := handler.NewStrg(services, log)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//r.Use(authMiddleware)
	r.POST("/customer/login", h.CustomerLogin)

	r.POST("/car", h.CreateCars)
	r.PUT("/car/:id", h.UpdateCar)
	r.GET("/car/:id", h.GetbyidCar)
	r.GET("/car", h.GetAllCarss)
	//r.GET("car/available", h.GetAvailableCars)
	r.DELETE("/car/:id", h.DeleteCar)

	r.POST("/customer", h.Createcus)
	//	r.PUT("/customer/:id", h.Updacustomer)
	//r.PATCH("/customer", h.ChangePasswordCustomer)
	r.GET("/customer/:id", h.GetByIDCus)
	r.GET("/customer", h.Getallcus)
	r.GET("/customer/cars", h.GetAllCustomerCars)
	r.DELETE("/customer/:id", h.Deletecus)

	r.POST("/order", h.CreateOrder)
	r.PUT("/order/:id", h.UpdateOrder)
	r.GET("/order/:id", h.GetByID)
	r.GET("/order", h.GetAllOrder)

	r.POST("/customer/register", h.CustomerRegister)

	//r.DELETE("/order/:id", h.)

	return r
}

// func authMiddleware(c *gin.Context) {
// 	auth := c.GetHeader("Authorization")
// 	if auth == "" {
// 		c.AbortWithError(http.StatusUnauthorized, errors.New("____________________-unauthorized"))
// 	}
// 	c.Next()
// }
