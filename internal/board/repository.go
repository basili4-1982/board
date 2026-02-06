package board

import (
	"github.com/google/uuid"

	"board/internal/announcement"
)

type Repository struct {
	announcements []announcement.Announcement
}

func NewRepository() *Repository {
	return &Repository{
		announcements: make([]announcement.Announcement, 0),
	}
}

func (r *Repository) AddAnnouncement(announcement announcement.Announcement) error {
	r.announcements = append(r.announcements, announcement)
	return nil
}

func (r *Repository) List() []announcement.Announcement {
	return r.announcements
}

func (r *Repository) Get(id uuid.UUID) (announcement.Announcement, error) {
	for _, a := range r.announcements {
		if a.ID == id {
			return a, nil
		}
	}

	return announcement.Announcement{}, NotFoundAnnouncement
}
