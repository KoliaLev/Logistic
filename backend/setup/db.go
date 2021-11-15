package setup

import (
	"fmt"
	"logistics/models"
	"os"

	"github.com/go-pg/pg/v10"
	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

func MigrateDB(config *models.DataBaseConfig) {
	// dbConf := conf.DataBase
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s", config.UserName, config.Password, config.Host, config.DataBaseName)
	migration, err := migrate.New("file://migrations", connStr+"?sslmode=disable")
	if err != nil {
		logrus.Fatal("err conection to data base ", err)
	}

	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		println(err.Error())
		os.Exit(1)
	} else if err == migrate.ErrNoChange {
		println("Database schema is up to date")
	} else {
		println("Successfully migrated database schema")
	}

}

func PostgresConnect(config *models.DataBaseConfig) *pg.DB {
	MigrateDB(config)
	return pg.Connect(&pg.Options{
		Addr:     config.Host,
		User:     config.UserName,
		Password: config.Password,
		Database: config.DataBaseName,
	})
}
