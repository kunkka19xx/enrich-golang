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

	name            string    `json:"name" bson:"name" binding:"required"`
	email           string    `json:"email" bson:"email" binding:"required"`
	password        string    `json:"password" bson:"password" binding:"required,min=8"`
	passwordConfirm string    `json:"passwordConfirm" bson:"passwordConfirm,omitempty" binding:"required"`
	role            string    `json:"role" bson:"role"`
	verified        bool      `json:"verified" bson:"verified"`
	createdDate     time.Time `json:"created_date" bson:"created_date"`
	updatedDate     time.Time `json:"updated_date" bson:"update_date"`
}

type SingInInput struct {
	email    string
	password string
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
