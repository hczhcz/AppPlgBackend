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

func jsonHandler(fn func([]byte) interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		res, err := json.Marshal(newResponse(fn(body)))
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

func userNew(body []byte) interface{} {
	type request struct {
		Login
		User
	}

	var req request
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Println(err)
		return ActionInternalError{}
	}

	return UserNew(req.Login, req.User)
}

func userLogin(body []byte) interface{} {
	type request struct {
		Login
	}

	var req request
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Println(err)
		return ActionInternalError{}
	}

	return UserLogin(req.Login)
}

func userVerify(body []byte) interface{} {
	return nil
}

func userGet(body []byte) interface{} {
	return nil
}

func userUpdate(body []byte) interface{} {
	return nil
}
