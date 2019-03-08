package userService

import (
	"../../../model"

	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user, err := CreateUser(&model.User{
		Username: "test1",
		Password: "admin1234",
	})
	fmt.Println("user:", user, "err:", err)
}

func TestReadUser(t *testing.T) {
	user, err := ReadUser(6)
	fmt.Println("user:", user, "err:", err)
}

func TestUpdateUser(t *testing.T) {
	user, err := UpdateUser(&model.User{
		ID:        6,
		Firstname: "xiang",
		Lastname:  "tian",
		Avatar:    "hello world",
	})
	fmt.Println("user:", user, "err:", err)
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser(6)
	fmt.Println("err:", err)
}

func TestReadUsers(t *testing.T) {
	users, total, err := ReadUsers("admin", 0, 0, "", 0)
	fmt.Println("users:", users, "total:", total, "err:", err)
}

func TestReadUserByUsername(t *testing.T) {
	user, err := ReadUserByUsername("user3")
	fmt.Println(user, err)
}
