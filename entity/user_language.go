package entity

type User_Language struct {
	UserID      uint `gorm:"type:integer" json:"user_id"`
	LanguagesID uint `gorm:"type:integer" json:"languages_id"`
}
