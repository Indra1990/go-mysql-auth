package dto

type UserCreateRequest struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

type GetUserResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetUserByIDRequest struct {
	ID uint64 `json:"id"`
}
