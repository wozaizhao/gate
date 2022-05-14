package models

import (
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// 导航
type FeNav struct {
	ID                  uint           `json:"id" gorm:"primaryKey"`
	CreatedAt           time.Time      `json:"createdAt"`
	UpdatedAt           time.Time      `json:"updatedAt"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	NavTitle            string         `json:"navTitle" gorm:"type:varchar(20);NOT NULL;comment:导航名称"`
	Icon                string         `json:"icon" gorm:"type:varchar(20);NOT NULL;comment:图标"`
	Image               string         `json:"image" gorm:"type:varchar(20);NOT NULL;comment:图片"`
	WhiteBgIconColor    string         `json:"whiteBgIconColor" gorm:"type:varchar(20);NOT NULL;comment:白色背景图标颜色"`
	ColorBgIconColor    string         `json:"colorBgIconColor" gorm:"type:varchar(20);NOT NULL;comment:彩色背景图标颜色"`
	ColorBgColor        string         `json:"colorBgColor" gorm:"type:varchar(20);NOT NULL;comment:彩色背景颜色"`
	GradientBgIconColor string         `json:"gradientBgIconColor" gorm:"type:varchar(20);NOT NULL;comment:渐变背景图标颜色"`
	GradientBgColor     string         `json:"gradientBgColor" gorm:"type:varchar(20);NOT NULL;comment:渐变背景颜色"`
	Priority            uint           `json:"priority" gorm:"NOT NULL;comment:优先级"`
	Status              uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
	NavType             uint           `json:"navType" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:导航类型"`
	NavPath             string         `json:"navPath" gorm:"type:varchar(50);NOT NULL;comment:导航路径"`
	NavParams           string         `json:"navParams" gorm:"type:varchar(50);NOT NULL;comment:导航参数"`
}

func CreateFeNav(navTitle, icon, image, whiteBgIconColor, colorBgIconColor, colorBgColor, gradientBgIconColor, gradientBgColor string, navType uint, navPath, navParams string) (err error) {
	feNav := FeNav{NavTitle: navTitle, Icon: icon, Image: image, WhiteBgIconColor: whiteBgIconColor, ColorBgIconColor: colorBgIconColor, ColorBgColor: colorBgColor, GradientBgIconColor: gradientBgIconColor, GradientBgColor: gradientBgColor, NavType: navType, NavPath: navPath, NavParams: navParams}
	err = DB.Create(&feNav).Error
	return
}

func GetFeNavs() (feNavs []FeNav, err error) {
	err = DB.Scopes(FieldEqual("status", common.STATUS_NORMAL)).Find(&feNavs).Order("priority desc, createdAt desc").Error
	return
}
