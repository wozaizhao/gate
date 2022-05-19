package models

import (
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// 作者或团队
type FeAuthor struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	AuthorName string         `json:"authorName" gorm:"type:varchar(20);NOT NULL;comment:作者姓名"`
	AuthorDesc string         `json:"authorDesc" gorm:"type:varchar(20);NOT NULL;comment:作者描述"`
	WebURL     string         `json:"webURL" gorm:"type:varchar(20);NOT NULL;comment:作者URL"`
	AvatarURL  string         `json:"avatarURL" gorm:"type:varchar(20);NOT NULL;comment:作者头像URL"`
	GithubURL  string         `json:"githubURL" gorm:"type:varchar(20);NOT NULL;comment:作者GithubURL"`
	LikeCount  uint           `json:"likeCount" gorm:"comment:点赞数"`
	ViewCount  uint           `json:"viewCount" gorm:"comment:浏览数"`
	Priority   uint           `json:"priority" gorm:"NOT NULL;comment:优先级"`
	Tags       string         `json:"tags" gorm:"type:varchar(255);NOT NULL;comment:标签"`
	Status     uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}

func CreateFeAuthor(authorName, authorDesc, webURL, avatarURL, gitHubURL string) error {
	author := FeAuthor{AuthorName: authorName, AuthorDesc: authorDesc, WebURL: webURL, AvatarURL: avatarURL, Priority: common.PRIORITY_NORMAL, Status: common.STATUS_NORMAL}
	err := DB.Create(&author).Error
	return err
}

func GetFeAuthorByID(id uint) (*FeAuthor, error) {
	var author FeAuthor
	err := DB.Scopes(FieldEqual("status", common.STATUS_NORMAL)).Where("id = ?", id).First(&author).Error
	return &author, err
}

func GetFeAuthors(pageNum, pageSize int) (feAuthors []FeAuthor, err error) {
	err = DB.Scopes(Paginate(pageNum, pageSize), FieldEqual("status", common.STATUS_NORMAL)).Find(&feAuthors).Order("priority desc, createdAt desc").Error
	return
}
