package controllers

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	// "wozaizhao.com/gate/config"
	"wozaizhao.com/gate/common"
	"wozaizhao.com/gate/middlewares"
)

type uploadForm struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func Upload(c *gin.Context) {
	// cfg := config.GetConfig()

	var form uploadForm
	if err := c.ShouldBind(&form); err != nil {
		RenderBadRequest(c, err)
		return
	}
	file, _ := c.FormFile("file")

	// if cfg.Mode != "production" {
	// 	RenderSuccess(c, "Fk0nSk_9_ck8vGdrt-fw6kgurGr2", "")
	// 	return
	// }

	ret, err := middlewares.Upload(file)
	if err != nil {
		common.LogError("upload", err)
		RenderError(c, err)
		return
	}
	RenderSuccess(c, ret, "")
	// etag, _ := helpers.GetEtag(buf)
}
