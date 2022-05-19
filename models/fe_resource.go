package models

import (
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// 资源
type FeResource struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	ResourceName      string         `json:"resource_name" gorm:"type:varchar(20);NOT NULL;comment:资源名称"`
	CateID            uint           `json:"cateID" gorm:"index;NOT NULL;comment:分类ID"`
	AuthorID          uint           `json:"authorID" gorm:"index;comment:作者ID"`
	Thumb             string         `json:"thumb" gorm:"type:varchar(255);NOT NULL;comment:缩略图"`
	Description       string         `json:"description" gorm:"type:varchar(255);NOT NULL;comment:描述"`
	Priority          uint           `json:"priority" gorm:"NOT NULL;comment:优先级"`
	WebURL            string         `json:"webURL" gorm:"type:varchar(255);comment:网址"`
	DocURL            string         `json:"docURL" gorm:"type:varchar(255);comment:文档网址"`
	DownloadURL       string         `json:"downloadURL" gorm:"type:varchar(255);comment:下载网址"`
	GithubURL         string         `json:"githubURL" gorm:"type:varchar(255);comment:Github网址"`
	MiniAppID         string         `json:"miniAppID" gorm:"type:varchar(255);comment:小程序ID"`
	OfficialAccountID string         `json:"officialAccountID" gorm:"type:varchar(255);comment:公众号ID"`
	LikeCount         uint           `json:"likeCount" gorm:"comment:点赞数"`
	ViewCount         uint           `json:"viewCount" gorm:"comment:浏览数"`
	GithubLikeCount   uint           `json:"githubLikeCount" gorm:"comment:Github点赞数"`
	Tags              string         `json:"tags" gorm:"type:varchar(255);comment:标签"`
	Status            uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}

func CreateFeResource(resourceName, thumb, description, webURL, docURL, githubURL, miniAppID, officialAccountID string, cateID, authorID uint) error {
	resource := FeResource{ResourceName: resourceName, CateID: cateID, AuthorID: authorID, Thumb: thumb, Description: description, WebURL: webURL, DocURL: docURL, GithubURL: githubURL, MiniAppID: miniAppID, OfficialAccountID: officialAccountID, Priority: common.PRIORITY_NORMAL, Status: common.STATUS_NORMAL}
	err := DB.Create(&resource).Error
	return err
}

func GetFeResourceByID(id uint) (*FeResource, error) {
	var resource FeResource
	err := DB.Scopes(FieldEqual("status", common.STATUS_NORMAL)).Where("id = ?", id).First(&resource).Error
	return &resource, err
}

func GetFeResourceByAuthorID(id uint) (*[]FeResource, error) {
	var resources []FeResource
	err := DB.Scopes(FieldEqual("status", common.STATUS_NORMAL)).Where("author_id = ?", id).Find(&resources).Group("cate_id").Error
	return &resources, err
}

func GetFeResourceByCateID(pageNum, pageSize int, cateID uint) (*[]FeResource, error) {
	var resources []FeResource
	err := DB.Where("cate_id = ?", cateID).Scopes(Paginate(pageNum, pageSize), FieldEqual("status", common.STATUS_NORMAL)).Find(&resources).Order("priority desc, createdAt desc").Error
	return &resources, err
}
