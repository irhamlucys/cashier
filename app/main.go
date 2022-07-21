package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/irhamlucys/cashier/bootstrap"
	_menuCategoryHttp "github.com/irhamlucys/cashier/menu_category/delivery/http"
	_menuCategoryRepo "github.com/irhamlucys/cashier/menu_category/repository/mongo"
	_menuCategoryUsecase "github.com/irhamlucys/cashier/menu_category/usecase"
)

func main() {
	// defer func() {
	// 	err := bootstrap.App.Maria.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	router := gin.Default()
	router.Use(cors.Default())

	timeoutContext := time.Duration(bootstrap.App.Config.GetInt("context.timeout")) * time.Second
	mongoDatabase := bootstrap.App.Mongo.Database(bootstrap.App.Config.GetString("mongo.name"))

	menuCategoryRepo := _menuCategoryRepo.NewMongoRepository(*mongoDatabase)
	menuCategoryUsecase := _menuCategoryUsecase.NewMenuCategoryUsecase(menuCategoryRepo, timeoutContext)
	_menuCategoryHttp.NewMenuCategoryHandler(router, menuCategoryUsecase)
}