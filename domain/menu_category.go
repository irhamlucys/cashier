package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type MenuCategory struct {
	// ID			uint64
	UUID		uuid.UUID
	BranchID	uint64
	Name		string
	Menus		[]*Menu
	CreatedAt	time.Time
	UpdatedAt	*time.Time
	DeletedAt	*time.Time
}

type MenuCategoryRepository interface {
	InsertOne(ctx context.Context, mc *MenuCategory) (*MenuCategory, error)
	FindOne(ctx context.Context, id string) (*MenuCategory, error)
}

type MenuCategoryUsecase interface {
	InsertOne(ctx context.Context, mc *MenuCategory) (*MenuCategory, error)
	FindOne(ctx context.Context, id string) (*MenuCategory, error)
}

