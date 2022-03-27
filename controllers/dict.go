package controllers

import (
	"github.com/gin-gonic/gin"
	"wozaizhao.com/gate/models"
)

type geoNodes struct {
	ParentCode string           `json:"parent_code" binding:"required"`
	Nodes      []models.GeoNode `json:"nodes" binding:"required"`
}

type geoNodeHasLetter struct {
	Letter     string `json:"letter" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Code       string `json:"code" binding:"required"`
	ParentCode string `json:"parent_code" binding:"required"`
}

func GetCities(c *gin.Context) {
	data, err := models.GetCities()
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, data, "")
}

func GetProvinces(c *gin.Context) {
	data, err := models.GetProvinces()
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, data, "")
}

func GetRegions(c *gin.Context) {
	code := c.DefaultQuery("code", "86")
	data, err := models.GetRegionsByCode(code)
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, data, "")
}

type geoNode struct {
	Name     string     `json:"name"`
	Code     string     `json:"code"`
	Children []*geoNode `json:"children"`
}

func GetAllRegions(c *gin.Context) {
	provinces, err := models.GetProvinceRegions()
	if err != nil {
		RenderError(c, err)
		return
	}
	all, err := models.GetAllRegions()
	if err != nil {
		RenderError(c, err)
		return
	}
	tree := make([]*geoNode, 0)
	for _, province := range provinces {
		provinceNode := &geoNode{
			Name:     province.Name,
			Code:     province.Code,
			Children: make([]*geoNode, 0),
		}
		for _, region := range all {
			if region.ParentCode == province.Code {
				var child = &geoNode{
					Name:     region.Name,
					Code:     region.Code,
					Children: make([]*geoNode, 0),
				}
				for _, area := range all {
					if area.ParentCode == region.Code {
						child.Children = append(child.Children, &geoNode{
							Name:     area.Name,
							Code:     area.Code,
							Children: make([]*geoNode, 0),
						})
					}
				}
				provinceNode.Children = append(provinceNode.Children, child)
			}
		}
		tree = append(tree, provinceNode)
	}
	// for _, v := range data {
	// 	if v.ParentCode == "86" {
	// 		tree = append(tree, &geoNode{Name: v.Name, Code: v.Code, Children: make([]*geoNode, 0)})
	// 	}
	// }
	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, tree, "")
}

// func makeTree(data interface{}) (tree []geoNode) {
// 	for _, v := range data.([]models.Region) {

// 	}
// 	return
// }

func AddCity(c *gin.Context) {

	var s geoNodeHasLetter
	if err := c.ShouldBindJSON(&s); err != nil {
		RenderBadRequest(c, err)
		return
	}

	err := models.CreateCity(s.Letter, s.Name, s.Code, s.ParentCode)

	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, true, "添加成功")
}

func AddProvince(c *gin.Context) {

	var s geoNodeHasLetter
	if err := c.ShouldBindJSON(&s); err != nil {
		RenderBadRequest(c, err)
		return
	}

	err := models.CreateProvince(s.Letter, s.Name, s.Code, s.ParentCode)

	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, true, "添加成功")
}

func AddRegion(c *gin.Context) {

	var s geoNodes
	if err := c.ShouldBindJSON(&s); err != nil {
		RenderBadRequest(c, err)
		return
	}

	err := models.CreateRegions(s.ParentCode, s.Nodes)

	if err != nil {
		RenderError(c, err)
		return
	}
	RenderSuccess(c, true, "添加成功")
}
