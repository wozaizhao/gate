package models

import (
	"time"

	"gorm.io/gorm"
	// "wozaizhao.com/gate/common"
)

// 基础
type FeFundamental struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Cate        string         `json:"cate" gorm:"type:varchar(20);NOT NULL;comment:分类"`
	CateCN      string         `json:"cateCn" gorm:"type:varchar(20);NOT NULL;comment:分类中文名称"`
	Topic       string         `json:"topic" gorm:"type:varchar(20);NOT NULL;comment:主题"`
	TopicCN     string         `json:"topicCn" gorm:"type:varchar(20);NOT NULL;comment:主题中文名称"`
	TopicDesc   string         `json:"topicDesc" gorm:"type:varchar(255);NOT NULL;comment:主题描述"`
	TopicDescCN string         `json:"topicDescCn" gorm:"type:varchar(255);NOT NULL;comment:主题描述中文名称"`
	Links       []FeLink       `json:"links" gorm:"many2many:fe_fundamental_link;"`
	Status      uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}

func CreateFeFundamental(cate, cateCn, topic, topicCn, topicDesc, topticDescCn string, links []FeLink) error {
	fund := FeFundamental{Cate: cate, CateCN: cateCn, Topic: topic, TopicCN: topicCn, TopicDesc: topicDesc, TopicDescCN: topticDescCn, Links: links}
	return DB.Create(&fund).Error
}

func GetFeFundamentals() ([]FeFundamental, error) {
	var fund []FeFundamental
	err := DB.Find(&fund).Error
	return fund, err
}

func AdminGetFundamentals(pageNum, pageSize int) ([]FeFundamental, error) {
	var fund []FeFundamental
	err := DB.Scopes(Paginate(pageNum, pageSize)).Preload("Links").Find(&fund).Error
	return fund, err
}

func GetFeFundamental(id uint) (FeFundamental, error) {
	var fund FeFundamental
	err := DB.First(&fund, id).Error
	return fund, err
}

func GetFundamentalCount() (int64, error) {
	var count int64
	err := DB.Model(&FeFundamental{}).Count(&count).Error
	return count, err
}

func UpdateFeFundamental(id, status uint, cate, cateCn, topic, topicCn, topicDesc, topicDescCn string, links []FeLink) error {
	err := DB.Model(&FeFundamental{}).Where("id = ?", id).Updates(FeFundamental{ID: id, Cate: cate, CateCN: cateCn, Topic: topic, TopicCN: topicCn, TopicDesc: topicDesc, TopicDescCN: topicDescCn, Links: links, Status: status}).Error
	return err
}
