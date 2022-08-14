package helper

import (
	_ "embed"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

//go:embed certs/ca.pem
var verifyKeyBuf []byte

//go:embed certs/ca.key
var signKeyBuf []byte

var (
	// AccessExpiresIn token access expires in duration
	AccessExpiresIn = 12 * time.Hour
)

// UserClaims 用户信息
type UserClaims struct {
	ClientID string // 客户端ID
	ID       string // 用户ID
	Name     string // 用户名
	Role     string // 当前角色
}

// JwtClaims ...
type JwtClaims struct {
	jwt.StandardClaims
	UserClaims
}

// CreateToken ...
func CreateToken(claims JwtClaims) (token map[string]string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signKeyBuf)
	if err != nil {
		logrus.Errorln("Token ParseRSAPrivateKeyFromPEM Error ", err)
		return
	}
	accessToken, err := jwtToken.SignedString(signKey)
	if err != nil {
		logrus.Errorln("Token SignedString Error ", err)
		return
	}

	expiresAt := strconv.FormatInt(claims.ExpiresAt, 10)
	expiresIn := strconv.FormatInt(int64(AccessExpiresIn/time.Second), 10)
	// if claims.Role == "attendant" {
	// 	expiresIn = strconv.FormatInt(int64(AccessExpiresIn/time.Second)*48*365, 10)
	// }
	token = map[string]string{
		"access_token": accessToken,
		"token_type":   "Bearer",
		"expires_in":   expiresIn,
		"expires_at":   expiresAt,
	}
	return
}

// ParseToken ...
func ParseToken(token string) (*JwtClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(parsedToken *jwt.Token) (interface{}, error) {
		// the key used to validate tokens
		return jwt.ParseRSAPublicKeyFromPEM(verifyKeyBuf)
	})
	if err != nil {
		logrus.Errorln("parseToken ParseWithClaims Error ", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, err
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				claims := parsedToken.Claims.(*JwtClaims)
				return claims, err
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, err
			} else {
				return nil, err
			}
		}
	}
	claims, ok := parsedToken.Claims.(*JwtClaims)
	if ok && parsedToken.Valid {
		return claims, nil
	}
	return &JwtClaims{}, err
}
