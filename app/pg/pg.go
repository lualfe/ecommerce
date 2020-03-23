package pg

import (
	"github.com/jinzhu/gorm"
	"github.com/lualfe/ecommerce/models"
)

// Postgres model
type Postgres struct {
	Instance *gorm.DB
}

func (p *Postgres) Migrate() {
	p.Instance.AutoMigrate(&models.User{})
	p.Instance.AutoMigrate(&models.Address{})
	p.Instance.AutoMigrate(&models.BillingAddress{})

	p.Instance.Model(&models.Address{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	p.Instance.Model(&models.BillingAddress{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}
