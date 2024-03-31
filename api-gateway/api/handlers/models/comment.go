package models

type Comments struct {
	Id        string `json:"id"`
	OwnerId   string `json:"owner_id"`
	PostId    string `json:"post_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}