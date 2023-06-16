package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"ls-kh-rl/internal/errors"
	"time"
)

const accessTokenKey = "ACCESS-TOKEN-KEY-FOR-RL"
const refreshTokenKey = "REFRESH-TOKEN-KEY-FOR-RL"

type MyClaim struct {
	Id        int64
	TokenType string
	jwt.StandardClaims
}

func GenerateAccessToken(claim *MyClaim) string {
	iss := time.Now()
	exp := iss.Add(24 * time.Hour)
	claim.IssuedAt = iss.Unix()
	claim.ExpiresAt = exp.Unix()
	claim.TokenType = "access_token"
	signedString, e := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(accessTokenKey))
	if e != nil {
		fmt.Println(e)
	}
	return signedString
}

func GenerateRefreshToken(claim *MyClaim) string {
	iss := time.Now()
	exp := iss.Add(7 * 24 * time.Hour)
	claim.IssuedAt = iss.Unix()
	claim.ExpiresAt = exp.Unix()
	claim.TokenType = "refresh_token"
	signedString, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(refreshTokenKey))
	return signedString
}

func ParseJWT(token string) (*MyClaim, error) {
	jwtToken1, err1 := jwt.ParseWithClaims(token, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessTokenKey), nil
	})
	jwtToken2, err2 := jwt.ParseWithClaims(token, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(refreshTokenKey), nil
	})
	jwtToken := &jwt.Token{}
	if err1 != nil && err2 != nil {
		return nil, errors.ErrTokenParse
	}
	if err1 == nil {
		jwtToken = jwtToken1
	}
	if err2 != nil {
		jwtToken = jwtToken2
	}
	if myClaim, ok := jwtToken.Claims.(*MyClaim); ok {
		return myClaim, nil
	}
	return nil, errors.ErrTokenParse
}

func IsAccessToken(token string) (id int64, err error) {
	claim, err := ParseJWT(token)
	if err != nil {
		return
	}
	if claim.ExpiresAt < time.Now().Unix() {
		err = errors.ErrTokenExp
		return
	}
	if claim.TokenType != "access_token" {
		err = errors.ErrTokenType
		return
	}
	return claim.Id, nil
}

func IsRefreshToken(token string) (id int64, err error) {
	claim, err := ParseJWT(token)
	if err != nil {
		return
	}
	if claim.ExpiresAt < time.Now().Unix() {
		err = errors.ErrTokenExp
		return
	}
	if claim.TokenType != "refresh_token" {
		err = errors.ErrTokenType
		return
	}
	return claim.Id, nil
}
