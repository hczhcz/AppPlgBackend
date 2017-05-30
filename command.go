package main

import (
	"encoding/json"
	"log"
)

func userNew(sessionID string, data json.RawMessage) (string, interface{}) {
	type request struct {
		Login
		User
	}
	type response struct {
	}

	var req request
	if err := json.Unmarshal(data, &req); err != nil {
		log.Println(err)
		return actionInvalidRequest("Invalid user_new request")
	}

	if err := req.Login.IsValid(); err != nil {
		return actionInvalidRequest(err.Error())
	}

	if err := UserNew(req.Login, req.User); err != nil {
		log.Println(err)
		switch err.(type) { // TODO
		case *ErrorDuplicatedEmail:
			return actionDuplicatedEmail() // TODO
		default:
			return actionInvalidRequest("???.") // TODO
		}
	}

	return "", response{}
}

func userLogin(sessionID string, data json.RawMessage) (string, interface{}) {
	type request struct {
		Login
	}
	type response struct {
	}

	var req request
	if err := json.Unmarshal(data, &req); err != nil {
		log.Println(err)
		return actionInvalidRequest("Invalid user_login request")
	}

	if err := req.Login.IsValid(); err != nil {
		return actionInvalidRequest(err.Error())
	}

	if err := UserLogin(req.Login); err != nil {
		log.Println(err)
		return actionInvalidRequest("???.") // TODO
	}

	return "", response{}
}

func userVerify(sessionID string, data json.RawMessage) (string, interface{}) {
	type request struct {
		Code string `json:"code"`
	}
	type response struct {
		UserID uint64 `json:"user_id"`
	}

	var req request
	if err := json.Unmarshal(data, &req); err != nil {
		log.Println(err)
		return actionInvalidRequest("Invalid user_verify request")
	}

	userID, err := UserVerify(sessionID, req.Code)
	if err != nil {
		log.Println(err)
		return actionInvalidRequest("???.") // TODO
	}

	return "", response{userID}
}

func userGet(sessionID string, data json.RawMessage) (string, interface{}) {
	type request struct {
		UserID uint64 `json:"user_id"`
	}
	type response struct {
		User User `json:"user"`
	}

	var req request
	if err := json.Unmarshal(data, &req); err != nil {
		log.Println(err)
		return actionInvalidRequest("Invalid user_get request")
	}

	user, err := UserGet(req.UserID)
	if err != nil {
		log.Println(err)
		return actionInvalidRequest("???.") // TODO
	}

	return "", response{user}
}

func userUpdate(sessionID string, data json.RawMessage) (string, interface{}) {
	type request struct {
	}
	type response struct {
	}

	return "", response{}
}
