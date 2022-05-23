package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leeerraaa/backend-app/internal/domain"
	"github.com/leeerraaa/backend-app/template"
)

func (h *Handler) downloadDocument(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdParse := fmt.Sprintf("%v", userId)

	id := c.Param("id")

	document, err := h.documentService.DocumentGet(userIdParse, id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	filePath, err := template.GenerateDocx(document)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	nameFile := "result.docx"

	file, err := ioutil.ReadFile(filePath + nameFile)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+nameFile+"")
	c.Writer.Header().Set("Content-Length", strconv.Itoa(len(file)))

	c.Writer.WriteHeader(http.StatusOK)

	c.Writer.Write(file)
}

func (h *Handler) documentList(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdParse := fmt.Sprintf("%v", userId)

	list, err := h.documentService.DocumentGetList(userIdParse)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getListInfo{
		Result: list,
	})
}

func (h *Handler) createDocument(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdParse := fmt.Sprintf("%v", userId)

	var document domain.DocumentInput
	if err := c.BindJSON(&document); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.documentService.DocumentCreate(document, userIdParse)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteDocument(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdParse := fmt.Sprintf("%v", userId)

	id := c.Param("id")

	err := h.documentService.DocumentDelete(userIdParse, id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
