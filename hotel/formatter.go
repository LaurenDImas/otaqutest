package hotel

import "time"

type HotelFormatter struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	ImageUrl   string    `json:"image_url"`
	StarRating int       `json:"star_rating"`
	Price      int       `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FormatHotel(hotel Hotel) HotelFormatter {
	hotelFormatter := HotelFormatter{}
	hotelFormatter.ID = hotel.ID
	hotelFormatter.Name = hotel.Name
	hotelFormatter.Address = hotel.Address
	hotelFormatter.ImageUrl = hotel.ImageUrl
	hotelFormatter.StarRating = hotel.StarRating
	hotelFormatter.Price = hotel.Price
	hotelFormatter.CreatedAt = hotel.CreatedAt
	hotelFormatter.UpdatedAt = hotel.UpdatedAt
	return hotelFormatter
}

func FormatAllHotel(hotels []Hotel) []HotelFormatter {
	hotelAllFormatter := []HotelFormatter{}

	for _, hotels := range hotels {
		hotelFormatter := FormatHotel(hotels)
		hotelAllFormatter = append(hotelAllFormatter, hotelFormatter)
	}

	return hotelAllFormatter
}
