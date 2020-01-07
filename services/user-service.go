package services

import (
	"github.com/dharmatin/bookstore-user-api/domain/users"
	"github.com/dharmatin/bookstore-user-api/utils/crypto"
	"github.com/dharmatin/bookstore-user-api/utils/date"
	"github.com/dharmatin/bookstore-user-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.ACTIVE_STATUS
	user.CreatedDate = date.GetNowDB()
	user.Password = crypto.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(isPartialUpdate bool, user users.User) (*users.User, *errors.RestError) {
	current, getErr := GetUser(user.Id)
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if getErr != nil {
		return nil, getErr
	}
	if !isPartialUpdate {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	} else {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	}
	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteUser(userId int64) *errors.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}

func GetUser(userId int64) (*users.User, *errors.RestError) {
	user := &users.User{Id: userId}
	if userId <= 0 {
		return nil, errors.NewBadRequestError("Invalid User Id")
	}

	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func FindByStatus(status string) (users.Users, *errors.RestError) {
	user := &users.User{}
	return user.FindByStatus(status)
}
