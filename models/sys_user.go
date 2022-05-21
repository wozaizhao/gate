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
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Nickname  string         `json:"nickname" gorm:"type:varchar(50);DEFAULT '';comment:昵称"`
	Bio       string         `json:"bio" gorm:"type:varchar(50);DEFAULT '';comment:简介"`
	AvatarURL string         `json:"avatarUrl" gorm:"type:varchar(255);DEFAULT '';comment:头像"`
	Gender    int            `json:"gender" gorm:"type:tinyint(1);DEFAULT '0';comment:性别"`
	Username  string         `json:"username" gorm:"type:varchar(30);DEFAULT '';comment:用户名"`
	Password  string         `json:"password" gorm:"type:varchar(64);DEFAULT '';comment:密码"`
	OpenID    string         `json:"open_id" gorm:"unique;type:varchar(40);DEFAULT ''"`
	Roles     []Role         `json:"roles" gorm:"many2many:user_role;"`
	Status    uint           `json:"status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0';comment:状态"`
}

func GetUserByOpenID(openID string) (User, error) {
	user, exist := openIDExist(openID)
	if exist {
		return user, nil
	} else {
		user, err := createUser(openID)
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

// 创建帐户
func createUser(openID string) (user User, err error) {
	user = User{OpenID: openID, Status: common.USER_STATUS_NORMAL}
	result := DB.Create(&user)
	err = result.Error
	return
}

func openIDExist(openID string) (user User, exist bool) {
	r := DB.Where("open_id = ? ", openID).Find(&user)
	exist = r.RowsAffected > 0
	return
}

// 获取用户信息
func GetUserByID(userID uint) (user User, err error) {
	r := DB.Where("id = ? ", userID).Find(&user)
	err = r.Error
	return
}

// type UserInfoWithRole struct {
// 	ID           uint      `json:"id"`
// 	Username     string    `json:"username"`
// 	Nickname     string    `json:"nickname"`
// 	AvatarURL    string    `json:"avatarUrl"`
// 	Gender       int       `json:"gender"`
// 	Status       uint      `json:"status"`
// 	CreatedAt    time.Time `json:"createdAt"`
// 	Bio          string    `json:"bio"`
// 	RoleNames    string    `json:"role_names"`
// 	RoleKeys     string    `json:"role_keys"`
// 	RoleStatuses string    `json:"role_statuses"`
// }

// 分页获取用户 联合查询用户role
func GetUsers(pageNum, pageSize int) (users []User, err error) {
	err = DB.Scopes(Paginate(pageNum, pageSize)).Preload("Roles").Find(&users).Error
	return users, err
	// r := DB.Scopes(Paginate(pageNum, pageSize)).Find(&users)
	// err = r.Error
	// sqlstr := `select u.id, u.username, u.nickname, u.avatar_url, u.gender, u.status, u.created_at,
	//         u.bio, GROUP_CONCAT(r.role_name) as role_names, GROUP_CONCAT(r.role_key) as role_keys, GROUP_CONCAT(r.status) as role_statuses
	// 		from users as u
	// 			left join user_roles as ur on u.id = ur.user_id
	// 			left join roles as r on r.id = ur.role_id GROUP BY u.id`
	// r := DB.Raw(sqlstr).Scan(&users)
	// // result := DB.Where("user_id = ?", userID).Find(&userRoles)
	// err = r.Error
	// return
}

func GetUserCount() (count int64) {
	DB.Model(&User{}).Count(&count)
	return
}

// 更新用户信息
func UpdateUser(userID uint, gender int, nickname, avatarURL, bio string) (res bool, err error) {
	var user User
	user.ID = userID
	r := DB.Model(&user).Updates(User{Gender: gender, Nickname: nickname, AvatarURL: avatarURL, Bio: bio})
	return r.RowsAffected > 0, r.Error
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
