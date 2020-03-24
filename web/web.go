package web

import "github.com/lualfe/ecommerce/app"

// Web model
type Web struct {
	PG app.Repository
	MG app.Migration
}

// NewWeb initializes an Web instance
func NewWeb(pg app.Repository, mg app.Migration) *Web {
	return &Web{pg, mg}
}
