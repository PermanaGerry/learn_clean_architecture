package converter

import (
	"learn-hexagonal-architecture/internal/entity"
	"learn-hexagonal-architecture/internal/model"
)

func UserToTokenResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Token:    user.Token,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
	}
}

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
	}
}

func UserToEvent(user *entity.User) *model.UserEvent {
	return &model.UserEvent{
		ID:       user.ID,
		Name:     user.Name,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
	}
}
