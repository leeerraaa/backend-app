package template

import (
	"os"

	"github.com/leeerraaa/backend-app/internal/domain"
	docx "github.com/lukasjarosch/go-docx"
	"github.com/sirupsen/logrus"
)

func GenerateDocx(document domain.Document) (string, error) {
	mydir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	pathToTemplate := mydir + "/template/template.docx"
	pathToNew := mydir + "/template/result.docx"

	if document.EducationalProgram == "'Системний аналіз' 2016" && document.Subject == "Комп'ютерна обробка зображень та мультимедіа" {
		pathToTemplate = mydir + "/template/template-analiz-media.docx"
	}

	if document.EducationalProgram == "'Системний аналіз' 2016" && document.Subject == "Інтернет-технології та проектування WEB-додатків" {
		pathToTemplate = mydir + "/template/template-analiz-web.docx"
	}

	replaceMap := docx.PlaceholderMap{
		"_specialty_":           document.Specialty,
		"_educational_level_":   document.EducationalLevel,
		"_educational_program_": document.EducationalProgram,
		"_subject_":             document.Subject,
		"_lectures_":            document.Lectures,
		"_practical_classes_":   document.PracticalClasses,
		"_laboratory_classes_":  document.LaboratoryClasses,
	}

	doc, err := docx.Open(pathToTemplate)
	if err != nil {
		return "", err
	}
	defer doc.Close()

	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		return "", err
	}

	err = doc.WriteToFile(pathToNew)
	if err != nil {
		return "", err
	}

	logrus.Println("Success")

	returnedPath := mydir + "/template/"

	return returnedPath, nil
}
