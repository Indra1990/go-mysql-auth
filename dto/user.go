package dto

type UserCreateRequest struct {
	Name         string         `form:"name" json:"name"`
	Email        string         `form:"email" json:"email"`
	Password     string         `form:"password" json:"password"`
	LanguageMany []LanguageMany `form:"lang" json:"language"`
}

type LanguageMany struct {
	ID uint `form:"lang_id" json:"lang_id"`
}

type GetAuthUserResponse struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserResponse struct {
	ID        uint64                 `json:"id"`
	Name      string                 `json:"name"`
	Email     string                 `json:"email"`
	Books     []BookResponse         `json:"books,omitempty"`
	Languages []UserLanguageResponse `json:"language,omitempty"`
}

type UserLanguageResponse struct {
	ID   uint   `json:"language_id"`
	Name string `json:"lang_name"`
}

type BookResponse struct {
	ID          uint64 `json:"idbook"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

type GetUserByIDRequest struct {
	ID uint64 `json:"id"`
}

type UserUpdateRequest struct {
	ID    uint64
	Name  string `form:"name"`
	Email string `form:"email"`
}

type GetAuthUserRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
