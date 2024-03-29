package mongo

import (
	"context"
	"strconv"

	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type menuCategoryMongoRepository struct {
	DB         mongo.Database
	Collection mongo.Collection
}

func NewMenuCategoryMongoRepository(DB mongo.Database) domain.MenuCategoryRepositoryContract {
	return &menuCategoryMongoRepository{
			DB: DB,
			Collection: *DB.Collection("menu_categories"),
		}
}

func (repo *menuCategoryMongoRepository) InsertOne(ctx context.Context, data *domain.MenuCategory) (*domain.MenuCategory, error) {
	var err error

	_, err = repo.Collection.InsertOne(ctx, data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repo *menuCategoryMongoRepository) FindOne(ctx context.Context, id string) (*domain.MenuCategory, error) {
	var menucategory domain.MenuCategory
	var err  error

	intID, _ := strconv.Atoi(id)

	err = repo.Collection.FindOne(ctx, bson.M{
								"$or":
									bson.A{
										bson.M{"uuid": id},
										bson.M{"id": intID},
									},
							},).
							Decode(&menucategory)
	if err != nil {
		return &menucategory, err
	}

	return &menucategory, nil
}