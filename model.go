package main

import (
	"fmt"
	"log"
)

type Login struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type User struct {
	Nickname    string `json:"nickname,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Description string `json:"description,omitempty"`
}

func UserNew(login Login, user User) {
	log.Println("Created user:", login, user)
}

func UserLogin(login Login) error {
	if login.Email != "" {
		log.Println("Sent mail to:", login.Email)
	} else if login.Phone != "" {
		log.Println("Sent sms to:", login.Phone)
	} else {
		return fmt.Errorf("") // TODO
	}

	return nil
}

func UserVerify(sessionID string, code string) (uint64, error) {
	log.Printf("Verify code %s for session %s\n", code, sessionID)

	return 1234, nil // TODO: mock
}

func UserGet(userID uint64) (User, error) {
	log.Println("Get user_id", userID)

	return User{"test", "male", "Not human"}, nil
}
