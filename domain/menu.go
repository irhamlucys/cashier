package domain

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID				uint64
	UUID			uuid.UUID
	Name			string
	Description		string
	Label			string
	Public			bool
	ImageUrl		string
	CreatedAt		time.Time
	UpdatedAt		*time.Time
	DeletedAt		*time.Time
}