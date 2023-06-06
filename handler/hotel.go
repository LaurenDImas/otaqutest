package handler

import (
	"log"
	"net/http"
	"os"
	"otaqutest/hotel"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type hotelHandler struct {
	hotelService hotel.Service
}

func NewhotelHandler(hotelService hotel.Service) *hotelHandler {
	return &hotelHandler{hotelService}
}

func (h *hotelHandler) CreateHotel(c *fiber.Ctx) error {
	var input hotel.HotelInput
	err := godotenv.Load()
	if err != nil {
		panic("Can't load configuration!")
	}

	res, err := http.Get(os.Getenv("LINK_SCRAPPING"))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	errDelete := h.hotelService.DeleteHotel()
	if errDelete != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.ErrBadRequest)
	}

	replacer := strings.NewReplacer("IDR", "", ".", "")
	doc.Find("#content-hotel").Children().Each(func(i int, sel *goquery.Selection) {
		row := input
		row.Name = sel.Find("h3").Text()
		row.Address = sel.Find(".loct").Text()
		row.ImageUrl, _ = sel.Find(".img-hotel").Attr("src")
		row.StarRating = sel.Find(".star-hotel").Length()
		priceString := sel.Find(".price-hotel").Text()
		replacePriceString := replacer.Replace(strings.ReplaceAll(priceString, " ", ""))
		row.Price, _ = strconv.Atoi(replacePriceString)

		// Save data to table hotel
		_, err := h.hotelService.CreateHotel(row)
		if err != nil {
			c.Status(fiber.ErrBadRequest.Code).JSON(fiber.ErrBadRequest)
			return
		}
	})

	return c.Status(fiber.StatusOK).JSON("Data Berhasil Disimpan")
}

func (h *hotelHandler) GetAllHotel(c *fiber.Ctx) error {
	hotels, err := h.hotelService.GetAllHotel()
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.ErrBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(hotel.FormatAllHotel(hotels))
}
