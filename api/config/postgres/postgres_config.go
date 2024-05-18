package postgres_config

import (
	"api/coworkings/infrastructure/postgres"
	"api/utils/logger"
	"fmt"

	"github.com/eldimious/golang-api-showcase/utils/env"
	gorm_postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	URL      string
	DB       *gorm.DB
}

var MIGRATIONS_PATH = "../migrations"

func NewConfig() *Config {
	env.CheckDotEnv()

	database_host := env.MustGet("DATABASE_HOST")
	database_port := env.MustGet("DATABASE_PORT")
	database_user := env.MustGet("DATABASE_USER")
	database_name := env.MustGet("DATABASE_DB")
	database_password := env.MustGet("DATABASE_PASSWORD")
	database_url := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", database_user, database_password, database_host, database_name)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		database_host, database_port, database_user, database_name, database_password)

	logger.GetLogger().Info("Connecting to postgres...", nil)
	db, err := gorm.Open(gorm_postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to db: " + err.Error())
	}

	logger.GetLogger().Info("Running migrations...", nil)
	err = db.AutoMigrate(&postgres.PostgresCoworking{}, &postgres.PostgresWorkspace{}, &postgres.CoworkingImage{})
	if err != nil {
		panic("Error running migrations: " + err.Error())
	}

	return &Config{
		Host:     database_host,
		Port:     database_port,
		User:     database_user,
		DBName:   database_name,
		Password: database_password,
		URL:      database_url,
		DB:       db,
	}
}
