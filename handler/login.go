package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"small_account/model"
)

func Login() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		d := RegisterAndLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		v, msg := Validate(d)
		if !v {
			w.Write([]byte(msg))
			w.WriteHeader(400)
			return
		}
		token, msg, err := model.Login(d.Email, d.Password)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
		if msg != "" {
			w.WriteHeader(400)
			w.Write([]byte(msg))
			return
		}
		w.Write([]byte(token))
	}
}
