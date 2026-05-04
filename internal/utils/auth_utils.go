package utils

import (
	"errors"
	"time"

	"github.com/crisnet-dev/task-api/internal/config"
	"github.com/crisnet-dev/task-api/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(user models.User) (string, string, error) {

	var env = config.GetEnv()
	var SecretKey = []byte(env.SecretKey)

	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"name":  user.Name,
		"iss":   env.Issuer,
		"type":  "access",
		"exp":   time.Now().Add(time.Duration(5) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", "", err
	}

	refreshTokenClaims := jwt.MapClaims{
		"sub":  user.ID,
		"iss":  env.Issuer,
		"type": "refresh",
		"exp":  time.Now().Add(time.Duration(5) * time.Minute).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(SecretKey)
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshTokenString, nil
}

// ?
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	var env = config.GetEnv()
	var SecretKey = []byte(env.SecretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("INVALID_METHOD")
		}
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("INVALID_TOKEN")
}

func ValidateEmail() {}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPasswordHash(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
