package services

import (
	"go/models"
)

type AuthService interface {
	signUpUser(*models.SingUpInput) (*models.DBResponse, error)

	signInUser(*models.SingInInput) (*models.DBResponse, error)
}
