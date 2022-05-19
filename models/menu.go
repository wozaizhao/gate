package models

import (
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// Menu 后台管理端菜单
type Menu struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	MenuName     string         `json:"menu_name" gorm:"type:varchar(20);NOT NULL;comment:菜单名称"`
	MenuKey      string         `json:"menu_key" gorm:"type:varchar(20);NOT NULL;comment:菜单标识"`
	MenuDesc     string         `json:"menu_desc" gorm:"type:varchar(50);NOT NULL;comment:菜单描述"`
	MenuParentID uint           `json:"menu_parent_id" gorm:"type:int(11);NOT NULL;DEFAULT '0':comment:父级菜单ID"`
	Status       uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}

// 创建菜单
func CreateMenu(menuName, menuKey, menuDesc string, menuParentID uint) (menu Menu, err error) {
	menu = Menu{MenuName: menuName, MenuKey: menuKey, MenuDesc: menuDesc, MenuParentID: menuParentID, Status: common.STATUS_NORMAL}
	result := DB.Create(&menu)
	err = result.Error
	return
}

// 查询菜单树
func GetMenuTree() (menus []Menu, err error) {
	err = DB.Where("status = ?", common.STATUS_NORMAL).Find(&menus).Error
	return
}
