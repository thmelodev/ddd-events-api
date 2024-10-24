package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thmelodev/ddd-events-api/src/providers/config"
)

type TokenClaims struct {
	Email  string  `json:"email"`
	UserID string  `json:"userId"`
	Exp    float64 `json:"exp"`
}

func GenerateToken(email string, userId string, cf *config.Config) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Duration(cf.Jwt.Expiration)).Unix(),
	})

	return token.SignedString([]byte(cf.Jwt.SecretKey))
}

func ValidateToken(tokenString string, cf *config.Config) (*TokenClaims, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not parse token")
		}
		return []byte(cf.Jwt.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("invalid token claims")
	}

	email := claims["email"].(string)
	userId := claims["userId"].(string)
	exp := claims["exp"].(float64)

	return &TokenClaims{
		Email:  email,
		UserID: userId,
		Exp:    exp,
	}, nil
}
