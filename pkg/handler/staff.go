package handler

import (
	"net/http"
	"time"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetDataUser(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	dataUser, err := h.services.Staff.GetDataUser(c, userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Get data user %s", c.Param("ID"))
	c.JSON(http.StatusOK, dataUser)
}

func (h *Handler) GetAccounts(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inp domain.Date
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	accounts, err := h.services.Staff.GetAccounts(c, userID, inp.Date)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Get accounts %s", c.Param("ID"))

	c.JSON(http.StatusOK, accounts)
}

func (h *Handler) AddAccount(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inp domain.AccountData
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	inp.UserID = userID
	inp.Percent = 0.25
	inp.CreateDate = time.Now()
	inp.StatusDelete = false

	if err := h.services.Staff.AddAccount(c, inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Add account %s", c.Param("ID"))
}

func (h *Handler) GetDataAccount(c *gin.Context) {
	accountID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inp domain.AccountPack
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	dataAccount, err := h.services.Staff.GetDataAccount(c, accountID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	packsAccount, err := h.services.Staff.GetPacksAccount(c, accountID, inp.Date)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	income := float32(0)
	for _, pack := range packsAccount {
		income += float32(pack.CountTask) * float32(pack.Payment)
	}
	dataAccount.Income = income

	h.logger.Infof("Get data account %s", c.Param("ID"))

	c.JSON(http.StatusOK, map[string]interface{}{
		"packsAccount": packsAccount,
		"dataAccount":  dataAccount,
	})
}

func (h *Handler) AddPack(c *gin.Context) {
	accountID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inp domain.AccountPack
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	if err := h.services.Staff.AddPack(c, accountID, inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Add pack %s", c.Param("ID"))
}

func (h *Handler) UpgradePack(c *gin.Context) {
	packID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inp domain.AccountPack
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	if err := h.services.Staff.UpgradePack(c, packID, inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Update pack %s", c.Param("ID"))
}

func (h *Handler) ApprovePack(c *gin.Context) {
	packID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Staff.ApprovePack(c, packID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Approve pack %s", c.Param("ID"))
}

func (h *Handler) DeletePack(c *gin.Context) {
	packID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Staff.DeletePack(c, packID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Delete pack %s", c.Param("ID"))
}

func (h *Handler) DeleteAccount(c *gin.Context) {
	accountID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Staff.DeleteAccount(c, accountID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Delete account %s", c.Param("ID"))
}

func (h *Handler) GetParamsMain(c *gin.Context) {
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

	var incomeAll float32
	var incomeAdmin float32

	if inp.Position == "staff" {
		incomeAll, incomeAdmin, err = h.services.Staff.GetParamsMainStaff(c, userID)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	} else if inp.Position == "teamlead" {
		incomeAll, incomeAdmin, err = h.services.Staff.GetParamsMainTeamlead(c, userID)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		incomeAll, incomeAdmin, err = h.services.Staff.GetParamsMainAdmin(c, userID)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	h.logger.Infof("Get params main staff %s", c.Param("ID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"incomeAll":   incomeAll,
		"incomeAdmin": incomeAdmin,
	})
}

func (h *Handler) GetParamsDate(c *gin.Context) {
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

	var incomeAll float32
	var incomeAdmin float32

	if inp.Position == "staff" {
		incomeAll, incomeAdmin, err = h.services.Staff.GetParamsDateStaff(c, userID, inp.Date)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	} else if inp.Position == "teamlead" {
		incomeAll, incomeAdmin, err = h.services.Staff.GetParamsDateTeamlead(c, userID, inp.Date)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		incomeAll, incomeAdmin, err = h.services.Staff.GetParamsDateAdmin(c, userID, inp.Date)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	h.logger.Infof("Get params date staff %s", c.Param("ID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"incomeAll":   incomeAll,
		"incomeAdmin": incomeAdmin,
	})
}

func (h *Handler) ChangeTeamlead(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inp domain.UserTeamlead
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	if err := h.services.Staff.ChangeTeamlead(c, userID, inp.TeamLead); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Change teamlead from %s to %s", c.Param("ID"), inp.TeamLead)
}