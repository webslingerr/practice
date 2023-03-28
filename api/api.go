package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, store, logger)

	r.POST("/customer", handler.CreateCustomer)
	r.GET("/customer/:id", handler.GetByIdCustomer)
	r.GET("/customer", handler.GetListCustomer)	
	r.PUT("/customer/:id", handler.UpdateCustomer)
	r.PATCH("/customer/:id", handler.UpdatePatchCustomer)
	r.DELETE("/customer/:id", handler.DeleteCustomer)

	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)	
	r.PUT("/user/:id", handler.UpdateUser)
	r.PATCH("/user/:id", handler.UpdatePatchUser)
	r.DELETE("/user/:id", handler.DeleteUser)

	r.POST("/courier", handler.CreateCourier)
	r.GET("/courier/:id", handler.GetByIdCourier)
	r.GET("/courier", handler.GetListCourier)	
	r.PUT("/courier/:id", handler.UpdateCourier)
	r.PATCH("/courier/:id", handler.UpdatePatchCourier)
	r.DELETE("/courier/:id", handler.DeleteCourier)

	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)	
	r.PUT("/category/:id", handler.UpdateCategory)
	r.PATCH("/category/:id", handler.UpdatePatchCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)

	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)	
	r.PUT("/product/:id", handler.UpdateProduct)
	r.PATCH("/product/:id", handler.UpdatePatchProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)

	r.POST("/order", handler.CreateOrder)
	r.GET("/order/:id", handler.GetByIdOrder)
	r.GET("/order", handler.GetListOrder)	
	r.PUT("/order/:id", handler.UpdateOrder)
	r.PATCH("/order/:id", handler.UpdatePatchOrder)
	r.DELETE("/order/:id", handler.DeleteOrder)


	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}