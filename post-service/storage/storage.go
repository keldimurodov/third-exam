package storage

import (
	"third-exam/post-service/storage/mongodb"
	"third-exam/post-service/storage/postgres"
	"third-exam/post-service/storage/repo"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

// IStorage ...
type IStorage interface {
	Post() repo.PostStorageI
}

type storagePg struct {
	db       *sqlx.DB
	postRepo repo.PostStorageI
}

type storageMongo struct {
	db       *mongo.Collection
	postRepo repo.PostStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		postRepo: postgres.NewPostRepo(db),
	}
}

func (s storagePg) Post() repo.PostStorageI {
	return s.postRepo
}

func (s storageMongo) Post() repo.PostStorageI {
	return s.postRepo
}

func NewStorageMongo(db *mongo.Collection) *storageMongo {
	return &storageMongo{
		db:       db,
		postRepo: mongodb.NewCommentRepoMongo(db),
	}
}
