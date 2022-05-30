package models

import (
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// Menu 后台管理端菜单
type Menu struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Name        string         `json:"name" gorm:"type:varchar(30);NOT NULL;comment:菜单名称"`
	Path        string         `json:"path" gorm:"type:varchar(60);NOT NULL;comment:菜单路径"`
	Redirect    string         `json:"redirect" gorm:"type:varchar(60);NOT NULL;comment:菜单重定向"`
	Title       string         `json:"title" gorm:"type:varchar(30);NOT NULL;comment:菜单标题"`
	Icon        string         `json:"icon" gorm:"type:varchar(30);NOT NULL;comment:菜单图标"`
	Component   string         `json:"component" gorm:"type:varchar(60);NOT NULL;comment:菜单组件"`
	Sort        int            `json:"sort" gorm:"type:int(11);NOT NULL;comment:菜单排序"`
	ParentID    uint           `json:"parentID" gorm:"type:int(11);NOT NULL;DEFAULT '0':comment:父级菜单ID"`
	Permissions string         `json:"permissions" gorm:"type:varchar(255);comment:菜单权限"`
	ExternalURL string         `json:"externalURL" gorm:"type:varchar(255);comment:外链地址"`
	Status      uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}

// 创建菜单
func CreateMenu(name, path, redirect, title, icon, component, permissions, externalURL string, sort int, parentID uint) (menu Menu, err error) {
	menu = Menu{Name: name, Path: path, Redirect: redirect, Title: title, Icon: icon, Component: component, Permissions: permissions, ExternalURL: externalURL, Sort: sort, ParentID: parentID, Status: common.STATUS_NORMAL}
	result := DB.Create(&menu)
	err = result.Error
	return
}

// 查询菜单树
func GetMenus() (menus []Menu, err error) {
	err = DB.Where("status = ?", common.STATUS_NORMAL).Find(&menus).Error
	return
}
