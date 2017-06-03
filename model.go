package main

import (
	"fmt"
	"log"
)

import "github.com/go-pg/pg"

type Login struct {
	ID    uint64 `json:"-"`
	Email string `sql:",null" json:"email,omitempty"`
	Phone string `sql:",null" json:"phone,omitempty"`
}

func (l Login) IsValid() error {
	if l.Email == "" && l.Phone == "" {
		return fmt.Errorf("The email and phone fields of Login are both empty.")
	}

	if l.Email != "" && l.Phone != "" {
		return fmt.Errorf("Can not determine email or phone is used in Login.")
	}

	return nil
}

type User struct {
	ID          uint64 `json:"-"`
	Nickname    string `json:"nickname,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Description string `json:"description,omitempty"`
}

var DBConn *pg.DB

func init() {
	DBConn = pg.Connect(&pg.Options{
		Addr:     "127.0.0.1:5432",
		User:     "hyclz",
		Password: "666666", // FIXME
		Database: "appplg",
	})
}

func UserNew(login Login, user User) error {
	log.Println("Create user:", login, user)

	tx, err := DBConn.Begin()
	if err != nil {
		return err
	}

	if err := tx.Insert(&login); err != nil {
		tx.Rollback()
		return err
	}

	user.ID = login.ID
	if err := tx.Insert(&user); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	if login.Email != "" {
		log.Println("Sent mail to:", login.Email)
	} else if login.Phone != "" {
		log.Println("Sent sms to:", login.Phone)
	} else {
		// Never reach
	}

	return nil
}

func UserLogin(login Login) error {
	log.Println("Login:", login)

	if login.Email != "" {
		log.Println("Sent mail to:", login.Email)
	} else if login.Phone != "" {
		log.Println("Sent sms to:", login.Phone)
	} else {
		// Never reach
	}

	return nil
}

func UserVerify(sessionID string, code string) (uint64, error) {
	log.Printf("Verify code %s for session %s\n", code, sessionID)

	return 1234, nil // TODO: mock
}

func UserGet(userID uint64) (User, error) {
	log.Println("Get user_id", userID)

	return User{0, "test", "male", "Not human"}, nil
}
