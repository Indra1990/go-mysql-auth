package dto

type CreateLanguageRequest struct {
	Name string `form:"name" json:"name"`
}

type GetLanguageListResponse struct {
	ID   int             `json:"id"`
	Name string          `json:"name"`
	User []UserLanguages `json:"users,omitempty"`
}

type UserLanguages struct {
	ID    int    `json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
