package app

import (
	"github.com/dharmatin/bookstore-user-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.GetLogger().Info("Application Started")
	router.Run(":8081")
}
