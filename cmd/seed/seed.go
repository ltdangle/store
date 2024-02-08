package main

import (
	"log"
	"os"
	"store/pkg/models"
	"time"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type config struct {
	MYSQL_DSN    string `env:"MYSQL_DSN"`
	POSTGRES_URL string `env:"POSTGRES_URL"`
}

func main() {
	// Load .env config.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err.Error())
	}

	// Init db (GORM)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer that std sql logs will write to
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(cfg.POSTGRES_URL), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}

	// Init seeder.
	seeder := models.NewSeeder(db)

	customer := seeder.CreateCustomer()
	cart := seeder.AddCart(customer)
	cartItem := seeder.AddCartItem(cart)
	product := seeder.CreateProduct()
	seeder.AddProduct(cartItem, product)

}
