package handler

import (
	"jas/config"
	"jas/src/service"
	"jas/src/storage/postgres"
	"jas/src/storage/redis"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter sets up the router for the application.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// get configs
	cnf := config.NewConfig()
	r.Static("/uploads", "./uploads")
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// get db
	db, err := postgres.NewPostgres(cnf)
	// check error
	if err != nil {
		log.Println(err)
	} else {
		log.Println("db connected")
	}

	// get redis
	redis, err := redis.NewRedisCache(cnf)
	// check error
	if err != nil {
		log.Println(err)
	} else {
		log.Println("redis connected")
	}

	// get service and handler
	services := service.NewService(db, redis)
	handler := NewHandler(services)

	// Routes
	//r.GET("/ping", handler.Ping)

	// swagger
	swagger := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swagger))


	auth := r.Group("auth")
	{
		auth.POST("singup", handler.SignUp)
		auth.POST("signin", handler.SignIn)
		auth.POST("signout", handler.SignOut)
	}


	// app.Use(middlewares.Auth())
	// url := app.Group("url")
	// {
	// 	url.POST("", handler.CreateUrl)
	// 	url.GET("", handler.GetUrls)
	// 	url.GET(":id", handler.GetUrlByID)
	// 	url.DELETE(":id", handler.DeleteUrl)
	// }

	return r
}
