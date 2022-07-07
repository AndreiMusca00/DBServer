package entities

type Company struct {
	Location_ID string `gorm:"column:location_id"`
	ID          string
	Name        string
	Location    Location
}
