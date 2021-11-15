package main

import (
	"encoding/json"
	"io/ioutil"
	"logistics/api"
	"logistics/models"
	"logistics/repositories"
	"logistics/services"
	"logistics/setup"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

func main() {
	rawByteConfig, err := ioutil.ReadFile("config.json")
	if err != nil {
		logrus.Info("err reading config: ", err)
	}

	// Parse config
	var config models.Config
	err = json.Unmarshal(rawByteConfig, &config)
	if err != nil {
		logrus.Info("err parsing config: ", err)
	}

	pgDB := setup.PostgresConnect(&config.DataBase)
	// pgDB.AddQueryHook(dbLogger{})

	authRepo := repositories.NewAuthRepo(pgDB)

	handler := setup.HTTPServer()
	gr := handler.Group("v1/api")

	authService := services.NewAuthService(authRepo)

	api.Auth(gr, authService)
	api.TransporUnits(gr)

	logrus.Fatal(handler.Run(config.WebHost))
}
