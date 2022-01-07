package middlewares

import (
	"encoding/json"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"wozaizhao.com/gate/config"
)

var client *dysmsapi20170525.Client

/**
 * 使用AK&SK初始化账号Client
 * @throws Exception
 */
func InitSmsClient() (err error) {
	cfg := config.GetConfig()
	openConfig := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &cfg.AliyunConfig.AccessKeyID,
		// 您的AccessKey Secret
		AccessKeySecret: &cfg.AliyunConfig.AccessKeySecret,
	}
	openConfig.Endpoint = tea.String(cfg.AliyunConfig.Endpoint)
	client, err = dysmsapi20170525.NewClient(openConfig)
	return err
}

type codeStruct struct {
	Code string `json:"code"`
}

var loginTemplateCode = config.GetConfig().AliyunConfig.LoginTemplateCode

func SendLoginSms(phone, code string) error {
	err := send(phone, loginTemplateCode, code)
	return err
}

func send(phone, templateCode, code string) (err error) {
	cfg := config.GetConfig()
	var registerCode = codeStruct{Code: code}
	jsonString, err := json.Marshal(registerCode)
	if err != nil {
		return err
	}
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(cfg.AliyunConfig.SignName),
		TemplateCode:  tea.String(templateCode),
		TemplateParam: tea.String(string(jsonString)), // "{\"code\": 333333}"
	}
	_, err = client.SendSms(sendSmsRequest)
	if err != nil {
		return err
	}
	return nil
}
