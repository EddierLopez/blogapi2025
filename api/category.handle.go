package api

import (
	"database/sql"
	"net/http"
	"restapi/dto"

	"github.com/gin-gonic/gin"
)

type createCategoryRequest struct {
	Name string `json:"name" binding:required`
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.CreateCategoryParams{
		Name: req.Name,
	}
	category, err := server.dbtx.CreateCategory(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var lastId, _ = category.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

// Obtener una categoria por ID
type getCategoryRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getCategory(ctx *gin.Context) {
	var req getCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	category, err := server.dbtx.GetCategoryById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, category)
}
func (server *Server) getCategories(ctx *gin.Context) {
	categories, err := server.dbtx.GetAllCategories(ctx)
	//fmt.Println(categories)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, categories)
}
