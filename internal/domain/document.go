package domain

import "time"

type Document struct {
	Id                 string    `json:"id"`
	UserId             string    `json:"user_id"`
	Specialty          string    `json:"specialty"`
	EducationalLevel   string    `json:"educational_level"`
	EducationalProgram string    `json:"educational_program"`
	Subject            string    `json:"subject"`
	Lectures           int32     `json:"lectures"`
	PracticalClasses   int32     `json:"practical_classes"`
	LaboratoryClasses  int32     `json:"laboratory_classes"`
	DateOfCreation     time.Time `json:"date_of_creation"`
}

type DocumentInput struct {
	Specialty          *string `json:"specialty"`
	EducationalLevel   *string `json:"educational_level"`
	EducationalProgram *string `json:"educational_program"`
	Subject            *string `json:"subject"`
	Lectures           *int32  `json:"lectures"`
	PracticalClasses   *int32  `json:"practical_classes"`
	LaboratoryClasses  *int32  `json:"laboratory_classes"`
}
