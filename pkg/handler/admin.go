package handler

import (
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/Traking-work/traking-backend.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (h *Handler) GetWorkers(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	workers, err := h.services.Admin.GetWorkers(c, userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, workers)
}

func (h *Handler) AddUser(c *gin.Context) {
	var inp domain.UserData
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	err := h.services.Admin.AddUser(c, inp)
	if err != nil {
		if errors.Is(err, domain.ErrReplayUsername) {
			c.JSON(http.StatusOK, "Такой username уже используется")
			return
		}
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Admin.DeleteUser(c, userID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}