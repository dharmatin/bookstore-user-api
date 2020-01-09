package users

import (
	"fmt"

	"github.com/dharmatin/bookstore-user-api/datasources/mysql/db"
	"github.com/dharmatin/bookstore-user-api/logger"
	dbUtil "github.com/dharmatin/bookstore-user-api/utils/db"
	"github.com/dharmatin/bookstore-user-api/utils/errors"
)

const (
	queryInsert     = "INSERT INTO users (first_name, last_name, email, created_date, status, password) VALUES (?, ?, ?, ?, ?, ?)"
	queryGetUser    = "SELECT id, first_name, last_name, email, created_date, status, password FROM users WHERE id=?"
	queryUpdate     = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	queryDelete     = "DELETE FROM users WHERE id=?"
	queryFindStatus = "SELECT id, first_name, last_name, email, created_date, status FROM users WHERE status=?"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	stmt, err := db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when prepare statement", err)
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedDate, &user.Status, &user.Password); err != nil {
		logger.Error("error scan struct", err)
		return dbUtil.ParseError(err)
	}

	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := db.Client.Prepare(queryInsert)
	if err != nil {
		logger.Error("error when prepare statement", err)
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedDate, user.Status, user.Password)
	if err != nil {
		return dbUtil.ParseError(err)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	user.Id = userId
	return nil
}

func (newUser *User) Update() *errors.RestError {
	stmt, err := db.Client.Prepare(queryUpdate)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(newUser.FirstName, newUser.LastName, newUser.Email, newUser.Id)
	if err != nil {
		return dbUtil.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := db.Client.Prepare(queryDelete)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	if _, err = stmt.Exec(user.Id); err != nil {
		return dbUtil.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) (Users, *errors.RestError) {
	stmt, err := db.Client.Prepare(queryFindStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedDate, &user.Status); err != nil {
			return nil, dbUtil.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No user matching with status %s", status))
	}

	return results, nil
}
