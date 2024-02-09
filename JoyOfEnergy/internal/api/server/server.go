package server

import (
	"JoyOfEnergy/internal/api/router"
	"JoyOfEnergy/pkg/logger"
	"github.com/gin-gonic/gin"
	"strings"
)

func Init() {
	log := logger.NewLogger()
	if strings.EqualFold("prodcution", "production") {
		gin.SetMode(gin.ReleaseMode)
	}
	r := router.NewRouter()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	err := r.Run(":8080")
	if err != nil {
		log.Error("Error occured while starting server....")
	}
}
