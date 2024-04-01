package v1

import (
	"net/http"
	models "third-exam/api-gateway/api/handlers/models"
	tokens "third-exam/api-gateway/api/handlers/tokens"
	"third-exam/api-gateway/config"
	"third-exam/api-gateway/pkg/logger"
	"third-exam/api-gateway/services"

	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
	jwthandler     tokens.JWTHandler
}

// handlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	JWTHandler     tokens.JWTHandler
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		jwthandler:     c.JWTHandler,
	}
}

func HandleBadRequestWithErrorMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StandardErrorModel{
			Error: models.Error{
				Message: "Incorrect data supplied",
			},
		})
		l.Error(message, logger.Error(err))
		return true
	}
	return false
}
