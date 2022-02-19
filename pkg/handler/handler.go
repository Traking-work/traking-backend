package handler

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/Traking-work/traking-backend.git/pkg/logging"
	"github.com/Traking-work/traking-backend.git/pkg/service"
)

type Handler struct {
	services *service.Service
	logger   logging.Logger
}

func NewHandler(services *service.Service) *Handler {
	logger := logging.GetLogger()
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Error loading env variables: %s", err.Error())
	}

	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type,access-control-allow-origin, access-control-allow-headers,authorization,my-custom-header"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
	}))

	api := router.Group("/api")
	{
		api.POST("/login", h.Login)
		api.GET("/refresh", h.Refresh)
		api.GET("/logout", h.Logout)

		admin := api.Group("/admin", h.userIdentity)
		{
			admin.GET("/get-teamleads", h.GetTeamLeads)
			admin.POST("/add-user", h.AddUser)
			admin.GET("/:userID/delete-user", h.DeleteUser)
		}

		teamlead := api.Group("/teamlead", h.userIdentity)
		{
			teamlead.GET("/:userID/get-staff", h.GetStaff)
		}

		staff := api.Group("/staff", h.userIdentity)
		{
			staff.GET("/:userID/get-accounts", h.GetAccounts)
			staff.POST("/:userID/add-account", h.AddAccount)
			staff.GET("/:userID/:accountID", h.GetDataAccount)
			staff.POST("/:userID/:accountID/add-pack", h.AddPack)
			staff.GET("/:userID/:accountID/delete-account", h.DeleteAccount)
		}
	}

	return router
}
