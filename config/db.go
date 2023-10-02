package config

import (
	"backend/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	var err error
	var dsn = "host=ln.ichigozdata.win user=postgres dbname=projectln port=5432 password=postgres"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("postgres err: " + err.Error())
		panic(err)
	}
	migrateDB()
	return err
}

func migrateDB() {
	DB.AutoMigrate(models.Channel{})
	DB.AutoMigrate(models.Template{})
	DB.AutoMigrate(models.Savety{})
	DB.AutoMigrate(models.Saveunit{})
	DB.AutoMigrate(models.NodeMCU{})
}
