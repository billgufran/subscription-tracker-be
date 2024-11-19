package database

import (
	"fmt"
	"log"
	"os"

	"subscription-tracker/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	log.Println("Starting database initialization...")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
	)

	log.Printf("Connecting to database with DSN: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Successfully connected to database")

	// Register custom types. It will set the ULID to the field if it is empty.
	db.Callback().Create().Before("gorm:create").Register("set_ulid", func(tx *gorm.DB) {
		if tx.Statement.Schema != nil {
			for _, field := range tx.Statement.Schema.Fields {
				if field.FieldType.Name() == "ULID" {
					_, isZero := field.ValueOf(tx.Statement.Context, tx.Statement.ReflectValue)
					if isZero {
						field.Set(tx.Statement.Context, tx.Statement.ReflectValue, models.NewULID())
					}
				}
			}
		}
	})

	// Auto-migrate the schema
	log.Println("Starting database migration...")
	err = db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Currency{},
		&models.PaymentMethod{},
		&models.Subscription{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed successfully")

	// Seed default data
	log.Println("Starting to seed default data...")
	seedDefaultData(db)
	log.Println("Default data seeding completed")

	return db
}

func seedDefaultData(db *gorm.DB) {
	// Seed default categories if they don't exist
	for _, category := range models.DefaultCategories {
		var count int64
		db.Model(&models.Category{}).
			Where("name = ? AND is_default = ?", category.Name, true).
			Count(&count)

		if count == 0 {
			db.Create(&category)
		}
	}

	// Seed default currencies if they don't exist
	for _, currency := range models.DefaultCurrencies {
		var count int64
		db.Model(&models.Currency{}).
			Where("code = ?", currency.Code).
			Count(&count)

		if count == 0 {
			db.Create(&currency)
		}
	}
}
