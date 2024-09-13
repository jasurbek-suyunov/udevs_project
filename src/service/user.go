package service

import (
	"context"
	"encoding/json"
	"fmt"
	"jas/helper"
	"jas/models"
	"time"
)

func (s *Service) CreateUser(ctx context.Context, user *models.UserSignUpRequest) error {
	// check if user already
	// generate password hash
	pass_hash, err := helper.GeneratePasswordHash(user.Password)

	// check error
	if err != nil {
		return err
	}

	// create user
	result, err := s.storage.User().CreateUser(ctx, &models.User{
		Username:     user.Username,
		FullName:     user.FullName,
		Bio:          user.Bio,
		Email:        user.Email,
		PasswordHash: pass_hash,
		CreatedAt:    time.Now().Unix(),
	})

	// check error
	if err != nil {
		return err
	}
	// convert user to json
	user_string, err := json.Marshal(result)

	// check error
	if err != nil {
		return err
	}	
	fmt.Println("user_string", string(user_string))
	fmt.Println("result.Username", result.Username)
	// set username and user
	err = s.cache.Redis().Set(ctx, result.Username, string(user_string), time.Hour*24)

	// check error
	if err != nil {
		return err
	}

	// return result if no error
	return nil
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {

	// result
	var user *models.User

	// get user from redis
	user_string, err := s.cache.Redis().Get(ctx, username)

	// check error
	if err != nil {

		// get user by username from database
		user, err = s.storage.User().GetUserByUsername(ctx, username)

		// check error
		if err != nil {
			return nil, err
		}
	} else {

		err = json.Unmarshal([]byte(user_string), &user)
		if err != nil {
			return nil, err
		}
		err = s.cache.Redis().Delete(ctx, username)
		if err != nil {
			return nil, err
		}
	}

	// return result if no error
	return user, nil
}