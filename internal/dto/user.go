package dto

type (
	CreateUserRequest struct {
		Name     *string `json:"name"`
		Email    *string `json:"email"`
		Password *string `json:"password"`
	}
	UpdateUserRequest struct {
		Name     *string `json:"name"`
		Email    *string `json:"email"`
		Password *string `json:"password"`
	}
	UserResponse struct {
		ID    uint   `json:"id"`
		Name  string `json:"name" `
		Email string `json:"email"`
	}
)
