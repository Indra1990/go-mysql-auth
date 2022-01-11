package dto

type BookCreateRequest struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	Slug        string `form:"slug" json:"-"`
	UserID      int64  `form:"user_id" json:"-"`
}
type BookCreateMultipleRequest struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	Slug        string `form:"slug" json:"-"`
	UserID      int64  `form:"user_id" json:"-"`
}
type BookUpdateRequest struct {
	// ID          int
	Title       string `form:"title"`
	Description string `form:"description"`
	Slug        string `form:"slug" json:"-"`
}

type BookUpdateMultipleRequest struct {
	ID          int    `form:"idbook" json:"idbook" `
	Title       string `form:"title" json:"title"`
	Description string `form:"description" json:"description"`
	Slug        string `form:"slug" json:"-"`
	UserID      int64  `form:"user_id" json:"-"`
}

type GetBookResponse struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	UserID      int    `json:"user_id"`
	User        User   `json:"users"`
}

type User struct {
	ID    uint64 `json:"iduser"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type BookDeleteMultiple struct {
	ID uint `form:"idbook" json:"idbook"`
}
