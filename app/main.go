package main

import (
	"final_zoom/database"
	"fmt"
	"log"
	"os"

	userhandler "final_zoom/user/delivery/http"
	userrepository "final_zoom/user/repository/postgre"
	userusecase "final_zoom/user/usecase"

	photohandler "final_zoom/photo/delivery/http"
	photorepository "final_zoom/photo/repository/postgre"
	photousecase "final_zoom/photo/usecase"

	sosmedhandler "final_zoom/social_media/delivery/http"
	sosmedrepository "final_zoom/social_media/repository/postgre"
	sosmedusecase "final_zoom/social_media/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	APP_PORT := os.Getenv("APP_PORT")

	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	posgreDB, err := db.DB()
	err = posgreDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := posgreDB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()

	userRepo := userrepository.NewUserRepository(db)
	userUseCase := userusecase.NewUserUseCase(userRepo)
	userhandler.NewUserHandler(router, userUseCase)

	photoRepo := photorepository.NewPhotoRepository(db)
	photoUseCase := photousecase.NewPhotoUseCase(photoRepo)
	photohandler.NewPhotoHandler(router, photoUseCase)

	sosmedRepo := sosmedrepository.NewSosmedRepository(db)
	sosmedUseCase := sosmedusecase.NewSosmedUseCase(sosmedRepo)
	sosmedhandler.NewSosmedHandler(router, sosmedUseCase)

	router.Run(fmt.Sprintf("localhost:%s", APP_PORT))
	log.Println("Berjalan pada port :", APP_PORT)
}
