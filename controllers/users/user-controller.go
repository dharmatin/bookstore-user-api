package users

import (
	"net/http"
	"strconv"

	"github.com/dharmatin/bookstore-user-api/domain/users"
	"github.com/dharmatin/bookstore-user-api/services"
	"github.com/dharmatin/bookstore-user-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))

	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(getXPublic(c)))
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(getXPublic(c)))
}

func UpdateUser(c *gin.Context) {
	var user users.User
	userId, err := getUserId(c.Param("user_id"))

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	isPartialUpdate := c.Request.Method == http.MethodPatch
	user.Id = userId
	result, updateErr := services.UserService.UpdateUser(isPartialUpdate, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(getXPublic(c)))
}

func DeleteUser(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))

	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	if err := services.UserService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func SearchUser(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UserService.FindByStatus(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshal(getXPublic(c)))
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid request")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, err := services.UserService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(getXPublic(c)))
}

func getUserId(urlParam string) (int64, *errors.RestError) {
	userId, err := strconv.ParseInt(urlParam, 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("User id must be a number")
		return 0, restErr
	}
	return userId, nil
}

func getXPublic(c *gin.Context) bool {
	return c.GetHeader("X-Public") == "true"
}
