package models

import (
	"gorm.io/gorm"
	"time"
)

type UserLoginWithWechat struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	OpenID    string         `json:"open_id" gorm:"type:varchar(40);DEFAULT ''"` // openID 小程序登录获取
	UserID    uint           `json:"userID"`
}

// 创建用户使用微信登录记录
func CreateUserLoginWithWechat(openID string, userID uint) (err error) {
	var record = UserLoginWithWechat{OpenID: openID, UserID: userID}
	result := DB.Create(&record)
	err = result.Error
	return
}

func GetLoginRecordByOpenID(openID string) (record UserLoginWithWechat, err error) {
	r := DB.Where("open_id = ?", openID).Find(&record)
	err = r.Error
	return
}
