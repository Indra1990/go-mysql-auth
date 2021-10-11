package entity

type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID      string `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID" json:"user"`
}
