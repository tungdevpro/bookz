package middleware

import (
	"bookz/api/server/helper"
	"bookz/api/server/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const SECRET_KEY = "12345678aA@"

func GenToken(user models.User) (string, error) {
	if len(user.ID) == 0 {
		return "", errors.New("len(id) == 0")
	}

	claims := &models.JwtCustomClaims{
		UserId: "1331231",
		Role:   user.Role,
		Phone:  user.Phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3600).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err := token.SignedString([]byte(SECRET_KEY))
	helper.Fatal(err)

	return result, nil
}

func ParseToken(token string) any {
	if len(token) == 0 {
		return helper.ErrorResponse{
			StatusCode: 0,
			Message:    "",
		}
	}

	return helper.Response{
		StatusCode: 1,
		Message:    "Done",
	}
}
