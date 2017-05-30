package main

import (
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
		return ActionInvalidRequest{"The email and phone fields of Login are both empty."}
	}
	return ActionOK{}
}

func UserVerify(userID string, code string) interface{} {
	log.Printf("Verify code %s for user_id %s\n", code, userID)

	type response struct {
		UserID string `json:"user_id"`
	}
	return response{userID}
}
