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

type PostRequest struct {
	UserID string `json:"user_id"`
	Content string `json:"content"`
	Title   string `json:"title"`
	Likes   int64  `json:"likes"`
	Dislikes int64  `json:"dislikes"`
	Views   int64  `json:"views"`
	Categories string `json:"categories"`
}

type Post struct {
	Id         string `json:"id"`
	UserID     string `json:"userID"`
	Content    string `json:"content"`
	Title      string `json:"title"`
	Likes      int64  `json:"likes"`
	Dislikes   int64  `json:"dislikes"`
	Views      int64  `json:"views"`
	Categories string `json:"categories"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
	Comments   []Comments
}