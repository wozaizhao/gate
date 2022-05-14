package models

import (
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// 前端库
type FeRepo struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	OwnerName       string         `json:"ownerName" gorm:"type:varchar(20);NOT NULL;comment:所有者名称"`
	RepoName        string         `json:"repoName" gorm:"type:varchar(20);NOT NULL;comment:仓库名称"`
	RepoDesc        string         `json:"repoDesc" gorm:"type:varchar(80);NOT NULL;comment:仓库描述"`
	Language        string         `json:"language" gorm:"type:varchar(20);NOT NULL;comment:语言"`
	OwnerAvatarURL  string         `json:"ownerAvatarURL" gorm:"type:varchar(100);NOT NULL;comment:所有者头像"`
	HomePage        string         `json:"homePage" gorm:"type:varchar(50);NOT NULL;comment:主页"`
	GithubURL       string         `json:"githubURL" gorm:"type:varchar(255);comment:Github网址"`
	License         string         `json:"license" gorm:"type:varchar(20);NOT NULL;comment:许可证"`
	Topics          string         `json:"topics" gorm:"type:varchar(40);NOT NULL;comment:话题"`
	LikeCount       uint           `json:"likeCount" gorm:"comment:点赞数"`
	ViewCount       uint           `json:"viewCount" gorm:"comment:浏览数"`
	GithubLikeCount uint           `json:"githubLikeCount" gorm:"comment:Github点赞数"`
	Status          uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}

func CreateFeRepo(ownerName, repoName, repoDesc, language, ownerAvatarURL, homePage, githubURL, license, topics string, githubLikeCount uint) error {
	repo := FeRepo{OwnerName: ownerName, RepoName: repoName, RepoDesc: repoDesc, Language: language, OwnerAvatarURL: ownerAvatarURL, HomePage: homePage, GithubURL: githubURL, License: license, Topics: topics, GithubLikeCount: githubLikeCount, Status: common.STATUS_NORMAL}
	err := DB.Create(&repo).Error
	return err
}

func IsRepoExist(ownerName, repoName string) bool {
	var repo FeRepo
	DB.Where("owner_name = ? AND repo_name = ?", ownerName, repoName).First(&repo)
	return repo.ID > 0
}

func GetFeRepos(pageNum, pageSize int) ([]FeRepo, error) {
	var repos []FeRepo
	err := DB.Scopes(Paginate(pageNum, pageSize), FieldEqual("status", common.STATUS_NORMAL)).Order("github_like_count desc").Find(&repos).Error
	return repos, err
}
