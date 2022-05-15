package controllers

import (
	"github.com/gin-gonic/gin"
	"time"
	"wozaizhao.com/gate/common"
	"wozaizhao.com/gate/models"
)

func GetWikis(c *gin.Context) {
	pageNumParam := c.DefaultQuery("pageNum", "1")
	pageSizeParam := c.DefaultQuery("pageSize", "10")
	pageNum, _ := common.ParseInt(pageNumParam)
	pageSize, _ := common.ParseInt(pageSizeParam)
	wikis, err := models.GetFeWikis(int(pageNum), int(pageSize))
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, wikis, "")
}

func GetGists(c *gin.Context) {
	pageNumParam := c.DefaultQuery("pageNum", "1")
	pageSizeParam := c.DefaultQuery("pageSize", "10")
	pageNum, _ := common.ParseInt(pageNumParam)
	pageSize, _ := common.ParseInt(pageSizeParam)
	gists, err := models.GetFeGists(int(pageNum), int(pageSize))
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, gists, "")
}

func GetCates(c *gin.Context) {
	cates, err := models.GetFeCates()
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, cates, "")
}

func GetResources(c *gin.Context) {
	pageNumParam := c.DefaultQuery("pageNum", "1")
	pageSizeParam := c.DefaultQuery("pageSize", "10")
	cateIDParam := c.DefaultQuery("cateID", "")
	pageNum, _ := common.ParseInt(pageNumParam)
	pageSize, _ := common.ParseInt(pageSizeParam)
	cateID, _ := common.ParseInt(cateIDParam)
	resources, err := models.GetFeResourceByCateID(int(pageNum), int(pageSize), uint(cateID))
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, resources, "")
}

func GetAuthors(c *gin.Context) {
	pageNumParam := c.DefaultQuery("pageNum", "1")
	pageSizeParam := c.DefaultQuery("pageSize", "10")
	pageNum, _ := common.ParseInt(pageNumParam)
	pageSize, _ := common.ParseInt(pageSizeParam)
	authors, err := models.GetFeAuthors(int(pageNum), int(pageSize))
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, authors, "")
}

func GetAuthor(c *gin.Context) {
	authorIDParam := c.Param("authorID")
	authorID, _ := common.ParseInt(authorIDParam)
	author, err := models.GetFeAuthorByID(uint(authorID))
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, author, "")
}

func GetNavs(c *gin.Context) {
	navs, err := models.GetFeNavs()
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, navs, "")
}

type addWikiReq struct {
	WikiName    string `json:"wikiName" binding:"required"`
	WikiDesc    string `json:"wikiDesc"`
	WikiContent string `json:"wikiContent" binding:"required"`
}

func AddWiki(c *gin.Context) {
	var wiki addWikiReq
	if err := c.ShouldBindJSON(&wiki); err != nil {
		RenderError(c, err)
		return
	}
	if err := models.CreateFeWiki(wiki.WikiName, wiki.WikiDesc, wiki.WikiContent); err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, nil, "")
}

type addGistReq struct {
	GistName    string `json:"gistName" binding:"required"`
	GistDesc    string `json:"gistDesc"`
	GistContent string `json:"gistContent" binding:"required"`
}

func AddGist(c *gin.Context) {
	var gist addGistReq
	if err := c.ShouldBindJSON(&gist); err != nil {
		RenderError(c, err)
		return
	}
	if err := models.CreateFeGist(gist.GistName, gist.GistDesc, gist.GistContent); err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, nil, "")
}

type addCateReq struct {
	CateName string `json:"cateName" binding:"required"`
	CateDesc string `json:"cateDesc"`
}

func AddCate(c *gin.Context) {
	var cate addCateReq
	if err := c.ShouldBindJSON(&cate); err != nil {
		RenderError(c, err)
		return
	}
	if err := models.CreateFeCate(cate.CateName, cate.CateDesc); err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, nil, "")
}

type addResourceReq struct {
	ResourceName      string `json:"resourceName" binding:"required"`
	Thumb             string `json:"thumb"`
	Description       string `json:"description"`
	WebURL            string `json:"webURL" binding:"required"`
	DocURL            string `json:"docURL"`
	GitHubURL         string `json:"githubURL"`
	MiniAppID         string `json:"miniAppID"`
	OfficialAccountID string `json:"officialAccountID"`
	CateID            string `json:"cateID" binding:"required"`
	AuthorID          string `json:"authorID"`
}

func AddResource(c *gin.Context) {
	var resource addResourceReq
	if err := c.ShouldBindJSON(&resource); err != nil {
		RenderError(c, err)
		return
	}
	cateID, _ := common.ParseInt(resource.CateID)
	authorID, _ := common.ParseInt(resource.AuthorID)
	if err := models.CreateFeResource(resource.ResourceName, resource.Thumb, resource.Description, resource.WebURL, resource.DocURL, resource.GitHubURL, resource.MiniAppID, resource.OfficialAccountID, uint(cateID), uint(authorID)); err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, nil, "")
}

type addAuthorReq struct {
	AuthorName string `json:"authorName" binding:"required"`
	AuthorDesc string `json:"authorDesc"`
	WebURL     string `json:"webURL"`
	AvatarURL  string `json:"avatarURL"`
	GitHubURL  string `json:"githubURL"`
}

func AddAuthor(c *gin.Context) {
	var author addAuthorReq
	if err := c.ShouldBindJSON(&author); err != nil {
		RenderError(c, err)
		return
	}
	if err := models.CreateFeAuthor(author.AuthorName, author.AuthorDesc, author.WebURL, author.AvatarURL, author.GitHubURL); err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, nil, "")
}

type addNavReq struct {
	NavTitle            string `json:"navTitle" binding:"required"`
	icon                string
	image               string
	whiteBgIconColor    string
	colorBgIconColor    string
	colorBgColor        string
	gradientBgIconColor string
	gradientBgColor     string
	navType             uint
	navPath             string
	navParams           string
}

func AddNav(c *gin.Context) {
	var nav addNavReq
	if err := c.ShouldBindJSON(&nav); err != nil {
		RenderError(c, err)
		return
	}
	if err := models.CreateFeNav(nav.NavTitle, nav.icon, nav.image, nav.whiteBgIconColor, nav.colorBgIconColor, nav.colorBgColor, nav.gradientBgIconColor, nav.gradientBgColor, nav.navType, nav.navPath, nav.navParams); err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, nil, "")
}

type addRepoReq struct {
	OwnerName       string `json:"ownerName" binding:"required"`
	RepoName        string `json:"repoName" binding:"required"`
	RepoDesc        string `json:"repoDesc"`
	Language        string `json:"language"`
	OwnerAvatarURL  string `json:"ownerAvatarURL"`
	HomePage        string `json:"homePage"`
	GithubURL       string `json:"githubURL"`
	GithubCreatedAt string `json:"githubCreatedAt"`
	GithubUpdatedAt string `json:"githubUpdatedAt"`
	License         string `json:"license"`
	Topics          string `json:"topics"`
	GithubLikeCount int    `json:"githubLikeCount"`
	GithubID        int    `json:"githubID"`
}

func AddRepo(c *gin.Context) {
	var repo addRepoReq
	if err := c.ShouldBindJSON(&repo); err != nil {
		RenderError(c, err)
		return
	}
	if models.IsRepoExist(repo.OwnerName, repo.RepoName) {
		RenderFail(c, "repo already exist!")
		return
	}
	// githubLikeCount, _ := common.ParseInt(repo.GithubLikeCount)
	githubCreatedAt, _ := time.Parse(""+time.RFC3339+"", repo.GithubCreatedAt)
	githubUpdatedAt, _ := time.Parse(""+time.RFC3339+"", repo.GithubUpdatedAt)
	if err := models.CreateFeRepo(repo.OwnerName, repo.RepoName, repo.RepoDesc, repo.Language, repo.OwnerAvatarURL, repo.HomePage, repo.GithubURL, repo.License, repo.Topics, uint(repo.GithubLikeCount), uint(repo.GithubID), &githubCreatedAt, &githubUpdatedAt); err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, nil, "")
}
