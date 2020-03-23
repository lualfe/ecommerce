package models

// Address model
type Address struct {
	ID                int    `json:"id"`
	UserID            int    `gorm:"type:int" json:"user_id"`
	Address           string `gorm:"type:varchar(30)" json:"address" form:"address"`
	AddressNumber     string `gorm:"type:varchar(10)" json:"address_number" form:"address_number"`
	AddressComplement string `gorm:"type:varchar(10)" json:"address_complement" form:"address_complement"`
	District          string `gorm:"type:varchar(10)" json:"district" form:"district"`
	State             string `gorm:"type:varchar(2)" json:"state" form:"state"`
	City              string `gorm:"type:varchar(10)" json:"city" form:"city"`
	ZIP               string `gorm:"type:varchar(9)" json:"zip" form:"zip"`
}
