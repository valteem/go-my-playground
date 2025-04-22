package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"webapi/product-catalog/api"
	"webapi/product-catalog/config"
	"webapi/product-catalog/repository"
	"webapi/product-catalog/server"
	"webapi/product-catalog/services"
	"webapi/product-catalog/sqldb"

	"github.com/gin-gonic/gin"
)

func Run(configPath string) {

	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config info: %v", err)
	}

	pg, err := sqldb.New(cfg.PG.URL, sqldb.MaxPoolSize(cfg.PG.MaxPoolSize))
	if err != nil {
		log.Fatalf("failed to establish database connection: %v", err)
	}
	defer pg.Close()

	repo := repository.NewRepositories(pg)

	deps := &services.ServiceDependencies{
		Repositories: repo,
	}
	services := services.NewServices(deps)

	handler := &gin.Engine{}
	api.NewRouter(handler, services)

	server := server.New(handler, server.Addr(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("app interrupted by signal %q", s.String())
	case err = <-server.Notify():
		log.Printf("http server notifies %v", err)

	}

	err = server.Shutdown()
	if err != nil {
		log.Fatalf("failed to shut down server: %v", err)
	}

}
