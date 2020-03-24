package pg

import (
	"github.com/jinzhu/gorm"
	"github.com/lualfe/ecommerce/models"

	//needed for postgres initialization
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// Necessary in order to check for transaction retry error codes.
	_ "github.com/lib/pq"
)

// NewPostgres initializes a postgres instance
func NewPostgres(conn string) *Postgres {
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	pg := &Postgres{
		Instance: db,
	}
	return pg
}

// Postgres model
type Postgres struct {
	Instance *gorm.DB
}

// Migrate makes all postgres migrations and binds the foreign keys
func (p *Postgres) Migrate() {
	p.Instance.AutoMigrate(&models.User{})
	p.Instance.AutoMigrate(&models.Address{})
	p.Instance.AutoMigrate(&models.BillingAddress{})

	p.Instance.Model(&models.Address{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	p.Instance.Model(&models.BillingAddress{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}
