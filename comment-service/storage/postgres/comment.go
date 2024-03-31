package postgres

import (
	"log"
	c "third-exam/comment-service/genproto/comment"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type commentRepo struct {
	db *sqlx.DB
}

// NewPostRepo ...
func NewCommentRepo(db *sqlx.DB) *commentRepo {
	return &commentRepo{db: db}
}

func (r *commentRepo) CreateComment(comment *c.Comment) (*c.Comment, error) {

	idd := uuid.NewString()
	comment.Id = idd

	var res c.Comment

	query := `INSERT INTO Comments(
			id, 
			owner_id,
			post_id, 
			text, 
			created_at, 
        	updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING 
			id, 
			owner_id,
			post_id, 
			text, 
			created_at, 
        	updated_at`
	err := r.db.QueryRow(query, comment.Id, comment.OwnerId, comment.PostId, comment.Text, comment.CreatedAt, comment.UpdatedAt).Scan(
		&res.Id,
		&res.OwnerId,
		&res.PostId,
		&res.Text,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *commentRepo) GetComment(cc *c.GetCommentRequest) (*c.Comment, error) {

	var res c.Comment
	query := `
	SELECT
		id, 
		owner_id,
		post_id, 
		text, 
		created_at, 
        updated_at
	FROM 
		Comments
	WHERE
	    id = $1
		and
		deleted_at IS NULL`
	err := r.db.QueryRow(query, cc.Id).Scan(
		&res.Id,
		&res.OwnerId,
		&res.PostId,
		&res.Text,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *commentRepo) GetAllComments(cc *c.GetAllCommentRequest) (*c.GetAllCommentResponse, error) {
	var allComment c.GetAllCommentResponse
	query := `
	SELECT
		id, 
		owner_id,
		post_id, 
		text, 
		created_at, 
        updated_at
	FROM 
		post
	WHERE 
		deleted_at IS NULL
	LIMIT $1
	OFFSET $2`

	offset := cc.Limit * (cc.Page - 1)

	rows, err := r.db.Query(query, cc.Limit, offset)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var pr c.Comment

		err := rows.Scan(
			&pr.Id,
			&pr.OwnerId,
			&pr.PostId,
			&pr.Text,
			&pr.CreatedAt,
			&pr.UpdatedAt)

		if err != nil {
			return nil, err
		}

		allComment.Comments = append(allComment.Comments, &pr)
	}
	return &allComment, nil
}

func (r *commentRepo) UpdateComment(cc *c.Comment) (*c.Comment, error) {
	query := `
	UPDATE
		Comments
	SET
		text,
		updeted_at = CURRENT_TIMESTAMP
	WHERE
		id = $1
		and 
		post_id = $2
	RETURNING
		id, 
		owner_id,
		post_id, 
		text, 
		created_at,
		updeted_at`

	var respComment c.Comment
	err := r.db.QueryRow(query, cc.Text, cc.Id, cc.PostId).Scan(
		&respComment.Id,
		&respComment.OwnerId,
		&respComment.PostId,
		&respComment.Text,
		&respComment.CreatedAt,
		&respComment.UpdatedAt,
	)

	if err != nil {
		log.Println("Error updating user in postgres")
		return nil, err
	}
	return &respComment, nil
}

func (r *commentRepo) DeleteComment(pr *c.GetDeleteCommentRequest) (*c.Comment, error) {

	var res c.Comment

	query := `
	UPDATE
		Comments
	SET
		deleted_at=CURRENT_TIMESTAMP
	WHERE
		id=$1
	and
	    owner_id=$2
	and 
	    post_id=$3

	RETURNING
		id, 
		owner_id,
		post_id, 
		text, 
		created_at,
		updeted_at`

	err := r.db.QueryRow(query, pr.Id, pr.OwnerId, pr.PostId).Scan(
		&res.Id,
		&res.OwnerId,
		&res.PostId,
		&res.Text,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil

}
