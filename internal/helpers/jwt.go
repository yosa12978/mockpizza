package helpers

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yosa12978/mockPizza/internal/domain"
)

func NewJWT(usr domain.Usr) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": usr.Username,
			"role":     usr.Role,
			"id":       usr.ID,
		})
	return token.SignedString(signingKey)
}

func ParseJWT(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}
