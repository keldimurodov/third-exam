package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"

	models "third-exam/api-gateway/api/handlers/models"
	cc "third-exam/api-gateway/genproto/comment"
	l "third-exam/api-gateway/pkg/logger"
	"third-exam/api-gateway/pkg/utils"
)

// CreateComment ...
// @Summary CreateComment ...
// @Security ApiKeyAuth
// @Description Api for creating a new comment
// @Tags comment
// @Accept json
// @Produce json
// @Param Post body models.PostCommentsRequest true "createComment"
// @Success 200 {object} models.PostComments
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comment/ [post]
func (h *handlerV1) CreateComment(c *gin.Context) {
	var (
		body        models.Comments
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

	commentID := uuid.NewString()
	commentOwnerID := uuid.NewString()
	commentPostID := uuid.NewString()

	response, err := h.serviceManager.CommentService().CreateComment(ctx, &cc.Comment{
		Id:        commentID,
		OwnerId:   commentOwnerID,
		PostId:    commentPostID,
		Text:      body.Text,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		DeletedAt: body.DeletedAt,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to Create Comment", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetComment get comment by id
// @Summary GetComment
// @Security ApiKeyAuth
// @Description Api for getting comment by id
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.PostComments
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comment/{id} [get]
func (h *handlerV1) GetComment(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().GetComment(
		ctx, &cc.GetCommentRequest{
			Id: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to GET Comment", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAllComments returns list of comments from the service
// @Summary All comments from the service
// @Security ApiKeyAuth
// @Description Api returns list of comments from
// @Tags comment
// @Accept json
// @Produce json
// @Param page path int64 true "Page"
// @Param limit path int64 true "Limit"
// @Succes 200 {object} models.PostComments
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comments/ [get]
func (h *handlerV1) GetAllComments(c *gin.Context) {
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

	response, err := h.serviceManager.CommentService().GetAllComment(
		ctx, &cc.GetAllCommentRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to list Commments", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateComment updates Comment by id
// @Summary Update Commment
// @Security ApiKeyAuth
// @Description Api returns updates comment
// @Tags comment
// @Accept json
// @Produce json
// @Param Post body models.PostComments true "UpdateComment"
// @Succes 200 {Object} models.PostComments
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comment/{id} [put]
func (h *handlerV1) UpdateComment(c *gin.Context) {
	var (
		body        models.Comments
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

	response, err := h.serviceManager.CommentService().UpdateComment(ctx, &cc.UpdateRequest{
		Id:      body.Id,
		OwnerId: body.OwnerId,
		PostId:  body.PostId,
		Text:    body.Text,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to update Comment", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteComment deleted Comment by id
// @Summary Delete Comment by id
// @Security ApiKeyAuth
// @Description Api deletes post
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Succes 200 {Object} model.Comments
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comment/{id} [delete]
func (h *handlerV1) DeleteComment(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().DeleteComment(
		ctx, &cc.GetDeleteCommentRequest{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to delete Comment", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
