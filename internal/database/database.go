package database

import (
	"log"

	"subscription-tracker/internal/config"
	"subscription-tracker/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) *gorm.DB {
	log.Println("Starting database initialization...")
	dsn := cfg.Database.GetDSN()

	log.Printf("Connecting to database with DSN: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Successfully connected to database")

	// Register custom types to gorm. It will set the ULID to the field if it is empty.
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
		&models.BillingCycle{},
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
			Where("name = ? AND system_defined = ?", category.Name, true).
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

	// Seed default billing cycles if they don't exist
	for _, billingCycle := range models.DefaultBillingCycles {
		var count int64
		db.Model(&models.BillingCycle{}).
			Where("name = ? AND system_defined = ?", billingCycle.Name, true).
			Count(&count)

		if count == 0 {
			db.Create(&billingCycle)
		}
	}
}
