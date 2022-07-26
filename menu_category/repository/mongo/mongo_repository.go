package mongo

import (
	"context"

	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type mongoRepository struct {
	DB         mongo.Database
	Collection mongo.Collection
}

const (
	timeFormat     = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
	collectionName = "menu_categories"
)

func NewMongoRepository(DB mongo.Database) domain.MenuCategoryRepository {
	return &mongoRepository{
			DB,
			*DB.Collection(collectionName),
		}
}

func (m *mongoRepository) InsertOne(ctx context.Context, menucategory *domain.MenuCategory) (*domain.MenuCategory, error) {
	var (
		err error
	)

	_, err = m.Collection.InsertOne(ctx, menucategory)
	if err != nil {
		return menucategory, err
	}

	return menucategory, nil
}

func (m *mongoRepository) FindOne(ctx context.Context, id string) (*domain.MenuCategory, error) {
	var (
		menucategory domain.MenuCategory
		err  error
	)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &menucategory, err
	}

	err = m.Collection.FindOne(ctx, bson.M{"_id": idHex}).Decode(&menucategory)
	if err != nil {
		return &menucategory, err
	}

	return &menucategory, nil
}