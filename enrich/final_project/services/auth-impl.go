package services

import (
	"context"
	"strings"
	"time"

	"github.com/kunkka19xx/enrich-golang/models"
	// "go/models"

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
	user.CreatedDate = time.Now()
	user.UpdatedDate = user.CreatedDate
	user.Email = strings.ToLower(user.Email)
	user.PasswordConfirm = ""
	user.Verified = true
	user.Role = "user"
	user.Name = ""

	// hashedPassword, _ := utils.hashedPassword(user.password)

	return nil, nil
}

func newAuthService(collection *mongo.Collection, ctx context.Context) AuthService {
	return &AuthServiceImpl{collection, ctx}
}
