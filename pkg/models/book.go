package models

// struct for database schema
type BookDetail struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	BookName    string
	AuthorID    uint
	Publication string
}
