package http

import (
	"context"
	"net/http"

	"lucy/cashier/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ResponseError struct {
	Message string `json:"message"`
}

type MenuCategoryHandler struct {
	MenuCategoryUsecase domain.MenuCategoryUsecase
}

func NewMenuCategoryHandler(router *gin.Engine, mc domain.MenuCategoryUsecase) {
	handler := &MenuCategoryHandler{
		MenuCategoryUsecase: mc,
	}

	router.POST("/menu-categories", handler.InsertOne)
}

func (mch *MenuCategoryHandler) InsertOne(c *gin.Context) {
	var (
		menucategory domain.MenuCategory
		err error
	)

	err = c.BindJSON(&menucategory)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var ok bool
	if ok, err = isRequestValid(&menucategory); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	result, err := mch.MenuCategoryUsecase.InsertOne(ctx, &menucategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, result)
}

func isRequestValid(mc *domain.MenuCategory) (bool, error) {
	validate := validator.New()
	err := validate.Struct(mc)
	if err != nil {
		return false, err
	}
	return true, nil
}