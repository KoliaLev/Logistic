package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

const secret = "secret"

func CreateToken(userId int, email string) string {
	expireTime := time.Duration(60 * 60 * 1000000000) // 60 sec * 60 min * 1 000 000 000 nanosec
	tokenExpires := time.Now().Add(expireTime).Unix()
	tokenClaims := jwt.MapClaims{}
	tokenClaims["userId"] = userId
	tokenClaims["email"] = email
	tokenClaims["exp"] = tokenExpires

	logrus.Info("tokenClaims ", tokenClaims)

	at := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		logrus.Info("create token unsuccesful")
	}
	return token
}

func TokenIsValid(myToken string) bool {
	token, err := VerifyToken(myToken)
	if err != nil {
		return false
	}
	if err := token.Claims.Valid(); err != nil {
		return false
	}
	return true
}

// verify token
func VerifyToken(myToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
