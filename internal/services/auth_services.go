package services

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/crisnet-dev/task-api/internal/config"
	"github.com/crisnet-dev/task-api/internal/models"
	"github.com/crisnet-dev/task-api/internal/repository"
	"github.com/crisnet-dev/task-api/internal/utils"
)

func LoginService(user models.User) (string, string, error) {
	if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" {
		return "", "", errors.New("M_C")
	}

	userFounded, err := repository.FindUserByEmail(user.Email)
	if err != nil {
		return "", "", err
	}

	if !utils.CheckPasswordHash(userFounded.Password, user.Password) {
		return "", "", errors.New("INVALID_PASSWORD")
	}

	accessTokenString, refreshTokenString, err := utils.GenerateToken(userFounded)
	if err != nil {
		log.Println(err)
		return "", "", err
	}

	if err := repository.UpdateRefreshToken(userFounded.ID, refreshTokenString); err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func SignupService(user models.User) error {
	if strings.TrimSpace(user.Name) == "" || strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" {
		return errors.New("M_C")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		return err
	}

	err = repository.AddUser(user, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func RefreshTokenService(refreshToken string) (string, string, error) {
	claims, err := utils.ValidateToken(refreshToken)
	if err != nil {
		log.Println(err)

		if strings.Contains(strings.ToUpper(err.Error()), "EXPIRED") {
			return "", "", errors.New("Refresh token expired")
		}

		return "", "", errors.New("Invalid refresh token")
	}

	user_id_string := claims.Subject
	issuer := claims.Issuer
	tokenType := claims.Type

	user_id, err := strconv.Atoi(user_id_string)
	if err != nil {
		return "", "", err
	}

	if tokenType != "refresh" && issuer != config.GetEnv().Issuer {
		return "", "", errors.New("Invalid refresh token")
	}

	_, err = repository.WhereRefreshToken(int(user_id)) // "refreshTokenFounded" Alert!!!!
	if err != nil {
		return "", "", err
	}

	userFounded, err := repository.FindUserByID(int(user_id))
	if err != nil {
		return "", "", err
	}

	newAccessToken, newRefreshToken, err := utils.GenerateToken(userFounded)

	return newAccessToken, newRefreshToken, nil
}
