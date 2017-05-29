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

func NewUser(login Login, user User) (Action, interface{}) {
	log.Println("Created user: ", login, user)
	return nil, nil
}
