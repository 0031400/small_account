package middleware

import (
	"log"
	"net/http"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				log.Panicln(err)
				w.WriteHeader(500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
