package handler

import (
	"github.com/jasurbek-suyunov/udevs_project/models"
	"strconv"

	"github.com/gin-gonic/gin"
)
const (
	users_folder = "users"
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

	//convert user id to int
	userIDInt, err := strconv.Atoi(userID.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// check if user is following himself
	if userIDInt == follow.FollowedID {
		c.JSON(400, gin.H{"error": "You can't follow yourself"})
		return
	}

	// check if user is already following
	isFollowing, err := h.services.IsFollowing(c, userIDInt, follow.FollowedID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if isFollowing {
		c.JSON(400, gin.H{"error": "You are already following this user"})
		return
	}

	//create follow
	err = h.services.FollowUser(c, userIDInt, follow.FollowedID)

	// check error
	if err != nil {
		c.JSON(400, gin.H{"error":"error following user"})
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
	// convert user id to int
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}
	//convert user id to int
	userIDInt, err := strconv.Atoi(userID.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// check if user is following himself
	if userIDInt == follow.FollowedID {
		c.JSON(400, gin.H{"error": "You can't unfollow yourself"})
		return
	}

	// check if user is not following
	isFollowing, err := h.services.IsFollowing(c, userIDInt, follow.FollowedID)
	if err != nil {
		c.JSON(400, gin.H{"error":"error checking if user is following"})
		return
	}

	if !isFollowing {
		c.JSON(400, gin.H{"error": "You are not following this user"})
		return
	}

	// create unfollow
	err = h.services.UnFollowUser(c, userIDInt, follow.FollowedID)

	// check error
	if err != nil {
		c.JSON(400, gin.H{"error":"error unfollowing user"})
		return
	}

	// return result if no error
	c.JSON(201, gin.H{"message": "User unfollowed successfully"})
}

func (h *Handler) GetFollowers(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// convert user id to int
	userIDInt, err := strconv.Atoi(userID.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// create follow
	result, err := h.services.GetFollowers(c, userIDInt)

	// check error
	if err != nil {
		c.JSON(400, gin.H{"error":"error getting followers"})
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

	// convert user id to int
	userIDInt, err := strconv.Atoi(userID.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// create follow
	result, err := h.services.GetFollowing(c, userIDInt)

	// check error
	if err != nil {
		c.JSON(400, gin.H{"error":"error getting following"})
		return
	}

	// return result if no error
	c.JSON(200, gin.H{"result": result})
}

func (h *Handler) GetFollowersByUserID(c *gin.Context) {
	//get followers user id
	follow_userID := c.Param("user_id")
	// convert user id to int
	followUserIDInt, err := strconv.Atoi(follow_userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	//convert user id to int
	userIDInt, err := strconv.Atoi(userID.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// check if user is following himself
	if userIDInt == followUserIDInt {
		c.JSON(400, gin.H{"error": "You can't track yourself with id"})
		return
	}

	// check if user is already following
	isFollowing, err := h.services.IsFollowing(c, userIDInt, followUserIDInt)
	if err != nil {
		c.JSON(400, gin.H{"error":"error checking if user is following"})
		return
	}

	if !isFollowing {
		c.JSON(400, gin.H{"error": "You are not following this user "})
		return
	}

	// create follow
	result, err := h.services.GetFollowers(c, followUserIDInt)

	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return result if no error
	c.JSON(200, gin.H{"result": result})
}

func (h *Handler) GetFollowingByUserID(c *gin.Context) {
	//get followers user id
	follow_userID := c.Param("user_id")

	// convert user id to int
	followUserIDInt, err := strconv.Atoi(follow_userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	//convert user id to int
	userIDInt, err := strconv.Atoi(userID.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// check if user is following himself
	if userIDInt == followUserIDInt {
		c.JSON(400, gin.H{"error": "You can't track yourself with id"})
		return
	}

	// check if user is already following
	isFollowing, err := h.services.IsFollowing(c, userIDInt, followUserIDInt)
	if err != nil {
		c.JSON(400, gin.H{"error":"error checking if user is following"})
		return
	}

	if !isFollowing {
		c.JSON(400, gin.H{"error": "You are not following this user "})
		return
	}
	// create follow
	result, err := h.services.GetFollowing(c, followUserIDInt)

	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return result if no error
	c.JSON(200, gin.H{"result": result})
}

func (h *Handler) Search(c *gin.Context) {
	//get query
	query := c.Query("q")

	// create search
	result, err := h.services.Search(c, query)

	// check error
	if err != nil {
		c.JSON(400, gin.H{"error":"error searching"})
		return
	}

	// return result if no error
	c.JSON(200, gin.H{"result": result})
}

func (h *Handler) UploadProfileImage(c *gin.Context) {
	// Parse form file
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "No file is received"})
		return
	}
	defer file.Close()

	// Upload to S3
	fileURL, err := uploadToS3(file, fileHeader, users_folder)
	if err != nil {
		c.JSON(500, gin.H{"error":"error uploading file"})
		return
	}
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// upload file
	err = h.services.UploadProfileImage(c, userID.(string), fileURL)
	if err != nil {
		c.JSON(400, gin.H{"error":"error uploading profile image"})
		return
	}

	// return result if no error
	c.JSON(200, gin.H{"url": fileURL})
}