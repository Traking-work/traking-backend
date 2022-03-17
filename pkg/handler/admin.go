package handler

import (
	"errors"
	"net/http"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetTeamLeads(c *gin.Context) {
	teamleads, teamleadCreate, err := h.services.Admin.GetTeamLeads(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Info("Get teamleads")
	c.JSON(http.StatusOK, map[string]interface{}{
		"teamleads":      teamleads,
		"teamleadCreate": teamleadCreate,
	})
}

func (h *Handler) GetCountWorkers(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	countWorkers, err := h.services.Admin.GetCountWorkers(c, userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Get count workers %s", c.Param("ID"))
	c.JSON(http.StatusOK, countWorkers)
}

func (h *Handler) GetWorkers(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	workers, err := h.services.Admin.GetWorkers(c, userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Get workers %s", c.Param("ID"))
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
	h.logger.Infof("Add user %s", inp.Username)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inp domain.DataForParams
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	if err = h.services.Admin.DeleteUser(c, userID, inp.Position); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	h.logger.Infof("Delete user %s", c.Param("ID"))
}
