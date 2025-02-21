package connection_todoSM

import (
	"fmt"
	"log"

	model_todoSM "github.com/rrraf1/to-do_social_media/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host, User, Password, DBName string
}

func Migrate(db *gorm.DB) error {
	if err := model_todoSM.MigratePost(db); err != nil {
		log.Fatal(err)
		return gorm.ErrCheckConstraintViolated
	}
	return nil
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.User, config.Password, config.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
