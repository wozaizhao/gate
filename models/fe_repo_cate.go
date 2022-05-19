package models

import (
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// 库分类
type FeRepoCate struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	CateName     string         `json:"cate_name" gorm:"type:varchar(30);NOT NULL;comment:分类名称"`
	CateCNName   string         `json:"cate_cn_name" gorm:"type:varchar(20);NOT NULL;comment:分类中文名称"`
	CateDesc     string         `json:"cate_desc" gorm:"type:varchar(255);NOT NULL;comment:分类描述"`
	CateParentID uint           `json:"cate_parent_id" gorm:"type:int(11);NOT NULL;DEFAULT '0';comment:父级分类ID"`
	Status       uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}

func CreateFeRepoCate(cateName, cateCNName, cateDesc string, cateParentID uint) (cate FeRepoCate, err error) {
	cate = FeRepoCate{CateName: cateName, CateCNName: cateCNName, CateDesc: cateDesc, CateParentID: cateParentID, Status: common.STATUS_NORMAL}
	result := DB.Create(&cate)
	err = result.Error
	return
}

func GetFeRepoCates() (cates []FeRepoCate, err error) {
	err = DB.Where("status = ?", common.STATUS_NORMAL).Find(&cates).Error
	return
}

func GetFirstClassRepoCates() (cates []FeRepoCate, err error) {
	err = DB.Where("status = ? AND cate_parent_id = ?", common.STATUS_NORMAL, 0).Find(&cates).Error
	return
}
