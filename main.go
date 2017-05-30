package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user_new", jsonHandler(userNew))
	http.HandleFunc("/user_login", jsonHandler(userLogin))
	http.HandleFunc("/user_verify", jsonHandler(userVerify))
	http.HandleFunc("/user_get", jsonHandler(userGet))
	http.HandleFunc("/user_update", jsonHandler(userUpdate))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln(err)
	}
}

type Request struct {
	SessionID string          `json:"session_id"`
	Data      json.RawMessage `json:"data"`
}

type Response struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

// Note: We will have a session table in database
//       which stores login status and progress of login/register.
//       So, session ID would be fine to track a transaction instead of user ID

// func lookupUserIDBySessionID(sessionID string) uint64 {
// 	return 1000
// }

// func getUserID(r *http.Request) uint64 {
// 	if cookieSessionID, err := r.Cookie("session_id"); err == nil {
// 		return lookupUserIDBySessionID(cookieSessionID.Value)
// 	} else {
// 		return 0
// 	}
// }

func jsonHandler(fn func(string, json.RawMessage) (string, interface{})) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			// Note: It is not a legal command
			return
		}

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		var req Request
		if err := json.Unmarshal(body, &req); err != nil {
			log.Println(err)
			return
		}

		var res Response
		res.Action, res.Data = fn(req.SessionID, req.Data)

		returnBody, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
			return
		}

		if _, err = w.Write(returnBody); err != nil {
			log.Println(err)
			return
		}
	}
}

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

	if req.Login.Email == "" && req.Login.Phone == "" {
		return actionInvalidRequest("The email and phone fields of Login are both empty.")
	}
	if req.Login.Email != "" && req.Login.Phone != "" {
		return actionInvalidRequest("Can not determine email or phone is used in Login.")
	}

	if err := UserNew(req.Login, req.User); err != nil {
		log.Println(err)
		return actionInvalidRequest("???.") // TODO
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

	if req.Login.Email == "" && req.Login.Phone == "" {
		return actionInvalidRequest("The email and phone fields of Login are both empty.")
	}
	if req.Login.Email != "" && req.Login.Phone != "" {
		return actionInvalidRequest("Can not determine email or phone is used in Login.")
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
