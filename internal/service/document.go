package service

import (
	"github.com/leeerraaa/backend-app/internal/domain"
	repo "github.com/leeerraaa/backend-app/internal/repository/psql"
)

type DocumentService struct {
	repo repo.Document
}

func NewDocumentService(repo repo.Document) *DocumentService {
	return &DocumentService{
		repo: repo,
	}
}

func (d *DocumentService) DocumentGetList(userId string) ([]domain.Document, error) {
	return d.repo.DocumentGetList(userId)
}

func (d *DocumentService) DocumentGet(userId string, documentId string) (domain.Document, error) {
	return d.repo.DocumentGet(userId, documentId)
}

func (d *DocumentService) DocumentCreate(data domain.DocumentInput, userId string) (string, error) {
	return d.repo.DocumentCreate(data, userId)
}

func (d *DocumentService) DocumentDelete(userId string, documentId string) error {
	return d.repo.DocumentDelete(userId, documentId)
}
