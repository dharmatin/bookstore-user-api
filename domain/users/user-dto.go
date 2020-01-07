package users

import (
	"github.com/dharmatin/bookstore-user-api/utils/errors"
	"strings"
)

const (
	ACTIVE_STATUS   = "active"
	INACTIVE_STATUS = "inactive"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	CreatedDate string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user *User) Validate() *errors.RestError {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
