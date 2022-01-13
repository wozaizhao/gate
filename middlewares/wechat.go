package middlewares

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	log "github.com/sirupsen/logrus"
	"wozaizhao.com/gate/config"
)

var (
	wc              *wechat.Wechat
	miniProgram     *miniprogram.MiniProgram
	officialAccount *officialaccount.OfficialAccount
)

func InitWechat() {
	cfg := config.GetConfig()
	wc = wechat.NewWechat()
	redisOpts := &cache.RedisOpts{
		Host:        cfg.Redis.Host,
		Password:    cfg.Redis.Password,
		Database:    cfg.Redis.Database,
		MaxActive:   cfg.Redis.MaxActive,
		MaxIdle:     cfg.Redis.MaxIdle,
		IdleTimeout: cfg.Redis.IdleTimeout,
	}
	redisCache := cache.NewRedis(redisOpts)
	wc.SetCache(redisCache)

	miniCfg := &miniConfig.Config{
		AppID:     cfg.MiniProgramConfig.AppID,
		AppSecret: cfg.MiniProgramConfig.AppSecret,
	}
	log.Debugf("miniCfg=%+v", miniCfg)
	miniProgram = wc.GetMiniProgram(miniCfg)

	offCfg := &offConfig.Config{
		AppID:     cfg.OfficialAccountConfig.AppID,
		AppSecret: cfg.OfficialAccountConfig.AppSecret,
		Token:     cfg.OfficialAccountConfig.Token,
		//EncodingAESKey: "xxxx",
	}
	log.Debugf("offCfg=%+v", offCfg)
	officialAccount = wc.GetOfficialAccount(offCfg)
}

func GetMiniProgram() *miniprogram.MiniProgram {
	return miniProgram
}

func GetOfficialAccount() *officialaccount.OfficialAccount {
	return officialAccount
}
