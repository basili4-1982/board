package board

import (
	"fmt"

	"github.com/google/uuid"

	"board/internal/announcement"
)

type Board struct {
	rep *Repository
}

func NewBoard(rep *Repository) *Board {
	if rep == nil {
		panic("nil repository")
	}
	return &Board{rep: rep}
}

func (b Board) Add(userID uuid.UUID, text string) error {
	err := b.rep.AddAnnouncement(announcement.Announcement{
		ID:     uuid.New(),
		Text:   text,
		UserId: userID,
	},
	)
	if err != nil {
		return fmt.Errorf("board.Add: %w", err)
	}

	return nil
}

func (b Board) List() []announcement.Announcement {
	return b.rep.List()
}

func (b Board) Get(id uuid.UUID) (*announcement.Announcement, error) {
	a, err := b.rep.Get(id)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
