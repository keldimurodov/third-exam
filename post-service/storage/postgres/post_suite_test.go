package postgres

// import (
// 	"log"
// 	"testing"
// 	"third-exam/post-service/config"
// 	"third-exam/post-service/pkg/db"
// 	"third-exam/post-service/storage/repo"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/suite"

// 	pbp "third-exam/post-service/genproto/post"
// )

// type PostRepositorySuiteTest struct {
// 	suite.Suite
// 	CleanUpFunc func()
// 	Repository  repo.PostStorageI
// }

// func (s *PostRepositorySuiteTest) SetupSuite() {
// 	pgPoll, err, cleanUp := db.ConnectToDB(config.Load())
// 	if err != nil {
// 		log.Fatal("Error while connecting database with suite test")
// 		return
// 	}
// 	s.CleanUpFunc = cleanUp
// 	s.Repository = NewPostRepo(pgPoll)
// }

// // test func
// func (s *PostRepositorySuiteTest) TestUserCRUD() {
// 	// create post
// 	postRequest := &pbp.Post{
// 		Id:         uuid.NewString(),
// 		UserID:     uuid.NewString(),
// 		Title:      "test Title",
// 		Categories: "First Category",
// 		Content:    "New Content",
// 	}
// 	respPost, err := s.Repository.CreatePost(postRequest)
// 	s.Suite.NotNil(respPost)
// 	s.Suite.NoError(err)
// 	s.Suite.Equal(respPost, postRequest)

// 	// get post
// 	post, err := s.Repository.GetPost(&pbp.GetPostRequest{Id: respPost.Id})
// 	s.Suite.NoError(err)
// 	s.Suite.NotNil(post)
// 	s.Suite.Equal(post.Post.Id, respPost.Id)

// }

// func TestExampleTestSuite(t *testing.T) {
// 	suite.Run(t, new(PostRepositorySuiteTest))
// }
