package utils

import (
	"crypto/rsa"
	"github.com/go-open-auth/global"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type TokenClaims struct {
	UserID string                 `json:"userId"`
	Data   map[string]interface{} `json:"data"`
	jwt.RegisteredClaims
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func GenerateJWT(userId string, payloadData map[string]interface{}) (*Token, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv(global.TokenPrivateKey)))
	if err != nil {
		return nil, err
	}

	accessToken, err := generateToken(userId, payloadData, global.AccessTokenExpire, privateKey)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateToken(userId, payloadData, global.RefreshTokenExpire, privateKey)
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func generateToken(userId string, payloadData map[string]interface{}, duration time.Duration, privateKey *rsa.PrivateKey) (string, error) {
	claims := TokenClaims{
		UserID: userId,
		Data:   payloadData,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "open-auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
