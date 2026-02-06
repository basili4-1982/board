package user

import (
	"encoding/json"
	"io"
	"net/http"

	"board/pkg/jwt"
)

func MakeLoginHandler(jwt *jwt.Jwt, loginService *LoginService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		var l Login

		if err := json.Unmarshal(body, &l); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		s, err := loginService.Login(l.Login, l.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte(s))
	}
}
