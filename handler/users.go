package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spargapees/users-api/dto"
	"net/http"
	"strconv"
)

func (h *Handler) createUser(c *gin.Context) {
	var input dto.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	response := map[string]interface{}{
		"all_users": users,
	}
	c.JSON(http.StatusOK, response)
}

func (h *Handler) updateUser(c *gin.Context) {
	userId, err := ValidateId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var input dto.UserUpdate

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateUser(userId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type statusResponse struct {
	Status string `json:"status"`
}

func ValidateId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return -1, err
	}

	if id <= 0 {
		return -1, errors.New("id cannot be negative")
	}
	return id, nil
}
