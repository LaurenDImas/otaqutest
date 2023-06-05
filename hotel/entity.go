package hotel

import "time"

type Hotel struct {
	ID         int
	Name       string
	Address    string
	ImageUrl   string
	StarRating int
	Price      int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Hotel) TableName() string {
	return "hotel"
}
