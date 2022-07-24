package services

import (
	"github.com/kunkka19xx/enrich-golang/models"
)

type AuthService interface {
	signUpUser(*models.SingUpInput) (*models.DBResponse, error)

	signInUser(*models.SingInInput) (*models.DBResponse, error)
}
