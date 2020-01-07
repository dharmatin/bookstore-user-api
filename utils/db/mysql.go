package db

import (
	"strings"

	"github.com/dharmatin/bookstore-user-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	if err != nil {
		sqlErr, ok := err.(*mysql.MySQLError)
		if !ok {
			if strings.Contains(err.Error(), errorNoRows) {
				return errors.NewNotFoundError("no record matching")
			}
			return errors.NewInternalServerError("error parsing database error")
		}
		return errors.NewInternalServerError(sqlErr.Message)
	}
	return errors.NewInternalServerError(err.Error())
}
