package dto

type UserCreateRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email"  binding:"email,required"`
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

type UserUpdateRequest struct {
	ID    uint64
	Name  string `form:"name"`
	Email string `form:"email"`
}
