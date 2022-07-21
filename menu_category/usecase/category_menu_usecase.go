package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/irhamlucys/cashier/domain"
)

type menuCategoryUsecase struct {
	menuCategoryRepo	domain.MenuCategoryRepository
	contextTimeout		time.Duration
}

func NewMenuCategoryUsecase(mc domain.MenuCategoryRepository, to time.Duration) domain.MenuCategoryUsecase {
	return &menuCategoryUsecase{
		menuCategoryRepo: mc,
		contextTimeout: to,
	}
}

func (menuCategory *menuCategoryUsecase) InsertOne(c context.Context, mc *domain.MenuCategory) (*domain.MenuCategory, error) {

	ctx, cancel := context.WithTimeout(c, menuCategory.contextTimeout)
	defer cancel()

	mc.UUID = uuid.New()
	mc.CreatedAt = time.Now()

	res, err := menuCategory.menuCategoryRepo.InsertOne(ctx, mc)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (menuCategory *menuCategoryUsecase) FindOne(c context.Context, id string) (*domain.MenuCategory, error) {

	ctx, cancel := context.WithTimeout(c, menuCategory.contextTimeout)
	defer cancel()

	res, err := menuCategory.menuCategoryRepo.FindOne(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}