package config

import (
	"flag"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	cfgFile = flag.String("config", "./config.yaml", "配置文件路径")

	cfg *Config
)

// Config example config
type Config struct {
	Listen string `yaml:"listen"`
	Mode   string `yaml:"mode"`
	Redis  struct {
		Host        string `yaml:"host"`
		Password    string `yaml:"password"`
		Database    int    `yaml:"database"`
		MaxActive   int    `yaml:"maxActive"`
		MaxIdle     int    `yaml:"maxIdle"`
		IdleTimeout int    `yaml:"idleTimeout"`
	} `yaml:"redis"`
	Mysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`
	JwtSecret              string `yaml:"jwtsecret"`
	*OfficialAccountConfig `yaml:"officialAccountConfig"`
	*MiniProgramConfig     `yaml:"miniProgramConfig"`
	*AliyunConfig          `yaml:"aliyunConfig"`
	*GeetestConfig         `yaml:"geetestConfig"`
}

// OfficialAccountConfig 公众号相关配置
type OfficialAccountConfig struct {
	AppID          string `yaml:"appID"`
	AppSecret      string `yaml:"appSecret"`
	Token          string `yaml:"token"`
	EncodingAESKey string `yaml:"encodingAESKey"`
}

// MiniProgramConfig 小程序相关配置
type MiniProgramConfig struct {
	AppID     string `yaml:"appID"`
	AppSecret string `yaml:"appSecret"`
}

// AliyunConfig 阿里云相关配置
type AliyunConfig struct {
	AccessKeyID          string `yaml:"accessKeyID"`
	AccessKeySecret      string `yaml:"accessKeySecret"`
	Endpoint             string `yaml:"endpoint"`
	SignName             string `yaml:"signName"`
	LoginTemplateCode    string `yaml:"loginTemplateCode"`
	RegisterTemplateCode string `yaml:"registerTemplateCode"`
}

// 极验相关配置
type GeetestConfig struct {
	GeetestID  string `yaml:"geetestID"`
	GeetestKey string `yaml:"geetestKey"`
}

// GetConfig 获取配置
func GetConfig() *Config {
	if cfg != nil {
		return cfg
	}
	bytes, err := ioutil.ReadFile(*cfgFile)
	if err != nil {
		panic(err)
	}

	cfgData := &Config{}
	err = yaml.Unmarshal(bytes, cfgData)
	if err != nil {
		panic(err)
	}
	cfg = cfgData
	return cfg
}
