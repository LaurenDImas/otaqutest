package main

import (
	"fmt"
	"log"
	"os"
	"otaqutest/handler"
	"otaqutest/hotel"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Can't load configuration!")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connection to database success")

	hotelRepository := hotel.NewRepository(db)
	hotelService := hotel.NewService(hotelRepository)
	hotelHandler := handler.NewhotelHandler(hotelService)

	api := fiber.New()
	router := api.Group("/hotel") // /api

	router.Post("/create", hotelHandler.CreateHotel)
	router.Get("/", hotelHandler.GetAllHotel)

	api.Listen(":8000")
}
