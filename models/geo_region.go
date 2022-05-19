package models

import (
	"time"

	"gorm.io/gorm"
)

// Region 行政区
type Region struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Name       string         `json:"name" gorm:"type:varchar(15);NOT NULL;comment:行政区名称"`
	Code       string         `json:"code" gorm:"type:varchar(10);NOT NULL;comment:行政区编码"`
	ParentCode string         `json:"parent_code" gorm:"type:varchar(10);NOT NULL;comment:父级编码"`
}

// 创建城市
func CreateRegion(name, code, parentCode string) (err error) {
	var region = Region{Name: name, Code: code, ParentCode: parentCode}
	result := DB.Create(&region)
	err = result.Error
	return
}

type GeoNode struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

func CreateRegions(parentCode string, nodes []GeoNode) (err error) {
	var regions []Region
	for _, node := range nodes {
		regions = append(regions, Region{Name: node.Name, Code: node.Code, ParentCode: parentCode})
	}

	result := DB.Create(&regions)

	return result.Error
}

type regionRes struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func GetRegionsByCode(code string) (regions []regionRes, err error) {
	r := DB.Table("regions").Where("parent_code = ?", code).Scan(&regions)
	err = r.Error
	return
}

type regionResWithParent struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	ParentCode string `json:"parent_code"`
}

func GetProvinceRegions() (all []regionResWithParent, err error) {
	result := DB.Table("regions").Where("parent_code = ?", "86").Scan(&all)
	err = result.Error
	return
}

func GetAllRegions() (all []regionResWithParent, err error) {
	result := DB.Table("regions").Scan(&all)
	err = result.Error
	return
}
