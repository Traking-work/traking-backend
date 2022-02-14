package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/Traking-work/traking-backend.git/internal/domain"
)

func (h *Handler) Login(c *gin.Context) {
	var inp domain.UserLogin
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	res, err := h.services.Authorization.Login(c, inp.Username, inp.Password)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    res.RefreshToken,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HttpOnly: true,
	})

	h.logger.Infof("Login admin %s", res.UserID)

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": res.AccessToken,
		"position": res.Position,
		//"refreshToken": res.RefreshToken,
		"userID": res.UserID,
	})
}

func (h *Handler) Refresh(c *gin.Context) {
	cookie, err := c.Request.Cookie("refreshToken")
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	refreshToken := cookie.Value

	res, err := h.services.Authorization.Refresh(c, refreshToken)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    res.RefreshToken,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HttpOnly: true,
	})

	h.logger.Infof("Refresh admin %s", res.UserID)

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": res.AccessToken,
		"position": res.Position,
		//"refreshToken": res.RefreshToken,
		"userID": res.UserID,
	})
}

func (h *Handler) Logout(c *gin.Context) {
	cookie, err := c.Request.Cookie("refreshToken")
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	refreshToken := cookie.Value

	if err := h.services.Authorization.Logout(c, refreshToken); err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "refreshToken",
		Value:  "",
		MaxAge: -1,
	})

	h.logger.Infof("Logout admin %s", refreshToken)

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
