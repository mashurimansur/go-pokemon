package main

import (
	"go-pokemon/config"
	"go-pokemon/controller"
	"go-pokemon/database"
	"go-pokemon/entity"
	"go-pokemon/repository"
	"go-pokemon/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mtslzr/pokeapi-go"
	"gorm.io/gorm"
)

func main() {
	// Init Database
	configuration := config.New()
	db := database.ConnectionDatabase(configuration)
	// close DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Disable checking for cached data
	pokeapi.CacheSettings.UseCache = false

	//Migrate
	err := db.AutoMigrate(&entity.MyPokemon{})
	if err != nil {
		log.Fatal(err)
	}

	app := gin.Default()

	// static file
	app.Static("/assets", "./public")

	//Load Templates
	app.LoadHTMLGlob("templates/*")

	// Route
	route(app, db)
	log.Fatal(app.Run(":8080"))
}

func route(app *gin.Engine, db *gorm.DB) {
	pokemonRepository := repository.NewPokemonRepository(db)
	pokemonService := service.NewPokemonService(&pokemonRepository)
	pokemonController := controller.NewPokemonController(&pokemonService)
	pokemonController.Route(app)
}
