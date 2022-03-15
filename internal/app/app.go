package app

import (
	"net/http"
	"os"
	//"crypto/tls"

	"github.com/joho/godotenv"
	"github.com/Traking-work/traking-backend.git/pkg/handler"
	"github.com/Traking-work/traking-backend.git/pkg/logging"
	"github.com/Traking-work/traking-backend.git/pkg/repository"
	"github.com/Traking-work/traking-backend.git/pkg/service"
	"github.com/spf13/viper"
)

type Server struct {
	httpServer *http.Server
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func (s *Server) Run() error {
	logger := logging.GetLogger()

	if err := initConfig(); err != nil {
		logger.Fatalf("Error initializing configs: %s", err.Error())
	} else {
		logger.Info("Success initializing configs")
	}

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Error loading env variables: %s", err.Error())
	} else {
		logger.Info("Success loading env variables")
	}

	db, err := repository.NewMongoDB(os.Getenv("MONDO_DB_URL"))
	if err != nil {
		logger.Fatalf("Error connect mongodb: %s", err.Error())
	} else {
		logger.Info("Success connect mongodb")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	routes := handlers.InitRoutes()

	//cert, _ := tls.LoadX509KeyPair("fullchain.pem", "privkey.pem")

	s.httpServer = &http.Server{
		Addr:           ":" + viper.GetString("http.port"),
		Handler:        routes,
		MaxHeaderBytes: viper.GetInt("http.maxHeaderBytes"),
		ReadTimeout:    viper.GetDuration("http.readTimeout"),
		WriteTimeout:   viper.GetDuration("http.writeTimeout"),
		//TLSConfig: &tls.Config{
		//	Certificates: []tls.Certificate{cert},
		//},
	}

	logger.Info("Listen server...")
	return s.httpServer.ListenAndServe()
	//return s.httpServer.ListenAndServeTLS("", "")
}
