package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/leeerraaa/backend-app/internal/domain"
	"github.com/leeerraaa/backend-app/internal/service"
)

type User interface {
	GenerateToken(login, password string) (string, error)
	ParseToken(accessToken string) (string, error)
	UserInfo(userId string) (domain.User, error)
}

type Document interface {
	DocumentGetList(userId string) ([]domain.Document, error)
	DocumentGet(userId string, documentId string) (domain.Document, error)
	DocumentCreate(data domain.DocumentInput, userId string) (string, error)
	DocumentDelete(userId string, documentId string) error
}

type Handler struct {
	userService     User
	documentService Document
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		userService:     service.User,
		documentService: service.Document,
	}
}

type getListInfo struct {
	Result []domain.Document `json:"result"`
}

type getUserInfo struct {
	Result domain.User `json:"result"`
}

type getCreationToken struct {
	AccessToken string `json:"token"`
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.Use(h.loggingMiddleware)

		auth.POST("/sign-in", h.SignIn)
	}

	user := router.Group("/user")
	{
		user.Use(h.userIdentify)
		user.Use(h.loggingMiddleware)

		user.GET("/", h.UserInfo)
	}

	document := router.Group("/document")
	{
		document.Use(h.userIdentify)
		document.Use(h.loggingMiddleware)

		document.GET("/", h.documentList)
		document.GET("/:id", h.downloadDocument)
		document.POST("/", h.createDocument)
		document.DELETE("/:id", h.deleteDocument)
	}

	return router
}
