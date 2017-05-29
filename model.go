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

func NewUser(login Login, user User) (uint64, interface{}) {
	log.Println("Created user: ", login, user)
	return 0, nil
}