package api

import (
	"database/sql"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"restapi/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createPostRequest struct {
	UserID     int32  `json:"user_id"`
	CategoryID int32  `json:"category_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Image      string `json:"image"`
}

func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.CreatePostParams{
		UserID:     req.UserID,
		CategoryID: req.CategoryID,
		Title:      req.Title,
		Content:    req.Content,
		Image:      req.Image,
	}
	result, err := server.dbtx.CreatePost(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var lastId, _ = result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

func (server *Server) getAllPost(ctx *gin.Context) {
	posts, err := server.dbtx.GetAllPost(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, posts)

}

type getPostByUserRequest struct {
	UserID int32 `uri:"user_id" binding:"required,min=1"`
}

func (server *Server) getPostByUser(ctx *gin.Context) {
	var req getPostByUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	posts, err := server.dbtx.GetPostsByUser(ctx, req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, posts)

}

type getPostByCategoryRequest struct {
	CategoryID int32 `uri:"category_id" binding:"required,min=1"`
}

func (server *Server) getPostByCategory(ctx *gin.Context) {
	var req getPostByCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	posts, err := server.dbtx.GetPostsByCategory(ctx, req.CategoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, posts)
}

type deletePostRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deletePost(ctx *gin.Context) {
	var req deletePostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	result, err := server.dbtx.DeletePost(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var rows, _ = result.RowsAffected()
	ctx.JSON(http.StatusOK, gin.H{"rows_affected": rows})

}
func (server *Server) uploadPostImg(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("file0")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	defer file.Close()
	uploadDir := "utils/images/posts"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	filename := uuid.New().String() + "_" + filepath.Base(fileHeader.Filename)
	destinationFile, err := os.Create(filepath.Join(uploadDir, filename))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer destinationFile.Close()
	_, err = io.Copy(destinationFile, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"file_name": filename,
		"message":   "Archivo cargado exitosamente",
	})
}

type getPostImageRequest struct {
	Name string `uri:"name",binding:"required"`
}

func (server *Server) getPostImg(ctx *gin.Context) {
	var req getPostImageRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fileUrl := "utils/images/posts/" + req.Name
	ctx.File(fileUrl)
}
