package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"wozaizhao.com/gate/common"
)

// User 用户
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Nickname  string         `json:"nickname" gorm:"type:varchar(50);DEFAULT '';comment:昵称"`
	Bio       string         `json:"bio" gorm:"type:varchar(50);DEFAULT '';comment:简介"`
	AvatarURL string         `json:"avatarUrl" gorm:"type:varchar(255);DEFAULT '';comment:头像"`
	Gender    int            `json:"gender" gorm:"type:tinyint(1);DEFAULT '0';comment:性别"`
	Phone     string         `json:"phone" gorm:"type:varchar(20);unique;DEFAULT '';comment:手机号"`
	Username  string         `json:"username" gorm:"type:varchar(30);DEFAULT '';comment:用户名"`
	Password  string         `json:"password" gorm:"type:varchar(64);DEFAULT '';comment:密码"`
	Status    uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
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

func VerifyUser(username, password string) (user User, err error) {
	r := DB.Where("username = ? and password = ?", username, password).Find(&user)
	exist := r.RowsAffected > 0
	if exist {
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

// result := db.Create(&user) // pass pointer of data to Create

// user.ID             // returns inserted data's primary key
// result.Error        // returns error
// result.RowsAffected // returns inserted records count

// 创建帐户
func createUser(phone string) (user User, err error) {
	user = User{Phone: phone, Nickname: "用户" + phone[5:], Status: common.USER_STATUS_NORMAL}
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

type UserInfoWithRole struct {
	ID           uint      `json:"id"`
	Username     string    `json:"username"`
	Nickname     string    `json:"nickname"`
	AvatarURL    string    `json:"avatarUrl"`
	Phone        string    `json:"phone"`
	Gender       int       `json:"gender"`
	Status       uint      `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	Bio          string    `json:"bio"`
	RoleNames    string    `json:"role_names"`
	RoleKeys     string    `json:"role_keys"`
	RoleStatuses string    `json:"role_statuses"`
}

// 分页获取用户 联合查询用户role
func GetUsers(pageNum, pageSize int) (users []UserInfoWithRole, err error) {
	// r := DB.Scopes(Paginate(pageNum, pageSize)).Find(&users)
	// err = r.Error
	sqlstr := `select u.id, u.username, u.nickname, u.avatar_url, u.phone, u.gender, u.status, u.created_at,
	        u.bio, GROUP_CONCAT(r.role_name) as role_names, GROUP_CONCAT(r.role_key) as role_keys, GROUP_CONCAT(r.status) as role_statuses
			from users as u 
				left join user_roles as ur on u.id = ur.user_id
				left join roles as r on r.id = ur.role_id GROUP BY u.id`
	r := DB.Raw(sqlstr).Scan(&users)
	// result := DB.Where("user_id = ?", userID).Find(&userRoles)
	err = r.Error
	return
}

func GetUserCount() (count int64) {
	DB.Model(&User{}).Count(&count)
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

// // 查询用户角色
// func GetUserRole(userID uint) int {
// 	var user User
// 	r := DB.Where("id = ? ", userID).Find(&user)
// 	if r.RowsAffected > 0 {
// 		return int(user.Role)
// 	}
// 	return -1
// }

// 设置用户积分
func UpdateUserCredit(userID uint, credit int) {

}

// 查询用户积分
func GetUserCredit(userID uint) {

}
