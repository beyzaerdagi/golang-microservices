package domain

import (
	"fmt"
	"net/http"

	"github.com/beyzaerdagi/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "Beyza", LastName: "Erdagi", Email: "byzerdagi@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exists", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
