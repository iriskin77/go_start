package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gostart "github.com/iriskin77/go_start"
)

func (h *Handler) signUp(c *gin.Context) {
	// Данные о пользователей
	var input gostart.User

	// Здесь парсим ответ от пользователя
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return

	}

	// Когда мы распарсили тело json ответа, его нужно передать на уровень ниже
	// в services, где будет реализована бизнес-логика

}

func (h *Handler) signIn(c *gin.Context) {

}
