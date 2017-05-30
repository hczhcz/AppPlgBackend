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
