package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	// dsn := "db_host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db_host := os.Getenv("DBHOST")
	db_user := os.Getenv("DBUSER")
	db_password := os.Getenv("DBPASSWORD")
	db_name := os.Getenv("DBNAME")
	db_port := os.Getenv("DBPORT")
	db_sslmode := os.Getenv("DBSSLMODE")
	tz := os.Getenv("DBTZ")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", db_host, db_user, db_password, db_name, db_port, db_sslmode, tz)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %s", err)
	}
	// populate global DB variable
	DB = db
}

func DBMigrate() {

}
