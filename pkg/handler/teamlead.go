package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Traking-work/traking-backend.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetStaff(c *gin.Context) {
	var inp domain.UserID
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	userID, err := primitive.ObjectIDFromHex(inp.ID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	staff, err := h.services.Teamlead.GetStaff(c, userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, staff)
}