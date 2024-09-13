package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cronnoss/tk-api/internal/app"
	"github.com/cronnoss/tk-api/internal/logger"
	internalhttp "github.com/cronnoss/tk-api/internal/server/http"
	"github.com/cronnoss/tk-api/internal/storage"
)

func main() {
	conf := NewConfig().TicketConf
	storage := storage.NewStorage(conf.Storage)
	logger := logger.NewLogger(conf.Logger.Level, os.Stdout)
	ticket, _ := app.NewTicket(logger, conf, storage)
	httpsrv := internalhttp.NewServer(logger, ticket, conf.HTTP.Host, conf.HTTP.Port)

	ticket.Run(httpsrv)

	filename := filepath.Base(os.Args[0])
	fmt.Printf("%s stopped\n", filename)
}
