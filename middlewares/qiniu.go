package middlewares

import (
	"context"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"wozaizhao.com/gate/common"
	"wozaizhao.com/gate/config"
)

func Upload(file *multipart.FileHeader) (ret storage.PutRet, err error) {
	qiniuConfig := config.GetConfig().QiniuConfig
	f, err := file.Open()
	if err != nil {
		common.LogError("file.Open", err)
		return
	}

	defer f.Close()

	putPolicy := storage.PutPolicy{
		Scope: qiniuConfig.Bucket,
	}
	mac := qbox.NewMac(qiniuConfig.AccessKey, qiniuConfig.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	// ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	if err != nil {
		return
	}
	err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, io.Reader(f), file.Size, &putExtra)
	return
}
