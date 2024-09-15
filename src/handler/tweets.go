package handler

import (
	"errors"
	"fmt"
	"jas/helper"
	"jas/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTweet(c *gin.Context) {
	// Parse form file
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}
	defer file.Close()
	// get "tweet" from the form data
	tweet := c.PostForm("tweet")
	fmt.Println(tweet)
	// Folder name where the file will be saved
	folder := "test"
	// Upload to S3
	// how to use .png or .jpg or video file
	fmt.Println(fileHeader.Filename)
	fileURL, err := uploadToS3(file, fileHeader, folder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		MediaUrl:  tweet.MediaUrl,
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
