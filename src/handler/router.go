package handler

import (
	"jas/config"
	"jas/middleware"
	"jas/src/service"
	"jas/src/storage/postgres"
	"jas/src/storage/redis"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// SetupRouter sets up the router for the application.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "page not found",
		})
	})

	// get configs
	cnf := config.NewConfig()

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

	//Auth middleware
	r.Use(middleware.Auth())

	// tweet routes
	tweet := r.Group("tweet")
	{
		tweet.GET("", handler.GetTweets)
		tweet.POST("", handler.CreateTweet)
		tweet.PUT(":id", handler.UpdateTweet)
		tweet.GET(":id", handler.GetTweetByID)
		tweet.DELETE(":id", handler.DeleteTweet)
		tweet.GET("user/:id", handler.GetTweetsByUserID)
		tweet.POST("like/:tweet_id", handler.LikeTweet)
		tweet.POST("unlike/:tweet_id", handler.UnLikeTweet)
		tweet.POST("retweet/:tweet_id", handler.RetweetTweet)
		tweet.POST("unretweet/:tweet_id", handler.UnRetweetTweet)
	}

	// user routes
	user := r.Group("user")
	{
		user.POST("follow", handler.FollowUser)
		user.POST("unfollow", handler.UnFollowUser)
		user.GET("followers", handler.GetFollowers)
		user.GET("following", handler.GetFollowing)
		user.GET("following/:user_id", handler.GetFollowingByUserID)
		user.GET("followers/:user_id", handler.GetFollowersByUserID)
	}

	// search routes
	search := r.Group("search")
	{
		search.GET("", handler.Search)
	}

	return r
}
