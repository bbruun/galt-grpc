package initializers

import (
	"log"

	"github.com/bbruun/galt/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.Machine{})
	if err != nil {
		log.Fatalf("error migrating database: %s", err)
	}
}
