package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type UserLoginWithWechat struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	OpenID    string         `json:"open_id" gorm:"unique;type:varchar(40);DEFAULT ''"` // openID 小程序登录获取
	UserID    uint           `json:"userID"`
}

// 创建或更新用户使用微信登录记录
func UpsertUserLoginWithWechat(openID string, userID uint) (err error) {
	// 如果openID的记录存在，就更新；否则创建
	var record UserLoginWithWechat
	record.OpenID = openID
	record.UserID = userID
	r := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "open_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"user_id"}),
	}).Create(&record)
	err = r.Error
	return
}

func GetLoginRecordByOpenID(openID string) (record *UserLoginWithWechat, exist bool, err error) {
	r := DB.Where("open_id = ?", openID).Find(&record)
	err = r.Error
	exist = r.RowsAffected > 0
	return
}
