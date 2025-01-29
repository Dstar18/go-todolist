package models

type Notes struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"type:varchar(100);not null" json:"title"`
	IsCompleted int    `gorm:"type:boolean;default:0;not null" json:"is_completed"`
	CreatedAt   string `gorm:"type:datetime;default:null" json:"created_at"`
	UpdatedAt   string `gorm:"type:datetime;default:null" json:"updated_at"`
}
