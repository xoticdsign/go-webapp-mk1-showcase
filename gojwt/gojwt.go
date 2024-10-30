package gojwt

import (
	"go-webapp-mk1-showcase/gorm"

	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GLOBAL VARIABLES

type CustomClaims struct {
	jwt.RegisteredClaims
}

// JWT

func ConfigJWT(userCreds gorm.User) (string, error) {
	claims := CustomClaims{
		jwt.RegisteredClaims{
			Subject:   userCreds.Username,
			Audience:  jwt.ClaimStrings{userCreds.Role},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	secretStr := os.Getenv("JWT_SECRET")
	secret := []byte(secretStr)

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return token, err
	}

	if !token.Valid {
		return token, err
	}

	return token, nil
}
