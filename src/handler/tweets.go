package handler

import (
	"errors"
	"github.com/jasurbek-suyunov/udevs_project/helper"
	"github.com/jasurbek-suyunov/udevs_project/models"

	"github.com/gin-gonic/gin"
)

const (
	tweets_folder = "tweets"
)

func (h *Handler) CreateTweet(c *gin.Context) {

	// Parse form file
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "No file is received"})
		return
	}
	defer file.Close()

	// Get tweet
	tweet := c.PostForm("tweet")

	// Upload to S3
	fileURL, err := uploadToS3(file, fileHeader, tweets_folder)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	var tweetModel = models.Tweet{
		UserID:    userID.(string),
		Content:   tweet,
		MediaUrl:  fileURL,
		CreatedAt: helper.GetCurrentTime(),
	}

	// create tweet
	createdTweet, err := h.services.CreateTweet(c, &tweetModel)
	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return result if no error
	c.JSON(201, createdTweet)
}

func (h *Handler) GetTweetsByUserID(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get tweets
	tweets, err := h.services.GetTweetsByUserID(c, userID.(string))
	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return result if no error
	c.JSON(200, tweets)
}

func (h *Handler) DeleteTweet(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get tweet id
	tweetID := c.Param("id")
	if len(tweetID) == 0 {
		c.JSON(400, gin.H{"error": errors.New("invalid tweet id")})
		return
	}
	// delete tweet
	err := h.services.DeleteTweet(c, userID.(string), tweetID)
	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return result if no error
	c.JSON(200, gin.H{"message": "tweet deleted"})
}

func (h *Handler) UpdateTweet(c *gin.Context) {
	// variable
	var tweet models.TweetRequest

	// bind
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get tweet id
	tweetID := c.Param("id")

	var tweetModel = models.Tweet{
		ID:        tweetID,
		UserID:    userID.(string),
		Content:   tweet.Content,
		CreatedAt: helper.GetCurrentTime(),
	}

	// update tweet
	updatedTweet, err := h.services.UpdateTweet(c, &tweetModel)
	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return result if no error
	c.JSON(200, updatedTweet)
}

func (h *Handler) GetTweetByID(c *gin.Context) {
	//get tweet id
	tweetID := c.Param("id")

	// get tweet
	tweet, err := h.services.GetTweetByID(c, tweetID)
	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return result if no error
	c.JSON(200, tweet)
}

func (h *Handler) GetTweets(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get tweets
	tweets, err := h.services.GetTweets(c, userID.(string))
	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return result if no error
	c.JSON(200, tweets)
}

func (h *Handler) LikeTweet(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get tweet id
	tweetID := c.Param("tweet_id")
	if len(tweetID) == 0 {
		c.JSON(400, gin.H{"error": errors.New("invalid tweet id")})
		return
	}

	err := h.services.LikeTweet(c, userID.(string), tweetID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "like saved"})
}

func (h *Handler) UnLikeTweet(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get tweet id
	tweetID := c.Param("tweet_id")
	if len(tweetID) == 0 {
		c.JSON(400, gin.H{"error": errors.New("invalid tweet id")})
		return
	}

	err := h.services.UnLikeTweet(c, userID.(string), tweetID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "unlike saved"})
}

func (h *Handler) RetweetTweet(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get tweet id
	tweetID := c.Param("tweet_id")
	if len(tweetID) == 0 {
		c.JSON(400, gin.H{"error": errors.New("invalid tweet id")})
		return
	}

	err := h.services.RetweetTweet(c, userID.(string), tweetID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "retweet saved"})
}

func (h *Handler) UnRetweetTweet(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get tweet id
	tweetID := c.Param("tweet_id")
	if len(tweetID) == 0 {
		c.JSON(400, gin.H{"error": errors.New("invalid tweet id")})
		return
	}

	err := h.services.UnRetweetTweet(c, userID.(string), tweetID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "unretweet saved"})
}