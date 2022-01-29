package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
	"wozaizhao.com/gate/common"
	"wozaizhao.com/gate/config"
	"wozaizhao.com/gate/controllers"
	"wozaizhao.com/gate/middlewares"
)

// 腾讯验证码 服务地址
const wechatVerifyURL string = "https://api.weixin.qq.com/wxa/checkverificationcode"

func TencentCaptcha(c *gin.Context) {
	mini := middlewares.GetMiniProgram()
	auth := mini.GetAuth()
	accessToken, err := auth.GetAccessToken()
	if err != nil {
		controllers.RenderError(c, err)
		return
	}
	var fullURL = wechatVerifyURL + "?access_token=" + accessToken
	common.LogDebug("fullURL", fullURL)
	cfg := config.GetConfig()
	var userip string = c.ClientIP()
	common.LogDebug("userip", userip)
	var message string
	if cfg.Mode != "production" {
		message = "开发模式下可输入任意验证码"
	} else {
		message = "验证码已发送"
	}

	// 前端传来的数据
	openID := c.Query("openID")
	ticket := c.Query("ticket")
	randstr := c.Query("randstr")
	phone := c.Query("phone")

	type Request struct {
		OpenID  string `json:"openid"`
		Ticket  string `json:"ticket"`
		Randstr string `json:"randstr"`
		UserIP  string `json:"userip"`
	}

	req := Request{
		OpenID:  openID,
		Ticket:  ticket,
		Randstr: randstr,
		UserIP:  userip,
	}

	data, _ := json.Marshal(req)
	// 发起post请求
	// 设置5s超时
	cli := http.Client{Timeout: time.Second * 5}
	resp, err := cli.Post(fullURL, "application/json", bytes.NewBuffer(data))
	if err != nil || resp.StatusCode != 200 {
		// 当请求发生异常时，应放行通过，以免阻塞业务。
		common.LogError("TencentCaptcha Post: ", err)
		err = controllers.Send(phone)
		if err != nil {
			controllers.RenderError(c, err)
			return
		}
		controllers.RenderSuccess(c, true, message)
		return
	}

	defer resp.Body.Close()

	res_json, _ := ioutil.ReadAll(resp.Body)
	var res_map map[string]interface{}
	// 响应json数据如：{"errcode": 0, "errmsg": "ok", "response":1, "evil_level":70, "err_msg":""}
	if err := json.Unmarshal(res_json, &res_map); err == nil {
		response := res_map["response"]
		if response == 1 {
			err = controllers.Send(phone)
			if err != nil {
				controllers.RenderError(c, err)
				return
			}
			controllers.RenderSuccess(c, true, message)
		} else {
			reason := res_map["errmsg"]
			common.LogError("TencentCaptcha", reason)
			controllers.RenderFail(c, fmt.Sprintf("%v", reason))
		}
	}
}
