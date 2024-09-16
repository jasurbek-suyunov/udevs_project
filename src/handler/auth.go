package handler

import (
	"net/http"

	"github.com/jasurbek-suyunov/udevs_project/helper"
	"github.com/jasurbek-suyunov/udevs_project/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	// variable
	var user models.UserSignUpRequest

	// bind
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//checking password and confirm password
	if user.Password != user.ConfirmPassword {
		c.JSON(400, models.Error{
			Error: "password and confirm password does not match",
		})
		return
	}

	// create user
	err := h.services.CreateUser(c, &user)

	// check error
	if err != nil {
		c.JSON(400, models.Error{
			Error: "error creating user",
		})
		return
	}

	// return result if no error
	c.JSON(201, models.Message{
		Message: "User created successfully",
	})
}

func (h *Handler) SignIn(c *gin.Context) {

	// variable
	var login models.UserLoginRequest

	// bind
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, models.Error{
			Error: err.Error(),
		})
		return
	}

	// get user by username
	user, err := h.services.GetUserByUsername(c, login.Username)

	// check error
	if err != nil {
		c.JSON(404, models.Error{
			Error: "invalid username",
		})
		return
	}
	// check password
	if !helper.CheckPassword(user.PasswordHash, login.Password) {
		c.JSON(401, models.Error{
			Error: "invalid password",
		})
		return
	}

	// generate token
	param := &models.Token{
		UserId:    user.ID,
		UserAgent: c.Request.UserAgent(),
	}
	token := helper.GenerateJWT(param)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", token, 36000, "", "", false, false)
	c.JSON(http.StatusOK, &models.LoginResponse{
		Data: &models.UserResponse{
			ID:              user.ID,
			Username:        user.Username,
			FullName:        user.FullName,
			Bio:             user.Bio,
			Email:           user.Email,
			ProfileImageURL: user.ProfileImageURL,
			CreatedAt:       user.CreatedAt,
		},
		Error: "",
		Code:  0,
	})
}

func (h *Handler) SignOut(c *gin.Context) {
	c.SetCookie("token", "", 0, "", "", false, false)

	c.JSON(http.StatusOK, models.DefaultResponse{
		Data:  "succes logout",
		Error: "",
		Code:  200,
	})
}
