package v1

import (
	"context"
	// "fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"

	models "third-exam/api-gateway/api/handlers/models"
	// "third-exam/api-gateway/api/handlers/tokens"
	p "third-exam/api-gateway/genproto/post"
	l "third-exam/api-gateway/pkg/logger"
	"third-exam/api-gateway/pkg/utils"
)

// CreatePost ...
// @Summary Create Post ...
// @Security ApiKeyAuth
// @Description Api for creating a new post
// @Tags post
// @Accept json
// @Produce json
// @Param Post body models.PostRequest true "createPost"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/post/ [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		body        models.Post
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	// h.jwthandler = tokens.JWTHandler{
	// 	Sub:       body.,
	// 	Role:      "admin",
	// 	SigninKey: "admin",
	// 	Log:       h.log,
	// }

	// access, refresh, err := h.jwthandler.GenerateAuthJWT()
	// if err != nil {
	// 	c.JSON(http.StatusConflict, gin.H{
	// 		"error": "error while generating jwt",
	// 	})
	// 	h.log.Error("error generate new jwt tokens", l.Error(err))
	// 	return
	// }
	// fmt.Println("Tokens are working well")

	postID := uuid.NewString()

	response, err := h.serviceManager.PostService().CreatePost(ctx, &p.Post{
		Id:       postID,
		UserID:   body.UserID,
		Content:  body.Content,
		Title:    body.Title,
		Likes:    body.Likes,
		Dislikes: body.Dislikes,
		Views:    body.Views,
	})

	// respBody := &models.Post{
	// 	Id:       response.Id,
	// 	UserID:   response.UserID,
	// 	Content:  response.Content,
	// 	Title:    response.Title,
	// 	Likes:    response.Likes,
	// 	Dislikes: response.Dislikes,
	// 	Views:    response.Views,
	// }

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to Create Post", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetPost get post by id
// @Summary GetPost
// @Security ApiKeyAuth
// @Description Api for getting post by id
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/post/{id} [get]
func (h *handlerV1) GetPost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetPost(
		ctx, &p.GetPostRequest{
			Id: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to GET Post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAllPosts returns list of posts from the service
// @Summary All posts
// @Security ApiKeyAuth
// @Description Api returns list of posts
// @Tags post
// @Accept json
// @Produce json
// @Param page path int64 true "Page"
// @Param limit path int64 true "Limit"
// @Succes 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/ [get]
func (h *handlerV1) GetAllPosts(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("Failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetAllPosts(
		ctx, &p.GetAllRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to list post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdatePost updates post by id
// @Summary UpdatePost
// @Security ApiKeyAuth
// @Description Api returns updates post
// @Tags post
// @Accept json
// @Produce json
// @Param Product body models.Post true "UpdatePost"
// @Succes 200 {Object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/post/{id} [put]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		body        models.Post
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().UpdatePost(ctx, &p.Post{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to update post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeletePost deleted Post by id
// @Summary Delete Post
// @Security ApiKeyAuth
// @Description Api deletes post
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Succes 200 {Object} model.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/post/{id} [delete]
func (h *handlerV1) DeletePost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().DeletePost(
		ctx, &p.GetDeletePostRequest{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to delete post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
