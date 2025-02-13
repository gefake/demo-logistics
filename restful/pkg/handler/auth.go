package handler

import (
	"github.com/gin-gonic/gin"
	"logistic_api/pkg/database"
	"net/http"
)

type Auth struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary Логин в системе
// @Tags         Auth
// @Description Позволяет залогиниться пользователю в системе
// @Accept json
// @Produce json
// @Param input body Auth true "Данные для входа"
// @Success 200 {object} database.User
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var input Auth

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	hasLogin, err := h.Services.Login(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if !hasLogin {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	usr, err := h.Services.UserRepository.GetByLogin(input.Login)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, usr)
}

type authRegResp struct {
	ID uint `json:"id"`
}

// Register godoc
// @Summary Регистрация нового пользователя
// @Tags         Auth
// @Description Создает нового пользователя в системе
// @Accept json
// @Produce json
// @Param input body database.User true "Детали о пользователе"
// @Success 200 {object} authRegResp
// @Router /auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	var input database.User

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userRole, err := h.Services.RoleRepository.GetRoleByName(database.UserRoleName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	input.RoleID = userRole.ID

	defPos, err := h.Services.PositionRepository.GetPositionByName(database.Positions[0])
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	input.PositionID = defPos.ID

	id, err := h.Services.UserRepository.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
