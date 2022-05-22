package psql

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/leeerraaa/backend-app/internal/domain"
)

type DocumentRepo struct {
	db *sql.DB
}

func NewDocumentRepo(db *sql.DB) *DocumentRepo {
	return &DocumentRepo{db: db}
}

func (d *DocumentRepo) DocumentGet(userId string, documentId string) (domain.Document, error) {
	rows, err := d.db.Query("SELECT * FROM documents WHERE user_id = $1 AND id = $2", userId, documentId)
	if err != nil {
		return domain.Document{}, err
	}

	var document domain.Document
	for rows.Next() {
		if err := rows.Scan(&document.Id, &document.UserId, &document.Specialty, &document.EducationalLevel, &document.EducationalProgram, &document.Subject, &document.Lectures, &document.PracticalClasses, &document.LaboratoryClasses, &document.DateOfCreation); err != nil {
			return domain.Document{}, err
		}
	}

	return document, nil
}

func (d *DocumentRepo) DocumentGetList(userId string) ([]domain.Document, error) {
	rows, err := d.db.Query("SELECT * FROM documents WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}

	documents := make([]domain.Document, 0)
	for rows.Next() {
		var document domain.Document
		if err := rows.Scan(&document.Id, &document.UserId, &document.Specialty, &document.EducationalLevel, &document.EducationalProgram, &document.Subject, &document.Lectures, &document.PracticalClasses, &document.LaboratoryClasses, &document.DateOfCreation); err != nil {
			return nil, err
		}

		documents = append(documents, document)
	}

	return documents, rows.Err()
}

func (d *DocumentRepo) DocumentCreate(data domain.DocumentInput, userId string) (string, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return "", err
	}

	var id string
	row, err := tx.Prepare("INSERT INTO documents(id, user_id, specialty, educational_level, educational_program, subject, lectures, practical_classes, laboratory_classes, date_of_creation) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id")
	if err != nil {
		return "", err
	}

	defer row.Close()

	if err = row.QueryRow(uuid.NewString(), userId, data.Specialty, data.EducationalLevel, data.EducationalProgram, data.Subject, data.Lectures, data.PracticalClasses, data.LaboratoryClasses, time.Now()).Scan(&id); err != nil {
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return id, nil
}

func (d *DocumentRepo) DocumentDelete(userId string, documentId string) error {
	_, err := d.db.Exec("DELETE FROM documents WHERE id = $1 AND user_id = $2", documentId, userId)

	return err
}
