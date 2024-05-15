package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
    Username string `json:"username"`
    Role     string `json:"role"`
    jwt.StandardClaims
}

func GenerateJWT(username, role string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: username,
        Role:     role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (*Claims, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            return nil, err
        }
        return nil, err
    }

    if !token.Valid {
        return nil, err
    }

    return claims, nil
}
