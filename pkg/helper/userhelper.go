package helper

import (
	"time"

	"github.com/ashikask2002/ecomerce.git/pkg/utils/models"
	"github.com/golang-jwt/jwt/v5"
)

type authCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateTokenClients(user models.UserDetailsResponse) (string, error) {
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Hour * 48))
	issuedAt := jwt.NewNumericDate(time.Now())

	claims := &authCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "client",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("super-secret-key"))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
