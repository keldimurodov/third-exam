package postgres

import (
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
			text) 
		VALUES ($1, $2, $3, $4) 
		RETURNING 
			id, 
			owner_id,
			post_id, 
			text, 
			created_at, 
        	updeted_at`

	err := r.db.QueryRow(query, comment.Id, comment.OwnerId, comment.PostId, comment.Text).Scan(
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
        updeted_at
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
		&res.UpdatedAt)
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
        updeted_at
	FROM 
		Comments
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

func (r *commentRepo) UpdateComment(cc *c.UpdateRequest) (*c.Comment, error) {
	var res c.Comment
	query := `
	UPDATE
		Comments
	SET
		text = $1,
		updeted_at = CURRENT_TIMESTAMP
	WHERE
		id = $2
	and
		owner_id = $3
	and
		post_id = $4
	RETURNING
		id, 
		owner_id,
		post_id, 
		text, 
		created_at, 
		updeted_at`

	err := r.db.QueryRow(query, cc.Text, cc.Id, cc.OwnerId, cc.PostId).Scan(
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
		updeted_at,
		deleted_at`

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
