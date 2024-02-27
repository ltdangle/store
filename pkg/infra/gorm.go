package infra

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Gorm(postgresUrl string) *gorm.DB {

	// Init db (GORM)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer that std sql logs will write to
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(postgresUrl), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Sqlx(postgresUrl string) *sqlx.DB {
	sqlx, err := sqlx.Open("postgres", postgresUrl)
	if err != nil {
		panic("failed to connect database")
	}

	return sqlx
}
