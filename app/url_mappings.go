package app

import "github.com/frootlo/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	router.POST("/users", controllers.CreateUser)
	router.GET("/v1/products/list", controllers.GetAllProducts)
	router.GET("/v1/product/:id", controllers.GetProduct)
}
