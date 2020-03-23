package web

import "github.com/lualfe/ecommerce/app"

// Web model
type Web struct {
	App *app.App
}

// NewWeb initializes an Web instance
func NewWeb(a *app.App) *Web {
	return &Web{a}
}
