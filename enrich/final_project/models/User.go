package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SingUpInput struct {
	// name            string
	// email           string
	// password        string
	// passwordConfirm string
	// role            string
	// verified        bool
	// createDate      time.Time
	// updateDate      time.Time

	Name            string    `json:"name" bson:"name" binding:"required"`
	Email           string    `json:"email" bson:"email" binding:"required"`
	Password        string    `json:"password" bson:"password" binding:"required,min=8"`
	PasswordConfirm string    `json:"passwordConfirm" bson:"passwordConfirm,omitempty" binding:"required"`
	Role            string    `json:"role" bson:"role"`
	Verified        bool      `json:"verified" bson:"verified"`
	CreatedDate     time.Time `json:"created_date" bson:"created_date"`
	UpdatedDate     time.Time `json:"updated_date" bson:"update_date"`
}

type SingInInput struct {
	Email    string
	Password string
}

type DBResponse struct {
	Id              primitive.ObjectID
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
	Verified        string
	CreateDate      time.Time
	UpdateDate      time.Time
}

type UserResponse struct {
	Id         primitive.ObjectID
	Name       string
	Email      string
	Role       string
	CreateDate time.Time
	UpdateDate time.Time
}

func FilteredResponse(user *DBResponse) UserResponse {
	return UserResponse{
		Id:         User.Id,
		Email:      User.Email,
		Name:       User.Name,
		Role:       User.Role,
		CreateDate: User.CreateDate,
		UpdateDate: User.UpdateDate,
	}
}
