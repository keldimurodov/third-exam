package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	p "third-exam/post-service/genproto/post"
)

type postRepo struct {
	mdb *mongo.Collection
}

func (u *postRepo) CreatePost(post *p.Post) (*p.Post, error) {
	update := bson.M{
		"id":         post.Id,
		"userid":     post.UserID,
		"content":    post.Content,
		"title":      post.Title,
		"likes":      post.Likes,
		"dislikes":   post.Dislikes,
		"views":      post.Views,
		"category":   post.Categories,
		"created_at": time.Now().String(),
		"updatedat":  time.Now().String(),
	}
	_, err := u.mdb.InsertOne(context.Background(), update)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"id": post.Id}
	var postResp p.Post
	err = u.mdb.FindOne(context.Background(), filter).Decode(&postResp)
	if err != nil {
		return nil, err
	}

	return &postResp, nil
}

func (u *postRepo) GetPost(req *p.GetPostRequest) (*p.GetP, error) {
	filter := bson.M{"id": req.Id}
	var post p.GetP
	err := u.mdb.FindOne(context.Background(), filter).Decode(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (u *postRepo) GetAllPosts(req *p.GetAllRequest) (*p.GetAllResponse, error) {
	options := options.Find()
	options.SetLimit(req.Limit)
	options.SetSkip(req.Limit * (req.Page - 1))

	cursor, err := u.mdb.Find(context.Background(), bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var allPosts p.GetAllResponse
	for cursor.Next(context.Background()) {
		var post p.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		allPosts.Posts = append(allPosts.Posts, &post)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &allPosts, nil
}

func (u *postRepo) UpdatePost(req *p.Post) (*p.Post, error) {
	update := bson.M{
		"$set": bson.M{
			"id":        req.Id,
			"userid":    req.UserID,
			"content":   req.Content,
			"title":     req.Title,
			"likes":     req.Likes,
			"dislikes":  req.Dislikes,
			"views":     req.Views,
			"category":  req.Categories,
			"updatedat": time.Now().String(),
		},
	}
	filter := bson.M{"id": req.Id}
	updateResult, err := u.mdb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	if updateResult.ModifiedCount == 0 {
		return nil, fmt.Errorf("post with id %v not found", req.Id)
	}
	filter = bson.M{"id": req.Id}
	var post p.Post
	err = u.mdb.FindOne(context.Background(), filter).Decode(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (u *postRepo) DeletePost(req *p.GetDeletePostRequest) (*p.Post, error) {
	filter := bson.M{"id": req.Id}
	var post p.Post
	err := u.mdb.FindOne(context.Background(), filter).Decode(&post)
	if err != nil {
		return nil, err
	}

	result, err := u.mdb.DeleteOne(context.Background(), filter)
	if err != nil {
		return &p.Post{}, err
	}

	if result.DeletedCount == 0 {
		return &p.Post{}, fmt.Errorf("post with id %v not found", req.Id)
	}

	return &post, nil
}

func NewCommentRepoMongo(db *mongo.Collection) *postRepo {
	return &postRepo{mdb: db}
}
