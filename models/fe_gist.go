package models

import (
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// 代码片断
type FeGist struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	GistName    string         `json:"gistName" gorm:"type:varchar(20);NOT NULL;comment:代码片断名称"`
	GistDesc    string         `json:"gistDesc" gorm:"type:varchar(20);NOT NULL;comment:代码片断描述"`
	GistContent string         `json:"gistContent" gorm:"type:varchar(20);NOT NULL;comment:代码片断内容"`
	Priority    uint           `json:"priority" gorm:"NOT NULL;comment:优先级"`
	LikeCount   uint           `json:"likeCount" gorm:"comment:点赞数"`
	ViewCount   uint           `json:"viewCount" gorm:"comment:浏览数"`
	Status      uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}

func CreateFeGist(gistName, gistDesc, gistContent string) (err error) {
	feGist := FeGist{GistName: gistName, GistDesc: gistDesc, GistContent: gistContent, Priority: common.PRIORITY_NORMAL, Status: common.STATUS_NORMAL}
	err = DB.Create(&feGist).Error
	return
}

func GetFeGists(pageNum, pageSize int) (feGists []FeGist, err error) {
	err = DB.Scopes(Paginate(pageNum, pageSize), FieldEqual("status", common.STATUS_NORMAL)).Find(&feGists).Order("priority desc, createdAt desc").Error
	return
}
