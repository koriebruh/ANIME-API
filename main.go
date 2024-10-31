package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"insert_DM/cnf"
	"insert_DM/controller"
	"insert_DM/repository"
	"insert_DM/service"
)

func main() {
	// INSERT DATA FROM CSV
	//cnf.InsertDataCSV()
	db, _ := cnf.InitDB()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository)
	userController := controller.NewUserController(userService)

	animeRepository := repository.NewAnimeRepository()
	animeService := service.NewAnimeService(db, animeRepository)
	animeController := controller.NewAnimeController(animeService)

	r := fiber.New()
	r.Use(cors.New(cors.Config{}))

	r.Post("auth/user/register", userController.Register)
	r.Post("auth/user/login", userController.Login)
	r.Get("auth/user/logout", userController.Logout)

	authorized := r.Group("/", cnf.JWTAuthMiddleware)
	authorized.Get("/favorites", animeController.GetAllAnime)
	authorized.Post("/favorites", animeController.GetAnimeById)
	authorized.Post("/favorites/:anime_id", animeController.AddFavorite)
	authorized.Delete("/favorites/:anime_id", animeController.RemoveFavorite)

	err := r.Listen(":" + cnf.GetConfig().Server.Port)
	if err != nil {
		log.Fatal("connection down")
	}

}
