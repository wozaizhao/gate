package models

import (
	"crypto/rand"
	"gorm.io/gorm"
	"time"
)

// Captcha 验证码
type Captcha struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ExpiredAt time.Time      `json:"expired_at"`
	Phone     string         `json:"phone" gorm:"type:varchar(20);NOT NULL;comment:手机号"`
	Code      string         `json:"code" gorm:"type:varchar(6);NOT NULL;comment:验证码"`
}

func genCaptchaCode(len int) (string, error) {
	codes := make([]byte, len)
	if _, err := rand.Read(codes); err != nil {
		return "", err
	}

	for i := 0; i < len; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}

	return string(codes), nil
}

func CreateCaptcha(phone string) (code string, err error) {
	code, err = genCaptchaCode(6)
	if err != nil {
		return
	}
	var captcha = Captcha{Phone: phone, Code: code, ExpiredAt: time.Now().Add(5 * time.Minute)}
	err = DB.Create(&captcha).Error
	return
}

func CaptchaAvailable(phone, code string) bool {
	var c Captcha
	r := DB.Where("phone = ? AND code = ? AND expired_at > ?", phone, code, time.Now()).Find(&c)
	return r.RowsAffected > 0
}
