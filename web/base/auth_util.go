package base

import (
	"schat/types"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

const secretKey = "this is secret key"

type LoginClaims struct {
	Uid types.UserId
	jwt.RegisteredClaims
}

func GenerateToken(uid types.UserId, expireDuration time.Duration) (string, error) {
	expireTime := time.Now().Add(expireDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, LoginClaims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	})

	return token.SignedString([]byte(secretKey))
}

func ParseToken(tokenString string) (*LoginClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*LoginClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return claims, err
}
