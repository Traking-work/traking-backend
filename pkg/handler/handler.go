package handler

import (
	"os"

	"github.com/Traking-work/traking-backend.git/pkg/logging"
	"github.com/Traking-work/traking-backend.git/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL"), "http://703449-cv99377.tmweb.ru"},
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
			admin.GET("/:ID/get-count-workers", h.GetCountWorkers)
			admin.GET("/:ID/get-workers", h.GetWorkers)
			admin.POST("/add-user", h.AddUser)
			admin.POST("/:ID/delete-user", h.DeleteUser)
		}

		teamlead := api.Group("/teamlead", h.userIdentity)
		{
			teamlead.GET("/:userID/get-staff", h.GetStaff)
		}

		staff := api.Group("/staff", h.userIdentity)
		{
			staff.GET("/:ID/get-data-user", h.GetDataUser)
			staff.POST("/:ID/get-accounts", h.GetAccounts)
			staff.POST("/:ID/add-account", h.AddAccount)
			staff.POST("/:ID/change-teamlead", h.ChangeTeamlead)
			staff.POST("/:ID/get-data-account", h.GetDataAccount)
			staff.POST("/:ID/add-pack", h.AddPack)
			staff.POST("/:ID/upgrade-pack", h.UpgradePack)
			staff.GET("/:ID/delete-pack", h.DeletePack)
			staff.GET("/:ID/delete-account", h.DeleteAccount)
			staff.POST("/:ID/get-income", h.GetIncome)
			staff.POST("/:ID/get-employee-rating", h.GetEmployeeRating)
		}
	}

	return router
}
