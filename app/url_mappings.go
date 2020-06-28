package app

import "github.com/frootlo/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
