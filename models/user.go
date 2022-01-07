package models

import (
	"gorm.io/gorm"
	"time"
)

// User 用户
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	OpenID    string         `json:"open_id" gorm:"type:varchar(40);DEFAULT ''"`         // openID 小程序登录获取
	UnionID   string         `json:"union_id" gorm:"type:varchar(40);DEFAULT ''"`        // unionID 满足条件下小程序处获取
	NickName  string         `json:"nickname" gorm:"type:varchar(50);DEFAULT ''"`        // 昵称
	AvatarURL string         `json:"avatar_url" gorm:"type:varchar(255);DEFAULT ''"`     // 头像
	Gender    int            `json:"gender" gorm:"type:tinyint(1);DEFAULT '0'"`          // 性别
	Phone     string         `json:"phone" gorm:"type:varchar(20);unique;DEFAULT ''"`    // 手机号
	Username  string         `json:"username" gorm:"type:varchar(30);unique;DEFAULT ''"` // 用户名
	Password  string         `json:"password" gorm:"type:varchar(30);DEFAULT ''"`        // 密码
	Status    uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0'"` // 状态 1 初始值 正常 2 已失效
	Role      uint           `json:"role" gorm:"type:tinyint(1);DEFAULT '0'"`            // 角色 1 初始值 普通用户 2 管理员 3 vip
	Credit    int            `json:"credit" gorm:"type:int(11);DEFAULT '0'"`             // 用户积分
}

func GetUserByPhone(phone, openID string) (User, error) {
	user, exist := phoneExist(phone)
	if exist {
		return user, nil
	} else {
		user, err := createUser(phone, openID)
		return user, err
	}
}

// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

// result := db.Create(&user) // pass pointer of data to Create

// user.ID             // returns inserted data's primary key
// result.Error        // returns error
// result.RowsAffected // returns inserted records count

// 创建帐户
func createUser(phone, openID string) (user User, err error) {
	user = User{Phone: phone, OpenID: openID}
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
func GetUserByID(userID uint) {

}

// 更换手机号
func UpdatePhone(oldphone, phone string) {

}

// 关联用户openID
func LinkUserByOpenID(phone, openID string) {

}

// 使用微信信息设置用户昵称和头像
func SetUserInfoByWechat(openID, nickname, avatar string) {

}

// 设置用户昵称
func SetUserNickname(userID uint, nickname string) {

}

// 设置用户头像
func SetUserAvatar(userID uint, avatar string) {

}

// 设置用户性别
func SetUserGender(userID uint, gender int) {

}

// 设置用户名密码
func SetUsernameAndPassword(userID uint, username, password string) {

}

// 设置用户状态
func UpdateUserStatus(userID uint, status int) {

}

// 查询用户状态
func GetUserStatus(userID uint) int {
	return 1
}

// 设置用户角色
func UpdateUserRole(userID uint, role int) {

}

// 查询用户角色
func GetUserRole(userID uint) int {
	return 1
}

// 设置用户积分
func UpdateUserCredit(userID uint, credit int) {

}

// 查询用户积分
func GetUserCredit(userID uint) {

}
