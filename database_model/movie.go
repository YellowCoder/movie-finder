package database_model

type Movie struct {
	ID   uint `gorm:"primary_key"`
	Name string
}
