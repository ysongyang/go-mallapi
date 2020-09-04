package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	secretKey = "a2vwDwursBQGpAdBXIVLy92962cRZbjY" //JWT私钥
	maxAge    = 60 * 60 * 24 * 15                  //过期时间
	issuer    = "ysongyang"                        //wt签发者
)

//自定义Claims
type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

// @title 生成token
func CreateToken(userId int64) (*AccessToken, error) {
	expiresAt := time.Now().Add(time.Duration(maxAge) * time.Second).Unix()
	customClaims := &CustomClaims{
		UserId: userId, //用户id
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt, // 过期时间，必须设置
			Issuer:    issuer,    //签发者                                               // 非必须，也可以填充用户名，
		},
	}
	//采用HMAC SHA256加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}
	accessToken := &AccessToken{AccessToken: tokenString, ExpiresAt: expiresAt}
	return accessToken, nil
}

// @title 解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
