package models

type PostComments struct {
	Id        string `json:"id"`
	OwnerId   string `json:"owner_id"`
	PostId    string `json:"post_id"`
	Text      string `json:"text"`
}

type PostCommentsRequest struct {
    OwnerId string `json:"user_id"`
	PostId  string `json:"post_id"`
	Text    string `json:"text"`
}