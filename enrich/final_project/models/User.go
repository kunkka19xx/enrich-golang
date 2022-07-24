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
	id              primitive.ObjectID
	name            string
	email           string
	password        string
	passwordConfirm string
	role            string
	verified        string
	createDate      time.Time
	updateDate      time.Time
}

type UserResponse struct {
	id         primitive.ObjectID
	name       string
	email      string
	role       string
	createDate time.Time
	updateDate time.Time
}

func FilteredResponse(user *DBResponse) UserResponse {
	return UserResponse{
		id:         user.id,
		email:      user.email,
		name:       user.name,
		role:       user.role,
		createDate: user.createDate,
		updateDate: user.updateDate,
	}
}
