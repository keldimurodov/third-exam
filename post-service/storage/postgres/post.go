package postgres

import (
	"log"
	p "third-exam/post-service/genproto/post"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type postRepo struct {
	db *sqlx.DB
}

// NewPostRepo ...
func NewPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}

func (r *postRepo) CreatePost(post *p.Post) (*p.Post, error) {

	idd := uuid.NewString()
	post.Id = idd

	var re p.Post

	query := `INSERT INTO Posts(
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

func (r *postRepo) GetPost(pp *p.GetPostRequest) (*p.GetP, error) {

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
        return nil, errr
    }

	que := `SELECT 
		id, 
		owner_id, 
		post_id, 
		text, 
        created_at, 
        updeted_at 
	FROM 
		Comments 
	WHERE 
		post_id = $1`
	rows, err := r.db.Query(que, pp.Id)
	if err != nil {
		// agar comment yo'q bo'lsa shunchaki nil qaytarvoradi faqat comment uchun
		return &p.GetP{
			Post:     &respUser,
			Comments: nil,
		}, nil
	}

	var comments []*p.Comments

	for rows.Next() {
		var comment p.Comments

		err := rows.Scan(&comment.Id, &comment.OwnerId, &comment.PostId, &comment.Text)
		if err != nil {
			// birorta commentni olishda muammo bo'lsa shunchaki bo'sh string qaytarvorishi uchun
			log.Println("Commentlarni olishda muammo bo'ldi...")
		}

		comments = append(comments, &comment)
	}

	return &p.GetP{
		Post:     &respUser,
		Comments: comments,
	}, nil

	// return &respUser, nil
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
