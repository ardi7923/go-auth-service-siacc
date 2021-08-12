package helper

import (
	"auth/config"
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetToken(uuid string) (token_r string, err_r error) {
	sign := jwt.New(jwt.GetSigningMethod(config.Alg_jwt))
	claims := sign.Claims.(jwt.MapClaims)
	claims["UUID"] = uuid
	token, err := sign.SignedString([]byte(config.Secret_jwt))
	if err != nil {
		return "error", err
	}
	return token, nil
}

func TokenCheck(token_jwt string) (bool, error) {
	token, err := jwt.Parse(token_jwt, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(config.Alg_jwt) != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Secret_jwt), nil
	})
	if token != nil && err == nil && token.Valid {
		return true, nil
	} else {
		return false, err
	}
}

func GetPayloadToken(token_jwt string) (i jwt.MapClaims, err error) {
	token, err := jwt.Parse(token_jwt, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(config.Alg_jwt) != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Secret_jwt), nil
	})
	if !token.Valid {
		return nil, errors.New("Token invalid")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token invalid")
}

func GetPayloadFromHeader(ctx *gin.Context) (i jwt.MapClaims, err error) {
	header := ctx.Request.Header["Authorization"]
	token := strings.Split(header[0], "bearer ")
	claims, err := GetPayloadToken(token[1])
	if err != nil {
		return nil, err
	}
	return claims, nil
}
