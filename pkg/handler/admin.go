package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Traking-work/traking-backend.git/internal/domain"
)

func (h *Handler) GetTeamLeads(c *gin.Context) {
	var inp domain.UserID
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	teamleads, err := h.services.Admin.GetTeamLeads(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, teamleads)
}

func (h *Handler) AddUser(c *gin.Context) {
	var inp domain.UserData
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	var err error

	if inp.Position == "teamlead" {
		err = h.services.Admin.AddUser(c, inp)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		err = h.services.Admin.AddUser(c, inp)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}
}