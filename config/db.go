package config

import (
	"fmt"
	"form/models"
	"log"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	// 1. Get and Trim Environment Variables
	user := strings.TrimSpace(os.Getenv("DB_USER"))
	pass := strings.TrimSpace(os.Getenv("DB_PASSWORD"))
	host := strings.TrimSpace(os.Getenv("DB_HOST"))
	port := strings.TrimSpace(os.Getenv("DB_PORT"))
	name := strings.TrimSpace(os.Getenv("DB_NAME"))

	// 2. Build DSN (Data Source Name)
	// IMPORTANT: SkySQL Serverless requires tls=true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=true",
		user, pass, host, port, name)

	// Example using os.Getenv
	// dsn := os.Getenv("DATABASE_URL")

	// dsn := "root:root@tcp(localhost:3306)/kyc_demo?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "dbpgf21238876:D}ggV6AzHJLg2Tr7Q8KVz@tcp(serverless-us-east1.sysp0000.db2.skysql.com: 4014)/kyc_demo?charset=utf8mb4&parseTime=True&loc=Local"

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
