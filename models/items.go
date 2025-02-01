package models

type Items struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	IdNotes int    `gorm:"type:bigint;not null" json:"id_notes"`
	Name    string `gorm:"type:varchar(100);not null" json:"name"`
	Status  int    `gorm:"type:boolean;default:0;not null" json:"status"`
}
