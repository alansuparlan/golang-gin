package main

import (
	"io"
	"os"

	"github.com/alansuparlan/golang-gin/api"
	"github.com/alansuparlan/golang-gin/controller"
	"github.com/alansuparlan/golang-gin/docs"
	"github.com/alansuparlan/golang-gin/middlewares"
	"github.com/alansuparlan/golang-gin/repository"
	"github.com/alansuparlan/golang-gin/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
// @Param Authorization header string true "Insert your access token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWxhbiIsImFkbWluIjp0cnVlLCJleHAiOjE2MTM4OTIwMDcsImlhdCI6MTYxMzYzMjgwNywiaXNzIjoicHJhZ21hdGljcmV2aWV3cy5jb20ifQ.20lvN94xpXkuw8DgQ-AKJSeT-YR_RPebJRD_-047KNo)

func main() {
	setupLogOutput()
	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "alan test - Video API"
	docs.SwaggerInfo.Description = "kiw kiw"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "alan-test-app.herokuapp.com"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"https"}
	defer videoRepository.CloseDB()
	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())

	videoAPI := api.NewVideoAPI(loginController, videoController)

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("/token", videoAPI.Authenticate)
		}

		videos := apiRoutes.Group("/videos", middlewares.AuthorizeJWT())
		{
			videos.GET("", videoAPI.GetVideos)
			videos.POST("", videoAPI.CreateVideo)
			videos.PUT(":id", videoAPI.UpdateVideo)
			videos.DELETE(":id", videoAPI.DeleteVideo)
		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// We can setup this env variable from the EB console
	port := os.Getenv("PORT")

	// Elastic Beanstalk forwards requests to port 5000
	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)
}
