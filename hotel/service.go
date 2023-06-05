package hotel

type Service interface {
	CreateHotel(input HotelInput) (Hotel, error)
	GetAllHotel() ([]Hotel, error)
	DeleteHotel() error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateHotel(input HotelInput) (Hotel, error) {
	hotel := Hotel{}
	hotel.Name = input.Name
	hotel.Address = input.Address
	hotel.ImageUrl = input.ImageUrl
	hotel.StarRating = input.StarRating
	hotel.Price = input.Price

	newhotel, err := s.repository.Save(hotel)
	if err != nil {
		return newhotel, err
	}

	return newhotel, nil
}

func (s *service) GetAllHotel() ([]Hotel, error) {
	hotel, err := s.repository.FindAll()
	if err != nil {
		return hotel, err
	}

	return hotel, nil
}

func (s *service) DeleteHotel() error {
	err := s.repository.Destroy()
	if err != nil {
		return err
	}

	return nil
}
