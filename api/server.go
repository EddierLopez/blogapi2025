package api

import (
	"restapi/dto"
	"restapi/security"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

type Server struct {
	dbtx         *dto.DbTransaction
	tokenBuilder security.Builder
	router       *gin.Engine
}

func NewServer(dbtx *dto.DbTransaction) (*Server, error) {
	tokenBuilder, err := security.NewPasetoBuilder("12345678123456781234567812345678")
	if err != nil {
		return nil, err
	}
	server := &Server{
		dbtx:         dbtx,
		tokenBuilder: tokenBuilder,
	}
	router := gin.Default()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
	//RUTAS {ENDPOINTS} DEL API
	router.POST("api/v1/login", server.login)
	router.POST("api/v1/user", server.createUser)
	router.GET("api/v1/category/:id", server.getCategory)
	router.GET("api/v1/category", server.getCategories)
	router.GET("api/v1/post", server.getAllPost)
	router.GET("api/v1/post/:category_id", server.getPostByCategory)

	//RUTAS CON MIDDLEWARE
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenBuilder))
	authRoutes.POST("api/v1/category", server.createCategory)
	authRoutes.POST("api/v1/post", server.createPost)
	authRoutes.PUT("api/v1/user/:id", server.updateUser)
	authRoutes.PUT("api/v1/user", server.updateUser)
	authRoutes.PUT("api/v1/userpass", server.updatePassword)
	authRoutes.PUT("api/v1/userrole", server.updateRole)
	///FIN RUTAS///
	server.router = router
	return server, nil
}

func (server *Server) Start(url string) error {
	return server.router.Run(url)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
