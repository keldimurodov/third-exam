package models

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Website   string `json:"website"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdetedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type Users struct {
	Users []*User `json:"users"`
}
