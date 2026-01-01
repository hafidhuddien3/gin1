package db

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gin-quickstart/models"
)

var DB *gorm.DB

func InitDB() {
	dsn := "postgresql://neondb_owner:npg_XP28nmjWhEpg@ep-hidden-feather-a1zkp1zz-pooler.ap-southeast-1.aws.neon.tech/neondb?sslmode=require"
    // dsn := "postgresql://neondb_owner:password@host/neondb?sslmode=require"
	log.Println("Connecting to Neon database...")
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB.AutoMigrate(&models.Book{})
    DB.AutoMigrate(&models.Ingredient{})
    DB.AutoMigrate(&models.ProductIngredient{})
    DB.AutoMigrate(&models.Product{})
}
