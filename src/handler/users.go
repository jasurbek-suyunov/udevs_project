package handler

import (
	"fmt"
	"jas/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FollowUser(c *gin.Context) {
	// variable
	var follow models.FollowRequest

	// bind
	if err := c.ShouldBindJSON(&follow); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}
	fmt.Println(userID)
	// create follow
	err := h.services.FollowUser(c, userID.(int), follow.FollowedID)

	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return result if no error
	c.JSON(201, gin.H{"message": "User followed successfully"})
}

func (h *Handler) UnFollowUser(c *gin.Context) {
	// variable
	var follow models.FollowRequest

	// bind
	if err := c.ShouldBindJSON(&follow); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// create follow
	err := h.services.UnFollowUser(c, userID.(int), follow.FollowedID)

	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return result if no error
	c.JSON(201, gin.H{"message": "User unfollowed successfully"})
}

func (h *Handler) IsFollowing(c *gin.Context) {
	// variable
	var follow models.FollowRequest

	// bind
	if err := c.ShouldBindJSON(&follow); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// create follow
	result, err := h.services.IsFollowing(c, userID.(int), follow.FollowedID)

	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return result if no error
	c.JSON(200, gin.H{"result": result})
}

func (h *Handler) GetFollowers(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// create follow
	result, err := h.services.GetFollowers(c, userID.(int))

	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return result if no error
	c.JSON(200, gin.H{"result": result})
}

func (h *Handler) GetFollowing(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// create follow
	result, err := h.services.GetFollowing(c, userID.(int))

	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return result if no error
	c.JSON(200, gin.H{"result": result})
}