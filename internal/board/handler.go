package board

import (
	"net/http"

	"github.com/google/uuid"
)

func CreateAnnouncementHandler(b *Board) func(w http.ResponseWriter, r *http.Request) {
	if b == nil {
		panic("nil board service")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		r.Context()

		userID := uuid.New()
		text := ""

		err := b.Add(userID, text)
		if err != nil {
			return
		}

		w.Write([]byte("создание поста"))
	}
}
