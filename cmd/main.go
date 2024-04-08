package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andy-ahmedov/makves/internal/config"
	"github.com/andy-ahmedov/makves/internal/repository/postgres"
	"github.com/andy-ahmedov/makves/internal/service"
	"github.com/andy-ahmedov/makves/internal/transport/rest"
	"github.com/andy-ahmedov/makves/pkg/client"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.ErrorLevel)
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

	postgresDB, err := client.ConnectToDB(cfg.PostgresDB)
	if err != nil {
		log.Fatal(err)
	}

	postgresRepository := postgres.NewMakvesRepository(postgresDB)
	MakvesService := service.NewMakvesService(postgresRepository)
	handler := rest.NewHandler(MakvesService)

	srv := &http.Server{
		Addr:    ":9000",
		Handler: handler.InitGinRouter(),
	}

	logrus.Info("SERVER STARTED")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
