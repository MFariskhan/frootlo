package app

import (
	"github.com/gin-gonic/gin"
	"os"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	var port = os.Getenv("PORT")

	router.Run(":" + port)
}
