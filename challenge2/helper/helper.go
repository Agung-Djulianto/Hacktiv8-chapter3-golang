package helper

import (
	"chap3-challenge2/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateID() string {
	id := uuid.New()

	return id.String()
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(userID string, roleUser string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role":    roleUser,
	})

	tokenString, err := jwtToken.SignedString([]byte("k4nz3n___h0"))

	return tokenString, err
}

func VerifyToken(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, model.ErrorInvalidToken
		}
		return []byte("k4nz3n___h0"), nil
	})
	return jwtToken, err
}

func IsAdminTrue(role string) bool {
	return role == "admin"
}
