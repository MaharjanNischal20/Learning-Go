package main

import (
	"fmt"
	"log"
	"os"

	"example.com/auth/internal/auth"
	"example.com/auth/internal/database"
	apphttp "example.com/auth/internal/http"
	"example.com/auth/internal/user"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := database.Connect(dsn)

	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	userRepository := user.NewRepository(db)
	jwtService := auth.NewJWTService(os.Getenv("JWT_SECRET"))
	authService := auth.NewService(userRepository, jwtService)
	authHandler := auth.NewHandler(authService)
	router := apphttp.NewRouter(authHandler, jwtService)
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
