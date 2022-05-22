package psql

import (
	"database/sql"
	"github.com/leeerraaa/backend-app/internal/domain"
)

type User interface {
	GetUser(login, password string) (string, error)
	UserInfo(userId string) (domain.User, error)
}

type Document interface {
	DocumentGetList(userId string) ([]domain.Document, error)
	DocumentGet(userId string, documentId string) (domain.Document, error)
	DocumentCreate(data domain.DocumentInput, userId string) (string, error)
	DocumentDelete(userId string, documentId string) error
}

type Repository struct {
	User
	Document
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:     NewAuthRepository(db),
		Document: NewDocumentRepo(db),
	}
}
