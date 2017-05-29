package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user_new", userNew)
	http.HandleFunc("/user_login", userLogin)
	http.HandleFunc("/user_verify", userVerify)
	http.HandleFunc("/user_get", userGet)
	http.HandleFunc("/user_update", userUpdate)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

type response struct {
	Error uint64      `json:"error"`
	Data  interface{} `json:"data"`
}

func newResponse(code uint64, data interface{}) response {
	if data == nil {
		data = map[string]string{} // created an empty dict
	}
	return response{code, data}
}

func userNew(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Login
		User
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var req request
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println(err)
		return
	}

	code, ret := NewUser(req.Login, req.User)

	res, err := json.Marshal(newResponse(code, ret))
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

func userLogin(w http.ResponseWriter, r *http.Request) {

}

func userVerify(w http.ResponseWriter, r *http.Request) {

}

func userGet(w http.ResponseWriter, r *http.Request) {

}

func userUpdate(w http.ResponseWriter, r *http.Request) {

}
