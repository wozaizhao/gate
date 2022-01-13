package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IDInUri struct {
	ID string `uri:"id" binding:"required"`
}

func RenderError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, fmt.Sprintf("%+v", err.Error()))
}

func RenderBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, fmt.Sprintf("%+v", err.Error()))
}

func RenderFail(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{"status": "fail", "message": message})
}

func RenderSuccess(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": data, "message": message})
}
