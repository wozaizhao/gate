package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	// "wozaizhao.com/eden/config"
	"wozaizhao.com/gate/common"
	"wozaizhao.com/gate/middlewares"
	"wozaizhao.com/gate/models"
)

type AuthHeader struct {
	Token string `header:"Authorization"`
}

func TokenValidator(c *gin.Context) (*middlewares.Claims, error) {
	h := AuthHeader{}

	if err := c.ShouldBindHeader(&h); err != nil {
		return nil, err
	}
	var token string
	if len(strings.Split(h.Token, " ")) == 2 {
		token = strings.Split(h.Token, " ")[1]
	} else {
		return nil, errors.New("token is not found")
	}
	claims, errorParseToken := middlewares.ParseToken(token)
	if errorParseToken != nil {
		return nil, errorParseToken
	}
	return claims, nil
}

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, errorTokenValidator := TokenValidator(c)
		if errorTokenValidator != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		userStatus := models.GetUserStatus(claims.UserID)
		if common.USER_STATUS_NOT_ACTIVATED == userStatus {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": common.ERROR_USER_NOT_ACTIVATED})
			return

		} else if common.USER_STATUS_FORBIDDEN == userStatus {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": common.ERROR_USER_SUSPENDED})
			return
		}
		c.Set("userID", claims.UserID)
		c.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, errorTokenValidator := TokenValidator(c)
		if errorTokenValidator != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// todo check admin
		c.Set("userID", claims.UserID)
		c.Next()
		// if common.ADMIN_ROLE == models.GetUserRole(claims.UserID) {
		// 	c.Set("userID", claims.UserID)
		// 	c.Next()
		// } else {
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// }
	}

}
