package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"day-12-restfull-api/helper"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	if len(users) == 0 {
		return helper.Helper(c, http.StatusOK, "Success get all users", []User{})
	}

	return helper.Helper(c, http.StatusOK, "Success get all users", users)
}

// get user by id
func GetUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helper.Helper(c, http.StatusBadRequest, "Invalid argument type", nil)
	}

	for _, user := range users {
		if user.Id == userId {
			return helper.Helper(c, http.StatusOK, "Success get user", users[userId-1])
		}
	}

	return helper.Helper(c, http.StatusNotFound, fmt.Sprintf("User id %v not found", userId), nil)
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helper.Helper(c, http.StatusBadRequest, "Invalid argument type", nil)
	}

	newUsers := make([]User, 0)

	for i, user := range users {
		if user.Id == userId {
			newUsers = append(newUsers, users[i+1:]...)
			newUsers = append(newUsers, users[:i]...)

			users = newUsers

			return helper.Helper(c, http.StatusOK, fmt.Sprintf("Success delete User id %v", userId), newUsers)
		}
	}

	return helper.Helper(c, http.StatusNotFound, fmt.Sprintf("User id %v not found", userId), nil)
}

// update user by id
func UpdateUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helper.Helper(c, http.StatusBadRequest, "Invalid argument type", nil)
	}

	for i, user := range users {
		if user.Id == userId {
			var u User

			err = c.Bind(&u)
			if err != nil {
				return helper.Helper(c, http.StatusBadRequest, err.Error(), nil)
			}

			u.Id = users[i].Id

			users[i] = u

			return helper.Helper(c, http.StatusOK, fmt.Sprintf("Success update User id %v", userId), users)
		}
	}

	return helper.Helper(c, http.StatusNotFound, fmt.Sprintf("User id %v not found", userId), nil)
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		return helper.Helper(c, http.StatusBadRequest, err.Error(), nil)
	}

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	if result := strings.Contains(user.Email, "@"); !result {
		return helper.Helper(c, http.StatusBadRequest, "Invalid email format", user)
	} 

	users = append(users, user)
	return helper.Helper(c, http.StatusOK, "Success create User", user)
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
