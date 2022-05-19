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
	DeletedAt       gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	OwnerName       string         `json:"ownerName" gorm:"type:varchar(20);NOT NULL;comment:所有者名称"`
	RepoName        string         `json:"repoName" gorm:"type:varchar(50);NOT NULL;comment:仓库名称"`
	RepoDesc        string         `json:"repoDesc" gorm:"type:varchar(500);NOT NULL;comment:仓库描述"`
	Language        string         `json:"language" gorm:"type:varchar(20);NOT NULL;comment:语言"`
	OwnerAvatarURL  string         `json:"ownerAvatarURL" gorm:"type:varchar(100);NOT NULL;comment:所有者头像"`
	HomePage        string         `json:"homePage" gorm:"type:varchar(100);comment:主页"`
	GithubURL       string         `json:"githubURL" gorm:"type:varchar(255);comment:Github网址"`
	License         string         `json:"license" gorm:"type:varchar(20);comment:许可证"`
	Topics          string         `json:"topics" gorm:"type:varchar(500);comment:话题"`
	LikeCount       uint           `json:"likeCount" gorm:"comment:点赞数"`
	ViewCount       uint           `json:"viewCount" gorm:"comment:浏览数"`
	GithubID        uint           `json:"githubID" gorm:"comment:GithubID"`
	GithubLikeCount uint           `json:"githubLikeCount" gorm:"comment:Github点赞数"`
	GithubCreatedAt *time.Time     `json:"githubCreatedAt" gorm:"comment:Github创建时间"`
	GithubUpdatedAt *time.Time     `json:"githubUpdatedAt" gorm:"comment:Github更新时间"`
	Status          uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
	ChineseHomePage string         `json:"chineseHomePage" gorm:"type:varchar(100);comment:中文主页"`
	ChineseTags     string         `json:"tags" gorm:"type:varchar(255);comment:标签"`
	ChineseDesc     string         `json:"chineseDesc" gorm:"type:varchar(500);comment:中文描述"`
	CateID          uint           `json:"cateID" gorm:"comment:分类ID"`
	EcosystemID     uint           `json:"ecosystemID" gorm:"comment:生态ID"`
}

func CreateFeRepo(ownerName, repoName, repoDesc, language, ownerAvatarURL, homePage, githubURL, license, topics string, githubLikeCount, githubID uint, githubCreatedAt, githubUpdatedAt *time.Time) error {
	repo := FeRepo{OwnerName: ownerName, RepoName: repoName, RepoDesc: repoDesc, Language: language, OwnerAvatarURL: ownerAvatarURL, HomePage: homePage, GithubURL: githubURL, License: license, Topics: topics, GithubLikeCount: githubLikeCount, GithubID: githubID, GithubCreatedAt: githubCreatedAt, GithubUpdatedAt: githubUpdatedAt, Status: common.STATUS_NORMAL}
	err := DB.Create(&repo).Error
	return err
}

func IsRepoExist(ownerName, repoName string) bool {
	var repo FeRepo
	DB.Where("owner_name = ? AND repo_name = ?", ownerName, repoName).First(&repo)
	return repo.ID > 0
}

func GetRepos(pageNum, pageSize int) ([]FeRepo, error) {
	var repos []FeRepo
	err := DB.Scopes(Paginate(pageNum, pageSize), FieldEqual("status", common.STATUS_NORMAL)).Order("github_like_count desc").Find(&repos).Error
	return repos, err
}

func AdminGetRepos(pageNum, pageSize int) ([]FeRepo, error) {
	var repos []FeRepo
	err := DB.Scopes(Paginate(pageNum, pageSize)).Order("github_like_count desc").Find(&repos).Error
	return repos, err
}

func GetRepoByID(id uint) (*FeRepo, error) {
	var repo FeRepo
	err := DB.Scopes(FieldEqual("status", common.STATUS_NORMAL)).Where("id = ?", id).First(&repo).Error
	return &repo, err
}

func AdminGetRepoByID(id uint) (*FeRepo, error) {
	var repo FeRepo
	err := DB.Where("id = ?", id).First(&repo).Error
	return &repo, err
}

func UpdateRepoChineseInfo(id, cateID, ecosystemID, status uint, chineseHomePage, chineseDesc, tags string) error {
	err := DB.Model(&FeRepo{}).Where("id = ?", id).Updates(FeRepo{ChineseHomePage: chineseHomePage, ChineseDesc: chineseDesc, ChineseTags: tags, CateID: cateID, EcosystemID: ecosystemID, Status: status}).Error
	return err
}

func DeleteRepo(id uint) error {
	err := DB.Model(&FeRepo{}).Where("id = ?", id).Updates(FeRepo{Status: common.STATUS_DISABLED}).Error
	return err
}

func GetRepoCount() (count int64) {
	DB.Model(&FeRepo{}).Count(&count)
	return
}
