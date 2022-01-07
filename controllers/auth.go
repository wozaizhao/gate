package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// "wozaizhao.com/eden/config"
	"wozaizhao.com/gate/common"
	"wozaizhao.com/gate/middlewares"
	"wozaizhao.com/gate/models"
)

type AuthHeader struct {
	Token string `header:"token"`
}

func TokenValidator(c *gin.Context) (*middlewares.Claims, error) {
	h := AuthHeader{}

	if err := c.ShouldBindHeader(&h); err != nil {
		return nil, err
	}
	claims, errorParseToken := middlewares.ParseToken(h.Token)
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
		if common.STATUS_NORMAL == models.GetUserStatus(claims.UserID) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 1002, "errMsg": "帐户已禁用！"})
			return
		}
		c.Set("user-id", claims.UserID)
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
		if common.ADMIN_ROLE == models.GetUserRole(claims.UserID) {
			c.Set("user-id", claims.UserID)
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}

}
