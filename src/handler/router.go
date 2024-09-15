package handler

import (
	"github.com/jasurbek-suyunov/udevs_project/config"
	"github.com/jasurbek-suyunov/udevs_project/middleware"
	"github.com/jasurbek-suyunov/udevs_project/src/service"
	"github.com/jasurbek-suyunov/udevs_project/src/storage/postgres"
	"github.com/jasurbek-suyunov/udevs_project/src/storage/redis"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// SetupRouter sets up the router for the application.
func SetupRouter(cnf *config.Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "page not found",
		})
	})

	//amazon s3 storage
	err := connectS3(cnf)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("aws3 connected")
	}

	// get db
	db, err := postgres.NewPostgres(cnf)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("db connected")
	}

	// get redis
	redis, err := redis.NewRedisCache(cnf)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("redis connected")
	}

	// get service and handler
	services := service.NewService(db, redis)
	handler := NewHandler(services)

	//routes
	r.GET("/ping", handler.Ping)

	// auth routes
	auth := r.Group("auth")
	{
		auth.POST("singup", handler.SignUp)
		auth.POST("signin", handler.SignIn)
		auth.POST("signout", handler.SignOut)
	}

	//auth middleware
	r.Use(middleware.Auth())

	//api version v1 routes
	api := r.Group("api/v1")

	// twit routes
	twit := api.Group("twit")
	{
		twit.GET("", handler.GetTwits)
		twit.POST("", handler.CreateTwit)
		twit.PUT(":id", handler.UpdateTwit)
		twit.GET(":id", handler.GetTwitByID)
		twit.DELETE(":id", handler.DeleteTwit)
		twit.GET("user/:id", handler.GetTwitsByUserID)
		twit.POST("like/:twit_id", handler.LikeTwit)
		twit.POST("unlike/:twit_id", handler.UnLikeTwit)
		twit.POST("retwit/:twit_id", handler.RetwitTwit)
		twit.POST("unretwit/:twit_id", handler.UnRetwitTwit)
	}

	// user routes
	user := api.Group("user")
	{
		user.POST("upload", handler.UploadProfileImage)
		user.POST("follow", handler.FollowUser)
		user.POST("unfollow", handler.UnFollowUser)
		user.GET("followers", handler.GetFollowers)
		user.GET("following", handler.GetFollowing)
		user.GET("following/:user_id", handler.GetFollowingByUserID)
		user.GET("followers/:user_id", handler.GetFollowersByUserID)
	}

	// search routes
	search := api.Group("search")
	{
		search.GET("", handler.Search)
	}

	return r
}
