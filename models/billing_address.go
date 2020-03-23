package models

// BillingAddress model
type BillingAddress struct {
	ID                       int    `json:"id"`
	UserID                   int    `gorm:"type:int;UNIQUE" json:"user_id"`
	BillingAddress           string `gorm:"type:varchar(30)" json:"billing_address" form:"billing_address"`
	BillingAddressNumber     string `gorm:"type:varchar(10)" json:"billing_address_number" form:"billing_address_number"`
	BillingAddressComplement string `gorm:"type:varchar(10)" json:"billing_address_complement" form:"billing_address_complement"`
	BillingDistrict          string `gorm:"type:varchar(10)" json:"billing_district" form:"billing_district"`
	BillingState             string `gorm:"type:varchar(2)" json:"billing_state" form:"billing_state"`
	BillingCity              string `gorm:"type:varchar(10)" json:"billing_city" form:"billing_city"`
	BillingZIP               string `gorm:"type:varchar(9)" json:"billing_zip" form:"billing_zip"`
}
