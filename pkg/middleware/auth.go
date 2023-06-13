package middleware

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const SK = "secret"

type ClaimsWithScope struct {
	jwt.StandardClaims
	Scope string
}

func GetUserID(c *gin.Context) (uint, error) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return 0, err
	}
	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SK), nil
	})
	if err != nil {
		return 0, err
	}

	payload := token.Claims.(*ClaimsWithScope)
	id, _ := strconv.Atoi(payload.Subject)
	return uint(id), nil
}
