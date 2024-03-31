package postgres

import (
	"fmt"
	"log"
	u "third-exam/user-service/genproto/user"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *u.User) (*u.User, error) {

	var res u.User
	query := `
		INSERT INTO Users(
			first_name, 
			last_name,
			bio,
			website, 
			email,
			password
		)
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING 
			id, 
			first_name, 
			last_name,
			bio,
            website, 
			email,
			password,
			created_at,
			updeted_at
	`
	err := r.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Bio,
        user.Website,
        user.Email,
        user.Password).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Bio,
        &res.Website,
		&res.Email,
		&res.Password,
		&res.CreatedAt,
		&res.UpdetedAt)
	if err != nil {
		fmt.Println("Error Creating user")
		return nil, err
	}

	return &res, nil
}

func (r *userRepo) GetUser(id *u.GetUserRequest) (*u.User, error) {
	var user u.User
	query := `
	SELECT
		id,
		first_name,
		last_name,
		bio,
		website,
		email,
		password,
		created_at,
		updeted_at
	FROM 
		Users
	WHERE
		id=$1 
	AND deleted_at IS NULL
	`
	err := r.db.QueryRow(query, id.UserId).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Bio,
        &user.Website,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdetedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) UpdateUser(user *u.User) (*u.User, error) {
	var res u.User
	query := `
	UPDATE
		Users
	SET
		first_name=$1,
		last_name=$2,
		bio=$3,
        website=$4,
        email=$5,
        password=$6
		updeted_at=CURRENT_TIMESTAMP
	WHERE
		id=$7
	returning
		id, 
		first_name,
		last_name,
		bio,
        website, 
        email,
        password,
		created_at,
		updeted_at
	`
	err := r.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Bio,
        user.Website,
        user.Email,
		user.Password,
		user.Id).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Bio,
        &res.Website,
        &res.Email,
        &res.Password,
		&res.CreatedAt,
		&res.UpdetedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *userRepo) DeleteUser(user *u.GetUserRequest) (*u.User, error) {
	var res u.User
	query := `
	UPDATE
		Users
	SET
		deleted_at=CURRENT_TIMESTAMP
	WHERE
		id=$1
	RETURNING
		id, 
		first_name, 
		last_name,
		bio,
        website, 
        email,
        password,
		created_at,
		updeted_at,
		deleted_at
	`
	err := r.db.QueryRow(query, user.UserId).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Bio,
        &res.Website,
        &res.Email,
        &res.Password,
		&res.CreatedAt,
		&res.UpdetedAt,
		&res.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil

}

func (r *userRepo) GetAllUsers(user *u.GetAllRequest) (*u.GetAllResponse, error) {
	var allUser u.GetAllResponse
	query := `
	SELECT
		id,
		first_name,
		last_name,
		bio,
		website,
        email,
        password,
		created_at,
		updeted_at
	FROM 
		Users 
	WHERE 
		deleted_at IS NULL
	LIMIT $1
	OFFSET $2`
	offset := user.Limit * (user.Page - 1)
	rows, err := r.db.Query(query, user.Limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user u.User
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Bio,
            &user.Website,
            &user.Email,
            &user.Password,
			&user.CreatedAt,
			&user.UpdetedAt)
		if err != nil {
			return nil, err
		}
		allUser.Users = append(allUser.Users, &user)
	}
	return &allUser, nil
}

func (s *userRepo) CheckUniqueness(req *u.CheckUniquenessRequest) (*u.CheckUniquenessResponse, error) {
	var email int

	fmt.Println(req.Field, req.Value)

	query := fmt.Sprintf("SELECT count(1) from users WHERE %s = $1 ", req.Field)
	err := s.db.QueryRow(query, req.Value).Scan(&email)
	if err != nil {
		log.Fatal("error while checking!!!", err.Error())
	}
	if email == 1 {

		return &u.CheckUniquenessResponse{
			Result: true,
		}, nil
	}

	return &u.CheckUniquenessResponse{
		Result: false,
	}, nil
}

func SubVerification(db *sqlx.DB, user *u.UserDetail) (*u.User, error) {

	var res u.User

	query := `
		INSERT INTO Users(
			first_name, 
			last_name,
			bio,
			website,
			email,
			password
		)
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING 
			id, 
			first_name, 
			last_name,
			bio,
			website,
			email,
			password,
			created_at,
			updeted_at
	`
	err := db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Bio,
		user.Website,
		user.Email,
		user.Password).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Bio,
        &res.Website,
		&res.Email,
		&res.Password,
		&res.CreatedAt,
		&res.UpdetedAt)
	if err != nil {
		fmt.Println("Error subvalidation")
		return nil, err
	}

	return &res, nil
}

func (r *userRepo) Login(id *u.LoginRequest) (*u.User, error) {
	var user u.User
	query := `
	SELECT
		id,
		first_name,
		last_name,
		bio,
		website,
		email,
		password,
		created_at,
		updeted_at
	FROM 
		Users
	WHERE
		email=$1
	AND 
		password = $2
	AND deleted_at IS NULL
	`
	err := r.db.QueryRow(query, id.Email, id.Password).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Bio,
        &user.Website,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdetedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
