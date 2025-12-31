package config

import (
	"form/models"
	"log"
"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	// Example using os.Getenv
 dsn := os.Getenv("DATABASE_URL")  
	
	// dsn := "root:root@tcp(localhost:3306)/kyc_demo?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "dbpgf21238876:D}ggV6AzHJLg2Tr7Q8KVz@tcp(serverless-us-east1.sysp0000.db2.skysql.com: 4014)/kyc_demo?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}

	DB = database
	log.Println("✅ Database Connected!")

	// Auto migrate models
	DB.AutoMigrate(
		&models.Identity{},
		&models.Address{},
		&models.Promoter{},
		&models.Bank{},
		&models.Trading{},
		&models.Additional{},
	)
}
