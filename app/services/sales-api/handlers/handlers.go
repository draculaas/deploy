package handlers

import (
	"github.com/dimfeld/httptreemux/v5"
	"github.com/draculaas/deploy/app/services/sales-api/handlers/hackgrp"
	v1 "github.com/draculaas/deploy/business/web/v1"
)

type Routers struct{}

// Add implements the RouterAdder interface to add all routes.
func (Routers) Add(mux *httptreemux.ContextMux, apiCfg v1.APIMuxConfig) {
	hackgrp.Routes(mux)
}
