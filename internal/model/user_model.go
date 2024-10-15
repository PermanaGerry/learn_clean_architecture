package model

type UserResponse struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Token    string `json:"token,omitempty"`
	CreateAt int64  `json:"create_at,omitempty"`
	UpdateAt int64  `json:"update_at,omitempty"`
}

type LoginUserRequest struct {
	ID       string `json:"id" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type RegisterUserRequest struct {
	ID       string `json:"id" validate:"required,max=100"`
	Name     string `json:"name" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}
