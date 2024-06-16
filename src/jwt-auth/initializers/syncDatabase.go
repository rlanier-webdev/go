package initializers

import (
	"jwt-auth/models"
)

func SyncDatabase() {
	// Migrate the schema
	DB.AutoMigrate(&models.User{})
}
