package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"wozaizhao.com/gate/config"
)

var jwtSecret = config.GetConfig().JwtSecret

// Claims
type Claims struct {
	UserID uint   `json:"userID"`
	Phone  string `json:"phone"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(userID uint, phone string) (string, error) {
	nowtime := time.Now()
	expireTime := nowtime.Add(24 * time.Hour)
	claims := Claims{
		userID,
		phone,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "wozaizhao",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret))
	return token, err
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, err
		}
	}
	return nil, err
}
