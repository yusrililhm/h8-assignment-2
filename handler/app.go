package handler

import (
	"h8-assignment-2/docs"
	"h8-assignment-2/service"

	"h8-assignment-2/infra/config"
	"h8-assignment-2/infra/database"

	"h8-assignment-2/repository/item_repository/item_pg"
	"h8-assignment-2/repository/order_repository/order_pg"

	"github.com/gin-gonic/gin"

	swaggoFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name Yusril Ilham Kholid
// @contact.email yusrililham@gmail.com
// @contact.url http://github.com/yusrililhm
func StartApplication() {
	app := gin.Default()

	// load env
	config.LoadAppConfig()

	// init database
	database.InitiliazeDatabase()

	// dependencies injection
	db := database.GetDatabaseInstance()

	// swagger
	docs.SwaggerInfo.Title = "Order API Specs"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "Assignment 2 MSIB x Hacktiv8"

	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = append(docs.SwaggerInfo.Schemes, "https", "http")
	docs.SwaggerInfo.Host = "localhost:8080"

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggoFiles.Handler))

	itemRepository := item_pg.NewItemRepositoryImpl(db)
	orderRepository := order_pg.NewOrderRepositoryImpl(db)
	orderService := service.NewOrderServiceImpl(orderRepository, itemRepository)
	orderHandler := NewOrderHandler(orderService)

	// grouping route
	orders := app.Group("/orders")

	{
		orders.POST("/", orderHandler.CreateOrder)
		orders.GET("", orderHandler.FetchOrder)
		orders.PUT("/:orderId", orderHandler.UpdateOrder)
		orders.DELETE("/:orderId", orderHandler.DeleteOrder)
	}

	// run server
	port := config.GetAppConfig().Port
	app.Run(":" + port)
}
