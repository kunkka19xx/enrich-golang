package services

import (
	"context"
	"strings"
	"time"

	"github.com/kunkka19xx/enrich-golang/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

// signInUser implements AuthService
func (uc *AuthServiceImpl) signInUser(*models.SingInInput) (*models.DBResponse, error) {
	return nil, nil
}

// signUpUser implements AuthService
func (uc *AuthServiceImpl) signUpUser(user *models.SingUpInput) (*models.DBResponse, error) {
	user.createdDate = time.Now()
	user.updatedDate = user.createdDate
	user.email = strings.ToLower(user.email)
	user.passwordConfirm = ""
	user.verified = true
	user.role = "user"

	hashedPassword, _ := utils.hashedPassword(user.password)

	return nil, nil
}

func newAuthService(collection *mongo.Collection, ctx context.Context) AuthService {
	return &AuthServiceImpl{collection, ctx}
}
