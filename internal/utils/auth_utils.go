package utils

import (
	"errors"
	"strconv"

	"time"

	"github.com/crisnet-dev/task-api/internal/config"
	"github.com/crisnet-dev/task-api/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(user models.User) (string, string, error) {

	var env = config.GetEnv()
	var SecretKey = []byte(env.SecretKey)

	claims := models.Claims{
		Email: user.Email,
		Name:  user.Name,
		Type:  "access",

		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    env.Issuer,
			Subject:   strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(5) * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", "", err
	}

	refreshTokenClaims := models.Claims{
		Type: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(int(user.ID)),
			Issuer:    env.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(15) * time.Minute)),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(SecretKey)
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshTokenString, nil
}

func ValidateToken(tokenString string) (*models.Claims, error) {
	var env = config.GetEnv()
	var SecretKey = []byte(env.SecretKey)

	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("INVALID_METHOD")
		}
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
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
