package handler

import (
	"errors"

	"github.com/jasurbek-suyunov/udevs_project/helper"
	"github.com/jasurbek-suyunov/udevs_project/models"

	"github.com/gin-gonic/gin"
)

const (
	twits_folder = "twits"
)

func (h *Handler) CreateTwit(c *gin.Context) {

	// Parse form file
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "No file is received"})
		return
	}
	defer file.Close()

	// Get twit
	twit := c.PostForm("twit")

	// Upload to S3
	fileURL, err := uploadToS3(file, fileHeader, twits_folder)
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

	var twitModel = models.Twit{
		UserID:    userID.(string),
		Content:   twit,
		MediaUrl:  fileURL,
		CreatedAt: helper.GetCurrentTime(),
	}

	// create twit
	createdTwit, err := h.services.CreateTwit(c, &twitModel)
	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": "error creating twit"})
		return
	}

	// return result if no error
	c.JSON(201, createdTwit)
}

func (h *Handler) GetTwitsByUserID(c *gin.Context) {
	//get user id
	userID := c.Param("id")

	// check user id
	ok := helper.CheckIntegers(userID)
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get twits
	twits, err := h.services.GetTwitsByUserID(c, userID)
	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": "error getting twits"})
		return
	}

	// return result if no error
	c.JSON(200, twits)
}

func (h *Handler) DeleteTwit(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get twit id
	twitID := c.Param("id")
	if len(twitID) == 0 {
		c.JSON(400, gin.H{"error": errors.New("invalid twit id")})
		return
	}
	// delete twit
	err := h.services.DeleteTwit(c, userID.(string), twitID)
	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": "error deleting twit"})
		return
	}

	// return result if no error
	c.JSON(200, gin.H{"message": "twit deleted"})
}

func (h *Handler) UpdateTwit(c *gin.Context) {
	// variable
	var twit models.TwitRequest

	// bind
	if err := c.ShouldBindJSON(&twit); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get twit id
	twitID := c.Param("id")

	var twitModel = models.Twit{
		ID:        twitID,
		UserID:    userID.(string),
		Content:   twit.Content,
		CreatedAt: helper.GetCurrentTime(),
	}

	// update twit
	updatedTwit, err := h.services.UpdateTwit(c, &twitModel)
	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": "error updating twit"})
		return
	}

	// return result if no error
	c.JSON(200, updatedTwit)
}

func (h *Handler) GetTwitByID(c *gin.Context) {
	//get twit id
	twitID := c.Param("id")

	// check twit id
	ok := helper.CheckIntegers(twitID)
	if !ok {
		c.JSON(400, gin.H{"error": "twit_id not found"})
		return
	}

	// get twit
	twit, err := h.services.GetTwitByID(c, twitID)
	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": "error getting twit"})
		return
	}

	// return result if no error
	c.JSON(200, twit)
}

func (h *Handler) GetTwits(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get twits
	twits, err := h.services.GetTwits(c, userID.(string))
	// check error
	if err != nil {
		c.JSON(400, gin.H{"error": "error getting twits"})
		return
	}

	// return result if no error
	c.JSON(200, twits)
}

func (h *Handler) LikeTwit(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get twit id
	twitID := c.Param("twit_id")

	// check twit id
	ok = helper.CheckIntegers(twitID)
	if !ok {
		c.JSON(400, gin.H{"error": "twit_id not found"})
		return
	}

	err := h.services.LikeTwit(c, userID.(string), twitID)
	if err != nil {
		c.JSON(400, gin.H{"error": "error saving like"})
		return
	}

	c.JSON(200, gin.H{"message": "like saved"})
}

func (h *Handler) UnLikeTwit(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get twit id
	twitID := c.Param("twit_id")
	// check twit id
	ok =helper.CheckIntegers(twitID)
	if !ok {
		c.JSON(400, gin.H{"error": "twit_id not found"})
		return
	}

	err := h.services.UnLikeTwit(c, userID.(string), twitID)
	if err != nil {
		c.JSON(400, gin.H{"error": "error saving unlike"})
		return
	}

	c.JSON(200, gin.H{"message": "unlike saved"})
}

func (h *Handler) RetwitTwit(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get twit id
	twitID := c.Param("twit_id")
	// check twit id
	ok =helper.CheckIntegers(twitID)
	if !ok {
		c.JSON(400, gin.H{"error": "twit_id not found"})
		return
	}
	err := h.services.RetwitTwit(c, userID.(string), twitID)
	if err != nil {
		c.JSON(400, gin.H{"error": "error saving retwit"})
		return
	}

	c.JSON(200, gin.H{"message": "retwit saved"})
}

func (h *Handler) UnRetwitTwit(c *gin.Context) {
	//get user id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{"error": "user_id not found"})
		return
	}

	// get twit id
	twitID := c.Param("twit_id")
	// check twit id
	ok =helper.CheckIntegers(twitID)
	if !ok {
		c.JSON(400, gin.H{"error": "twit_id not found"})
		return
	}


	err := h.services.UnRetwitTwit(c, userID.(string), twitID)
	if err != nil {
		c.JSON(400, gin.H{"error": "error saving unretwit"})
		return
	}

	c.JSON(200, gin.H{"message": "unretwit saved"})
}
