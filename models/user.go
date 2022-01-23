package models

import (
	"gorm.io/gorm"
	"time"
	"wozaizhao.com/gate/common"
)

// User 用户
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// OpenID    string         `json:"open_id" gorm:"type:varchar(40);DEFAULT ''"`         // openID 小程序登录获取
	// UnionID   string         `json:"union_id" gorm:"type:varchar(40);DEFAULT ''"`        // unionID 满足条件下小程序处获取
	Nickname  string `json:"nickname" gorm:"type:varchar(50);DEFAULT ''"`        // 昵称
	Bio       string `json:"bio" gorm:"type:varchar(50);DEFAULT ''"`             // 简介
	AvatarURL string `json:"avatarUrl" gorm:"type:varchar(255);DEFAULT ''"`      // 头像
	Gender    int    `json:"gender" gorm:"type:tinyint(1);DEFAULT '0'"`          // 性别
	Phone     string `json:"phone" gorm:"type:varchar(20);unique;DEFAULT ''"`    // 手机号
	Username  string `json:"username" gorm:"type:varchar(30);DEFAULT ''"`        // 用户名
	Password  string `json:"password" gorm:"type:varchar(30);DEFAULT ''"`        // 密码
	Status    uint   `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0'"` // 状态 1 初始值 正常 2 已失效
	Role      uint   `json:"role" gorm:"type:tinyint(1);DEFAULT '0'"`            // 角色 1 初始值 普通用户 2 管理员 3 vip
	Credit    int    `json:"credit" gorm:"type:int(11);DEFAULT '0'"`             // 用户积分
}

func GetUserByPhone(phone string) (User, error) {
	user, exist := phoneExist(phone)
	if exist {
		return user, nil
	} else {
		user, err := createUser(phone)
		return user, err
	}
}

// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

// result := db.Create(&user) // pass pointer of data to Create

// user.ID             // returns inserted data's primary key
// result.Error        // returns error
// result.RowsAffected // returns inserted records count

// 创建帐户
func createUser(phone string) (user User, err error) {
	user = User{Phone: phone, Nickname: "用户" + phone[5:], Status: common.STATUS_NORMAL}
	result := DB.Create(&user)
	err = result.Error
	return
}

func phoneExist(phone string) (user User, exist bool) {
	r := DB.Where("phone = ? ", phone).Find(&user)
	exist = r.RowsAffected > 0
	return
}

// 获取用户信息
func GetUserByID(userID uint) (user User, err error) {
	r := DB.Where("id = ? ", userID).Find(&user)
	err = r.Error
	return
}

// 获取用户信息
// func GetUserByOpenID(openID string) (user User, err error) {
// 	r := DB.Where("open_id = ?", openID).Find(&user)
// 	err = r.Error
// 	return
// }

// 更新用户信息
func UpdateUser(userID uint, Gender int, Phone, Nickname, AvatarURL, Username, Password, Bio string) (res bool, err error) {
	var user User
	user.ID = userID
	r := DB.Model(&user).Updates(User{Gender: Gender, Phone: Phone, Nickname: Nickname, AvatarURL: AvatarURL, Username: Username, Password: Password, Bio: Bio})
	return r.RowsAffected > 0, r.Error
}

// 更换手机号
func UpdatePhone(oldphone, phone string) {

}

// 关联用户openID
// func LinkUserByOpenID(userID uint, openID string) {
// 	r := DB.Model(&User{}).Where("id = ?", userID).Update("open_id", openID)
// 	if r.Error != nil {
// 		common.LogError("LinkUserByOpenID", r.Error)
// 	}
// }

// 使用微信信息设置用户昵称和头像
func SetUserInfoByWechat(openID, nickname, avatar string) {

}

// 设置用户名密码
func SetUsernameAndPassword(userID uint, username, password string) {

}

// 设置用户状态
func UpdateUserStatus(userID uint, status int) {

}

// 查询用户状态
func GetUserStatus(userID uint) int {
	var user User
	r := DB.Where("id = ? ", userID).Find(&user)
	if r.RowsAffected > 0 {
		return int(user.Status)
	}
	return -1
}

// 设置用户角色
func UpdateUserRole(userID uint, role int) {

}

// 查询用户角色
func GetUserRole(userID uint) int {
	var user User
	r := DB.Where("id = ? ", userID).Find(&user)
	if r.RowsAffected > 0 {
		return int(user.Role)
	}
	return -1
}

// 设置用户积分
func UpdateUserCredit(userID uint, credit int) {

}

// 查询用户积分
func GetUserCredit(userID uint) {

}
