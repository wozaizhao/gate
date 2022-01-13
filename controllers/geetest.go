package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"wozaizhao.com/gate/common"
	"wozaizhao.com/gate/config"
	"wozaizhao.com/gate/middlewares"
	"wozaizhao.com/gate/models"
)

var cfg = config.GetConfig().GeetestConfig

// geetest 公钥
var CAPTCHA_ID string = cfg.GeetestID

// geetest 密钥
var CAPTCHA_KEY string = cfg.GeetestKey

// geetest 服务地址
const API_SERVER string = "http://gcaptcha4.geetest.com"

// geetest 验证接口
var URL = API_SERVER + "/validate" + "?captcha_id=" + CAPTCHA_ID

// GeetestVerify
func GeetestVerify(c *gin.Context) {

	// 前端传回的数据
	lot_number := c.Query("lot_number")
	captcha_output := c.Query("captcha_output")
	pass_token := c.Query("pass_token")
	gen_time := c.Query("gen_time")
	phone := c.Query("phone")
	// 生成签名
	// 生成签名使用标准的hmac算法，使用用户当前完成验证的流水号lot_number作为原始消息message，使用客户验证私钥作为key
	// 采用sha256散列算法将message和key进行单向散列生成最终的 “sign_token” 签名
	sign_token := hmac_encode(CAPTCHA_KEY, lot_number)

	// 向极验转发前端数据 + “sign_token” 签名
	form_data := make(url.Values)
	form_data["lot_number"] = []string{lot_number}
	form_data["captcha_output"] = []string{captcha_output}
	form_data["pass_token"] = []string{pass_token}
	form_data["gen_time"] = []string{gen_time}
	form_data["sign_token"] = []string{sign_token}

	// 发起post请求
	// 设置5s超时
	cli := http.Client{Timeout: time.Second * 5}
	resp, err := cli.PostForm(URL, form_data)
	if err != nil || resp.StatusCode != 200 {
		// 当请求发生异常时，应放行通过，以免阻塞业务。
		common.LogError("GeetestLogin PostForm", err)
		err = send(phone)
		if err != nil {
			RenderError(c, err)
		}
		RenderSuccess(c, true, "验证码已发送")
		return
	}

	res_json, _ := ioutil.ReadAll(resp.Body)
	var res_map map[string]interface{}
	// 根据极验返回的用户验证状态, 网站主进行自己的业务逻辑
	// 响应json数据如：{"result": "success", "reason": "", "captcha_args": {}}
	if err := json.Unmarshal(res_json, &res_map); err == nil {
		result := res_map["result"]
		if result == "success" {
			err = send(phone)
			if err != nil {
				RenderError(c, err)
			}
			RenderSuccess(c, true, "验证码已发送")
		} else {
			reason := res_map["reason"]
			common.LogError("GeetestVerify fail: ", reason)
			RenderFail(c, fmt.Sprintf("%v", reason))
		}
	}
}

// hmac-sha256 加密：  CAPTCHA_KEY,lot_number
func hmac_encode(key string, data string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

func send(phone string) error {
	cfg := config.GetConfig()
	code, err := models.CreateCaptcha(phone)
	if err != nil {
		return err
	}
	if cfg.Mode != "production" {
		return nil
	}
	err = middlewares.SendLoginSms(phone, code)
	if err != nil {
		return err
	}
	return nil
}
