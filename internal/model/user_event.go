package model

type UserEvent struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	CreateAt int64  `json:"create_at,omitempty"`
	UpdateAt int64  `json:"update_at,omitempty"`
}

func (u *UserEvent) GetId() string {
	return u.ID
}
