package main

import (
	"log"
)

type Login struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type User struct {
	Gender      string `json:"gender,omitempty"`
	Description string `json:"description,omitempty"`
}

func UserNew(login Login, user User) interface{} {
	log.Println("Created user: ", login, user)
	return ActionOK{}
}

func UserLogin(login Login) interface{} {
	if login.Email != "" {
		log.Println("Sent mail to: ", login.Email)
	} else if login.Phone != "" {
		log.Println("Sent sms to: ", login.Phone)
	} else {
		return ActionInternalError{}
	}
	return ActionOK{}
}

func UserVerify(code string) interface{} {
	return nil
}
