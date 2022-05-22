package service

import (
	"github.com/leeerraaa/backend-app/internal/domain"
	repo "github.com/leeerraaa/backend-app/internal/repository/psql"
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

type Service struct {
	User
	Document
}

func NewService(repo *repo.Repository) *Service {
	return &Service{
		User: NewAuthService(repo.User),
		Document: NewDocumentService(repo.Document),
	}
}
