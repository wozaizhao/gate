package main

import (
	"wozaizhao.com/gate/common"
	"wozaizhao.com/gate/config"
	"wozaizhao.com/gate/middlewares"
	"wozaizhao.com/gate/models"
	"wozaizhao.com/gate/router"
)

func main() {

	r := router.SetupRouter()
	cfg := config.GetConfig()

	// if cfg.Mode == "production" {
	common.LogToFile()
	// }

	middlewares.InitWechat()
	middlewares.InitSmsClient()

	models.DBinit()
	r.Run(cfg.Listen)
}
