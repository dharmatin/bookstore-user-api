package users

import (
	"encoding/json"
)

type PublicUser struct {
	Id          int64  `json:"id"`
	CreatedDate string `json:"date_created"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	CreatedDate string `json:"date_created"`
	Status      string `json:"status"`
}

func (users Users) Marshal(isPublicUser bool) []interface{} {
	results := make([]interface{}, len(users))
	for index, user := range users {
		results[index] = user.Marshall(isPublicUser)
	}

	return results
}

func (user *User) Marshall(isPublicUser bool) interface{} {
	if isPublicUser {
		return PublicUser{
			Id:          user.Id,
			CreatedDate: user.CreatedDate,
			Status:      user.Status,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
