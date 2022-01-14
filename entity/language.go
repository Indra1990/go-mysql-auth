package entity

type Languages struct {
	ID   uint   `gorm:"primary_key:auto_increment" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
	User []User `gorm:"many2many:user_languages"`
}
