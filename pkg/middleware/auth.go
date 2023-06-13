package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const SK = "secret"

type ClaimsWithScope struct {
	jwt.StandardClaims
	Scope string
}

func GetUserID(c *gin.Context) (uint, error) {
	cookie, err := c.Cookie("user")
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

func IsAuthorized(c *gin.Context) {
	tokenString, err := c.Cookie("user")
	if err != nil {
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, "cookie empty")
			return

		}
	}

	token, err := jwt.ParseWithClaims(tokenString, &ClaimsWithScope{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SK), nil
	})
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Unauthorized")
		return

	}

	payload := token.Claims.(*ClaimsWithScope)
	IsAdmin := strings.Contains(c.FullPath(), "/user")
	if payload.Scope == "admin" && IsAdmin || payload.Scope == "user" && !IsAdmin {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Unauthorized")
		return

	}
	c.Next()
}
