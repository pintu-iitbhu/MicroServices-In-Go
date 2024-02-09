package controller

import (
	"JoyOfEnergy/internal/constants"
	"JoyOfEnergy/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}
func (h *HealthController) Status() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		log := logger.NewLogger()
		version, err := os.ReadFile(constants.GetDeploymentVersionPath)
		if err != nil {
			log.Info("failed to read file deployedVersion")
		}
		c.JSON(http.StatusOK, gin.H{constants.VersionKey: string(version)})
		return
	}

	return fn

}
