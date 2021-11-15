package models

type DataBaseConfig struct {
	Host         string
	UserName     string
	Password     string
	DataBaseName string
}

type Config struct {
	DataBase DataBaseConfig
	WebHost  string
}
