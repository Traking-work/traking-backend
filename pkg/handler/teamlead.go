package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetStaff(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	staff, err := h.services.Teamlead.GetStaff(c, userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Get staff %s", c.Param("userID"))
	c.JSON(http.StatusOK, staff)
}
