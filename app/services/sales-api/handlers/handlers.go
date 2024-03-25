package handlers

import (
	"github.com/draculaas/deploy/app/services/sales-api/handlers/checkgroup"
	"github.com/draculaas/deploy/app/services/sales-api/handlers/hackgrp"
	v1 "github.com/draculaas/deploy/business/web/v1"
	"github.com/draculaas/deploy/core/web"
)

type Routers struct{}

// Add implements the RouterAdder interface to add all routes.
func (Routers) Add(app *web.App, apiCfg v1.APIMuxConfig) {
	hackgrp.Routes(app, hackgrp.Config{
		Auth: apiCfg.Auth,
	})

	checkgroup.Routes(app, checkgroup.Config{
		Build: apiCfg.Build,
		Log:   apiCfg.Log,
	})
}
