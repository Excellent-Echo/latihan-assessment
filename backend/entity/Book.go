package entity

type Book struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	UserID int    `json:"user_id"`
}

type BookInput struct {
	Author string `json:"author"`
	Year   int    `json:"year"`
	UserID int    `json:"user_id"`
}
