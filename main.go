package main

import (
	"fmt"
	"log"
	"otaqutest/handler"
	"otaqutest/hotel"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/otaqu_test?charset=utf8mb4&parseTime=True&loc=Local"
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
