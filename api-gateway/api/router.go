package api

import (
	_ "third-exam/api-gateway/api/docs" // swag
	v1 "third-exam/api-gateway/api/handlers/v1"
	"third-exam/api-gateway/config"
	"third-exam/api-gateway/pkg/logger"
	"third-exam/api-gateway/services"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// @Title Welcome to ProductAPI
// @Version 1.0
// @Description This is a example of Social Network
// @Host localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")

	// users
	api.POST("/create", handlerV1.CreateUser)
	api.GET("/get/:id", handlerV1.GetUser)
	api.GET("/all", handlerV1.GetALlUsers)
	api.PUT("/update/:id", handlerV1.UpdateUser)
	api.DELETE("/delete/:id", handlerV1.DeleteUser)

	// post
	api.POST("/create", handlerV1.CreatePost)
	api.GET("/get/:id", handlerV1.GetPost)
	api.GET("/all", handlerV1.GetAllPosts)
	api.PUT("/update/:id", handlerV1.UpdatePost)
	api.DELETE("/delete/:id", handlerV1.DeletePost)

	// register
	api.POST("/sign", handlerV1.SignUp)
	api.GET("/login", handlerV1.LogIn)
	api.GET("/verification", handlerV1.Verification)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
