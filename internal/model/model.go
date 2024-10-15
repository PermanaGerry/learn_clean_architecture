package model

type WebResponse[T any] struct {
	Data   T             `json:"data"`
	Paging *PageMetaData `json:"paging,omitempty"`
	Error  string        `json:"error,omitempty"`
}

type PageResponse[T any] struct {
	Data         []T `json:"data"`
	PageMetaData *PageMetaData
}

type PageMetaData struct {
	Page      int `json:"page"`
	Size      int `json:"size"`
	CountData int `json:"count_data"`
	CountPage int `json:"count_page"`
}
