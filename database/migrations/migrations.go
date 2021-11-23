package migrations

import (
	"github.com/jadson-medeiros/challengehash/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(models.Product{})
}
