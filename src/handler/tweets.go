package handler

import (
	"jas/helper"
	"jas/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTweet(c *gin.Context) {
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

	var tweetModel = models.Tweet{
		UserID:    userID.(string),
		Content:   tweet.Content,
		ImageUrl:  tweet.ImageUrl,
		VideuUrl:  tweet.VideoUrl,
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
