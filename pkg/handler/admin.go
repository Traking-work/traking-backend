package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Traking-work/traking-backend.git/internal/domain"
)

func (h *Handler) GetDataAdmin(c *gin.Context) {
	var inp domain.NewUser
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

func (h *Handler) AddUser(c *gin.Context) {
	var inp domain.NewUser
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