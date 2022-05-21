package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
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

type editRepoReq struct {
	ChineseHomePage string `json:"chineseHomePage"`
	ChineseDesc     string `json:"chineseDesc"`
	Tags            string `json:"tags"`
	CateID          int    `json:"cateID"`
	Status          int    `json:"status"`
	EcosystemID     int    `json:"ecosystemID"`
}

func EditRepo(c *gin.Context) {
	var idInUri IDInUri
	var repo editRepoReq
	if err := c.ShouldBindUri(&idInUri); err != nil {
		RenderError(c, err)
		return
	}
	if err := c.ShouldBindJSON(&repo); err != nil {
		RenderError(c, err)
		return
	}
	repoID, _ := common.ParseInt(idInUri.ID)
	if err := models.UpdateRepoChineseInfo(uint(repoID), uint(repo.CateID), uint(repo.EcosystemID), uint(repo.Status), repo.ChineseHomePage, repo.ChineseDesc, repo.Tags); err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, nil, "")
}

func DeleteRepo(c *gin.Context) {
	var idInUri IDInUri
	if err := c.ShouldBindUri(&idInUri); err != nil {
		RenderError(c, err)
		return
	}
	repoID, _ := common.ParseInt(idInUri.ID)
	if err := models.DeleteRepo(uint(repoID)); err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, nil, "")
}

type RepoRes struct {
	List  []models.FeRepo `json:"list"`
	Total int64           `json:"total"`
}

func GetRepos(c *gin.Context) {
	pageNumParam := c.DefaultQuery("pageNum", "1")
	pageSizeParam := c.DefaultQuery("pageSize", "10")
	name := c.DefaultQuery("name", "")
	pageNum, _ := common.ParseInt(pageNumParam)
	pageSize, _ := common.ParseInt(pageSizeParam)
	repos, err := models.AdminGetRepos(int(pageNum), int(pageSize), name)
	if err != nil {
		RenderError(c, err)
		return
	}
	var total = models.GetRepoCount(name)
	var pageCount = float64(total) / float64(pageSize)
	var res = RepoRes{
		List:  repos,
		Total: common.Round(pageCount),
	}
	RenderSuccess(c, res, "")
}

func GetRepo(c *gin.Context) {
	var idInUri IDInUri
	if err := c.ShouldBindUri(&idInUri); err != nil {
		RenderError(c, err)
		return
	}
	repoID, _ := common.ParseInt(idInUri.ID)
	repo, err := models.AdminGetRepoByID(uint(repoID))

	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, repo, "")
}

type addFeRepoCateReq struct {
	CateName     string `json:"cateName" binding:"required"`
	CateCNName   string `json:"cateCNName" binding:"required"`
	CateDesc     string `json:"cateDesc"`
	CateParentID int    `json:"cateParentID"`
}

func AddFeRepoCate(c *gin.Context) {
	var cate addFeRepoCateReq
	if err := c.ShouldBindJSON(&cate); err != nil {
		RenderError(c, err)
		return
	}
	repoCate, err := models.CreateFeRepoCate(cate.CateName, cate.CateCNName, cate.CateDesc, uint(cate.CateParentID))
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, repoCate, "")
}

type cateNode struct {
	ID         uint        `json:"id"`
	CateName   string      `json:"cateName"`
	CateCNName string      `json:"cateCNName"`
	CateDesc   string      `json:"cateDesc"`
	Children   []*cateNode `json:"children"`
}

func GetFeCates(c *gin.Context) {
	firstClasses, err := models.GetFirstClassRepoCates()
	if err != nil {
		RenderError(c, err)
		return
	}
	all, err := models.GetFeRepoCates()
	if err != nil {
		RenderError(c, err)
		return
	}
	tree := make([]*cateNode, 0)

	for _, first := range firstClasses {
		firstNode := &cateNode{
			CateName:   first.CateName,
			CateCNName: first.CateCNName,
			CateDesc:   first.CateDesc,
			ID:         first.ID,
			// Children:   make([]*cateNode, 0),
		}
		firstChildren := make([]*cateNode, 0)
		for _, cate := range all {
			if cate.CateParentID == firstNode.ID {
				var child = &cateNode{
					CateName:   cate.CateName,
					CateCNName: cate.CateCNName,
					CateDesc:   cate.CateDesc,
					ID:         cate.ID,
					// Children:   make([]*cateNode, 0),
				}
				firstChildren = append(firstChildren, child)
			}
		}
		if len(firstChildren) > 0 {
			firstNode.Children = firstChildren
		}
		tree = append(tree, firstNode)
	}
	RenderSuccess(c, tree, "")
}

type addEcosystemReq struct {
	EcosystemName string `json:"ecosystemName" binding:"required"`
	EcosystemDesc string `json:"ecosystemDesc"`
}

func AddEcosystem(c *gin.Context) {
	var req addEcosystemReq
	if err := c.ShouldBindJSON(&req); err != nil {
		RenderError(c, err)
		return
	}
	ecosystem, err := models.CreateFeEcosystem(req.EcosystemName, req.EcosystemDesc)
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, ecosystem, "")
}

func GetEcosystems(c *gin.Context) {
	ecosystems, err := models.GetFeEcosystems()
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, ecosystems, "")
}

type addFundamentalReq struct {
	Cate         string          `json:"cate" binding:"required"`
	CateCn       string          `json:"cateCn"`
	Topic        string          `json:"topic" binding:"required"`
	TopicCn      string          `json:"topicCn"`
	TopicDesc    string          `json:"topicDesc" binding:"required"`
	TopticDescCn string          `json:"topicDescCn"`
	Links        []models.FeLink `json:"links"`
}

func AddFundamental(c *gin.Context) {
	var fund addFundamentalReq
	if err := c.ShouldBindJSON(&fund); err != nil {
		RenderError(c, err)
		return
	}
	err := models.CreateFeFundamental(fund.Cate, fund.CateCn, fund.Topic, fund.TopicCn, fund.TopicDesc, fund.TopticDescCn, fund.Links)
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, nil, "")
}

type FundRes struct {
	List  []models.FeFundamental `json:"list"`
	Total int64                  `json:"total"`
}

func GetFundamental(c *gin.Context) {
	pageNumParam := c.DefaultQuery("pageNum", "1")
	pageSizeParam := c.DefaultQuery("pageSize", "10")
	pageNum, _ := common.ParseInt(pageNumParam)
	pageSize, _ := common.ParseInt(pageSizeParam)
	funds, err := models.AdminGetFundamentals(int(pageNum), int(pageSize))
	if err != nil {
		RenderError(c, err)
		return
	}
	var total, _ = models.GetFundamentalCount()
	var pageCount = float64(total) / float64(pageSize)
	var res = FundRes{
		List:  funds,
		Total: common.Round(pageCount),
	}
	RenderSuccess(c, res, "")
}

type editFundamentalReq struct {
	ID           int             `json:"id" binding:"required"`
	Cate         string          `json:"cate" binding:"required"`
	CateCn       string          `json:"cateCn"`
	Topic        string          `json:"topic" binding:"required"`
	TopicCn      string          `json:"topicCn"`
	TopicDesc    string          `json:"topicDesc" binding:"required"`
	TopticDescCn string          `json:"topicDescCn"`
	Status       int             `json:"status"`
	Links        []models.FeLink `json:"links"`
}

func EditFundamental(c *gin.Context) {
	var fund editFundamentalReq
	if err := c.ShouldBindJSON(&fund); err != nil {
		RenderError(c, err)
		return
	}
	err := models.UpdateFeFundamental(uint(fund.ID), uint(fund.Status), fund.Cate, fund.CateCn, fund.Topic, fund.TopicCn, fund.TopicDesc, fund.TopticDescCn, fund.Links)
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, nil, "")
}
