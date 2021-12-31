package dto

type BookCreateRequest struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	Slug        string `form:"slug" json:"-"`
	UserID      int64  `form:"user_id" json:"-"`
}

type GetBookResponse struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	UserID      int    `json:"user_id"`
	User        User   `json:"user"`
}

type User struct {
	ID    uint64 `json:"iduser"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
