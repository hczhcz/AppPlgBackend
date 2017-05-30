package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user_new", jsonHandler(userNew, true))
	http.HandleFunc("/user_login", jsonHandler(userLogin, true))
	http.HandleFunc("/user_verify", jsonHandler(userVerify, false))
	http.HandleFunc("/user_get", jsonHandler(userGet, false))
	http.HandleFunc("/user_update", jsonHandler(userUpdate, false))

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

type response struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

func newResponse(data interface{}) response {
	// data should not be nil
	if data == nil {
		data = ActionInternalError{}
	}

	res := response{}

	if d, ok := data.(Action); ok {
		res.Action = d.Action()
		res.Data = d
	} else {
		res.Data = data
	}

	return res
}

func lookupUserIDBySessionID(sessionID string) uint64 {
	return 1000
}

func getUserID(r *http.Request) uint64 {
	if cookieSessionID, err := r.Cookie("session_id"); err == nil {
		return lookupUserIDBySessionID(cookieSessionID.Value)
	} else {
		return 0
	}
}

func jsonHandler(fn func(uint64, []byte) interface{}, acceptInvalidSessionID bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}

		userID := getUserID(r)
		if userID == 0 && !acceptInvalidSessionID {
			return
		}

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		res, err := json.Marshal(newResponse(fn(userID, body)))
		if err != nil {
			log.Println(err)
			return
		}

		_, err = w.Write(res)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func userNew(_ uint64, body []byte) interface{} {
	type request struct {
		Login
		User
	}

	var req request
	if err := json.Unmarshal(body, &req); err != nil {
		log.Println(err)
		return ActionInvalidRequest{"Invalid user_new request"}
	}

	return UserNew(req.Login, req.User)
}

func userLogin(_ uint64, body []byte) interface{} {
	type request struct {
		Login
	}

	var req request
	if err := json.Unmarshal(body, &req); err != nil {
		log.Println(err)
		return ActionInvalidRequest{"Invalid user_login request"}
	}

	return UserLogin(req.Login)
}

func userVerify(userID uint64, body []byte) interface{} {
	type request struct {
		Code string `json:"code"`
	}

	var req request
	if err := json.Unmarshal(body, &req); err != nil {
		log.Println(err)
		return ActionInvalidRequest{"Invalid user_verify request"}
	}

	return UserVerify(userID, req.Code)
}

func userGet(_ uint64, body []byte) interface{} {
	type request struct {
		UserID uint64 `json:"user_id"`
	}

	var req request
	if err := json.Unmarshal(body, &req); err != nil {
		log.Println(err)
		return ActionInvalidRequest{"Invalid user_get request"}
	}

	return UserGet(req.UserID)
}

func userUpdate(userID uint64, body []byte) interface{} {
	return nil
}
