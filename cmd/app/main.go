package main

import (
	"github.com/Traking-work/traking-backend.git/internal/app"
	"github.com/Traking-work/traking-backend.git/pkg/logging"
)

func main() {
	logger := logging.GetLogger()

	server := new(app.Server)
	if err := server.Run(); err != nil {
		logger.Errorf("Error run server %s", err.Error())
	}
}
