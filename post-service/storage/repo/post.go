package repo

import (
	pb "third-exam/post-service/genproto/post"
)

// PostStorageI ...
type PostStorageI interface {
	CreatePost(pb *pb.Post) (*pb.Post, error)
	GetPost(req *pb.GetPostRequest) (*pb.GetP, error)
	GetAllPosts(*pb.GetAllRequest) (*pb.GetAllResponse, error)
	UpdatePost(*pb.Post) (*pb.Post, error)
	DeletePost(*pb.GetDeletePostRequest) (*pb.Post, error)
}
