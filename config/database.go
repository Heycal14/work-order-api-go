package config

import (
	"fmt"
	"log"
	"os"

	"work-order-api-go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := getEnv("DB_HOST", "127.0.0.1")
	port := getEnv("DB_PORT", "3306")
	user := getEnv("DB_USER", "root")
	pass := getEnv("DB_PASSWORD", "")
	name := getEnv("DB_NAME", "work_order_api")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"),
		user,pass,host,port,name

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nill {
		log.fatal("gagal koneksi database", err)
	}

	if err := db.AutoMigrate(&models.WorkOrder{}); err != nil{
		log.Fatal("gagal migrate: ", err)
	}

	DB = db
	fmt.Println("Database Connected & mingrated")
}
func getEnv(key, fallback string) string{
	if value, ok := os.LookupEnv(key); ok{
		retunt value
	}
	return fallback
}
