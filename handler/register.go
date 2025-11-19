package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		d := RegisterRequest{}
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		v, msg := Validate(d)
		if !v {
			w.Write([]byte(msg))
			w.WriteHeader(400)
		}
		b, err := json.Marshal(d)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Write(b)
	}
}
func Validate(d RegisterRequest) (bool, string) {
	if len(d.Password) < 8 || len(d.Password) > 20 {
		return false, "length of password must be between 8 and 20"
	}
	if !strings.HasSuffix(d.Email, "@gmail.com") && !strings.HasSuffix(d.Email, "@163.com") {
		return false, "email must be gmail or 163 mail"
	}
	return true, ""
}
