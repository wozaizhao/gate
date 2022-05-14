package models

import (
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// 分类
type FeCate struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CateName  string         `json:"cateName" gorm:"type:varchar(20);NOT NULL;comment:分类名称"`
	CateDesc  string         `json:"cateDesc" gorm:"type:varchar(20);NOT NULL;comment:分类描述"`
	Priority  uint           `json:"priority" gorm:"NOT NULL;comment:优先级"`
	Status    uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}

func CreateFeCate(cateName, cateDesc string) (err error) {
	feCate := FeCate{CateName: cateName, CateDesc: cateDesc, Priority: common.PRIORITY_NORMAL, Status: common.STATUS_NORMAL}
	err = DB.Create(&feCate).Error
	return
}

func GetFeCates() (feCates []FeCate, err error) {
	err = DB.Scopes(FieldEqual("status", common.STATUS_NORMAL)).Find(&feCates).Order("priority desc, createdAt desc").Error
	return
}
