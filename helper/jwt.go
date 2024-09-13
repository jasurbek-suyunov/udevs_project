package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"jas/models"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(param *models.Token) string {

	json, _ := json.Marshal(param)
	timeExpired, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRE_TIME"))
	if err != nil {
		log.Printf("failed from convert time expired: %v", err)
		timeExpired = 18000
	}

	unixTime := time.Now().Add(time.Duration(timeExpired * int(time.Second))).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": string(json),
		"exp": unixTime,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		log.Printf("failed from generate token: %v", err)
		return ""
	}

	return tokenString
}

func ValidateJWT(tokenString string) (*models.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, errors.New("failed expired time")
		}
	}

	jsonString := claims["sub"].(string)

	var param models.Token

	err = json.Unmarshal([]byte(jsonString), &param)
	if err != nil {
		return nil, errors.New("failed from unmarshal json")
	}
	return &param, nil
}
