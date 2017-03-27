package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"message": http.StatusText(http.StatusInternalServerError),
				})
			}
		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
