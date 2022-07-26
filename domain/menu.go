package domain

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID				uint64
	UUID			uuid.UUID
	MenuID			uint64
	Name			string
	Description		string
	Label			string
	Public			bool
	ImageUrl		*string
	CreatedAt		time.Time
	UpdatedAt		*time.Time
	DeletedAt		*time.Time
}