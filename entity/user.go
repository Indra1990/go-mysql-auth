package entity

type User struct {
	ID        uint64      `gorm:"primary_key:auto_increment" json:"id"`
	Name      string      `gorm:"type:varchar(255)" json:"name"`
	Email     string      `gorm:"unique:varchar(255)" json:"email"`
	Password  string      `gorm:"->;<-;not null " json:"-"`
	Books     []Book      `gorm:"foreignkey:UserID" json:"book,omitempty"`
	Languages []Languages `gorm:"many2many:user_languages"`
}
