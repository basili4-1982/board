package main

import (
	"net/http"
	"os"

	"board/internal/board"
	"board/internal/server"
	"board/internal/user"
	"board/pkg/jwt"
)

func main() {
	srv := server.NewServer()

	secret := "qwerty"
	if len(os.Args) > 1 {
		secret = os.Args[1]
	}

	srv.AddHandler(http.MethodGet, "/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})

	userRep := user.NewRepository()

	regService := user.NewRegistration(userRep)

	j := jwt.NewJwt(secret)

	loginService := user.NewLoginService(userRep, j)

	b := board.NewBoard(board.NewRepository())

	srv.AddHandler(http.MethodPost, "/", user.MakeMiddlewareAuth(j, loginService,
		board.CreateAnnouncementHandler(b),
	))

	srv.AddHandler(http.MethodPost, "/reg", user.MakeRegHandler(regService))

	srv.AddHandler(http.MethodPost, "/login", user.MakeLoginHandler(j, loginService))

	err := srv.Run(":8080")
	if err != nil {
		panic(err)
	}
}
