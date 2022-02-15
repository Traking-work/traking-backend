package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Traking-work/traking-backend.git/internal/domain"
)

func (h *Handler) GetTeamLeads(c *gin.Context) {
	teamleads, teamleadCreate, err := h.services.Admin.GetTeamLeads(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"teamleads": teamleads,
		"teamleadCreate": teamleadCreate,
	})
}

func (h *Handler) AddUser(c *gin.Context) {
	var inp domain.UserData
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	err := h.services.Admin.AddUser(c, inp)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}