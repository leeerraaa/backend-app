package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leeerraaa/backend-app/internal/domain"
)

func (h *Handler) SignIn(c *gin.Context) {
	var input domain.SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.userService.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getCreationToken{
		AccessToken: token,
	})
}

func (h *Handler) UserInfo(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdParse := fmt.Sprintf("%v", userId)

	userData, err := h.userService.UserInfo(userIdParse)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getUserInfo{
		Result: userData,
	})
}
