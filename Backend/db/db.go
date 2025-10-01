package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Replace these values with your Supabase database credentials
	host := "db.opvifyrwxktzyuyskhme.supabase.co"
	port := "5432"
	user := "postgres"
	password := "9VFaBW19fb1HwIa4"
	dbname := "postgres"

	// DSN for Postgres connection
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	DB = db
	fmt.Println("✅ Connected to Supabase Postgres successfully")
}
