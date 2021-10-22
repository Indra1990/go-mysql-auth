package dto

type BookCreateRequest struct {
	Title       string `form:"title" `
	Description string `form:"description" `
	UserID      int64  `form:"user_id" `
}
