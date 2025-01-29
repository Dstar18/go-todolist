package models

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Firstname string `gorm:"type:varchar(100);not null" json:"firstname"`
	Lastname  string `gorm:"type:varchar(100);not null" json:"lastname"`
	Email     string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string `gorm:"type:varchar(100);not null" json:"password"`
	CreatedAt string `gorm:"type:datetime;default:null" json:"created_at"`
	UpdatedAt string `gorm:"type:datetime;default:null" json:"updated_at"`
}
