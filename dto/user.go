package dto

type UserCreateRequest struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

type GetUserResponse struct {
	ID    uint64 `form:"id"`
	Name  string `form:"name"`
	Email string `form:"email"`
}
