package services

import (
	"errors"
	"log"
	"strings"

	"github.com/crisnet-dev/models"
	"github.com/crisnet-dev/repository"
	"github.com/crisnet-dev/utils"
)

func LoginService(user models.User) (string, error) {
	if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" {
		return "", errors.New("M_C")
	}

	userFounded, err := repository.FindUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(userFounded.Password, user.Password) {
		return "", errors.New("INVALID_PASSWORD")
	}

	tokenString, err := utils.GenerateToken(userFounded)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
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
