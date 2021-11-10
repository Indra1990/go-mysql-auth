package entity

type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"unique:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null " json:"-"`
	Books    []Book `gorm:"foreignkey:UserID" json:"book,omitempty"`
}

// func (user User) Validate() error {
// 	return validation.ValidateStruct(&user,
// 		validation.Field(&user.Name, validation.Required.Error("Name is Required")),
// 		validation.Field(&user.Email, validation.Required.Error("Email is Required")),
// 	)
// }
