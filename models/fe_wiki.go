package models

import (
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// 前端百科
type FeWiki struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	WikiName    string         `json:"wikiName" gorm:"type:varchar(20);NOT NULL;comment:百科名称"`
	WikiDesc    string         `json:"wikiDesc" gorm:"type:varchar(20);NOT NULL;comment:百科描述"`
	WikiContent string         `json:"wikiContent" gorm:"type:varchar(20);NOT NULL;comment:百科内容"`
	LikeCount   uint           `json:"likeCount" gorm:"comment:点赞数"`
	ViewCount   uint           `json:"viewCount" gorm:"comment:浏览数"`
	Priority    uint           `json:"priority" gorm:"NOT NULL;comment:优先级"`
	Status      uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}

func CreateFeWiki(wikiName, wikiDesc, wikiContent string) (err error) {
	feWiki := FeWiki{WikiName: wikiName, WikiDesc: wikiDesc, WikiContent: wikiContent, Priority: common.PRIORITY_NORMAL, Status: common.STATUS_NORMAL}
	err = DB.Create(&feWiki).Error
	return
}

func GetFeWikis(pageNum, pageSize int) (feWikis []FeWiki, err error) {
	err = DB.Scopes(Paginate(pageNum, pageSize), FieldEqual("status", common.STATUS_NORMAL)).Find(&feWikis).Order("priority desc, createdAt desc").Error
	return
}
