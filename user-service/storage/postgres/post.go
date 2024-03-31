package postgres

import (
	"log"
	u "third-exam/user-service/genproto/user"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

// NewPostRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(post *u.User) (*u.User, error) {

	idd := uuid.NewString()
	post.Id = idd

	var re u.User

	query := `INSERT INTO Users(
			id, 
			userID, 
			content, 
			title, 
			likes,
			dislikes,
			views, 
        	categories) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
		RETURNING 
			id, 
			userID, 
			content, 
			title, 
			likes,
			dislikes,
			views, 
        	categories,
			created_at, 
        	updeted_at`
	er := r.db.QueryRow(query, post.Id, post.UserID, post.Content, post.Title, post.Likes, post.Dislikes, post.Views, post.Categories).Scan(
		&re.Id,
		&re.UserID,
		&re.Content,
		&re.Title,
		&re.Likes,
		&re.Dislikes,
		&re.Views,
		&re.Categories,
		&re.CreatedAt,
		&re.UpdetedAt,
	)
	if er != nil {
		log.Println(er)
	}

	updated_query := `
	UPDATE 
		Posts 
	SET 
		views = views + 1
	WHERE
		id = $1
	RETURNING
		id,
		userID, 
		content, 
		title,
		likes,
    	dislikes,
		views,
    	categories, 
		created_at,
		updeted_at`

	var respUser p.Post
	errr := r.db.QueryRow(updated_query, post.Id).Scan(
		&respUser.Id,
		&respUser.UserID,
		&respUser.Content,
		&respUser.Title,
		&respUser.Likes,
		&respUser.Dislikes,
		&respUser.Views,
		&respUser.Categories,
		&respUser.CreatedAt,
		&respUser.UpdetedAt,
	)

	if errr != nil {
		return nil, errr
	}

	return &respUser, nil
}

func (r *postRepo) GetPost(pp *p.GetPostRequest) (*p.Post, error) {

	updated_query := `
	UPDATE 
		Posts 
	SET 
		views = views + 1
	WHERE
		id = $1
	and
		deleted_at IS NULL
	RETURNING
		id,
		userID, 
		content, 
		title,
		likes,
    	dislikes,
		views,
    	categories, 
		created_at,
		updeted_at`

	var respUser p.Post
	errr := r.db.QueryRow(updated_query, pp.Id).Scan(
		&respUser.Id,
		&respUser.UserID,
		&respUser.Content,
		&respUser.Title,
		&respUser.Likes,
		&respUser.Dislikes,
		&respUser.Views,
		&respUser.Categories,
		&respUser.CreatedAt,
		&respUser.UpdetedAt,
	)

	if errr != nil {
		log.Println("No such user found?")
	}

	return &respUser, nil
}

func (r *postRepo) GetAllPosts(pp *p.GetAllRequest) (*p.GetAllResponse, error) {
	var allPost p.GetAllResponse
	query := `
	SELECT
		id,
		userID,
		content,
		title,
		likes,
        dislikes,
		views,
        categories,
		created_at,
		updeted_at
	FROM 
		Posts
	WHERE 
		deleted_at IS NULL
	LIMIT $1
	OFFSET $2`

	offset := pp.Limit * (pp.Page - 1)

	rows, err := r.db.Query(query, pp.Limit, offset)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var pr p.Post

		err := rows.Scan(
			&pr.Id,
			&pr.UserID,
			&pr.Content,
			&pr.Title,
			&pr.Likes,
			&pr.Dislikes,
			&pr.Views,
			&pr.Categories,
			&pr.CreatedAt,
			&pr.UpdetedAt)

		if err != nil {
			return nil, err
		}

		allPost.Posts = append(allPost.Posts, &pr)
	}
	return &allPost, nil
}

func (r *postRepo) UpdatePost(prr *p.Post) (*p.Post, error) {
	query := `
	UPDATE
		posts
	SET
		content = $1,
		title = $2,
		categories = $3,
		updeted_at = CURRENT_TIMESTAMP
	WHERE
		id = $4
		and 
		userID = $5
	RETURNING
		id,
		userID, 
		content, 
		title,
		likes,
    	dislikes,
		views,
    	categories, 
		created_at,
		updeted_at`

	var respUser p.Post
	err := r.db.QueryRow(query, prr.Content, prr.Title, prr.Categories, prr.Id, prr.UserID).Scan(
		&respUser.Id,
		&respUser.UserID,
		&respUser.Content,
		&respUser.Title,
		&respUser.Likes,
		&respUser.Dislikes,
		&respUser.Views,
		&respUser.Categories,
		&respUser.CreatedAt,
		&respUser.UpdetedAt,
	)

	if err != nil {
		log.Println("Error updating user in postgres")
		return nil, err
	}
	return &respUser, nil
}

func (r *postRepo) DeletePost(pr *p.GetDeletePostRequest) (*p.Post, error) {

	var res p.Post

	query := `
	UPDATE
		posts
	SET
		deleted_at=CURRENT_TIMESTAMP
	WHERE
		id=$1
	and
		userID = $2
	RETURNING
	id,
	userID,
	content, 
    title,
    likes,
    dislikes,
	views,
    categories, 
    created_at,
	updeted_at,
	deleted_at`

	err := r.db.QueryRow(query, pr.Id, pr.UserID).Scan(
		&res.Id,
		&res.UserID,
		&res.Content,
		&res.Title,
		&res.Likes,
		&res.Dislikes,
		&res.Views,
		&res.Categories,
		&res.CreatedAt,
		&res.UpdetedAt,
		&res.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil

}
