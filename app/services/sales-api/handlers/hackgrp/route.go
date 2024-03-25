package hackgrp

import (
	"github.com/draculaas/deploy/business/web/v1/auth"
	"github.com/draculaas/deploy/business/web/v1/mid"
	"github.com/draculaas/deploy/core/web"
	"net/http"
)

type Config struct {
	Auth *auth.Auth
}

// Routes ... Add route for this specific group
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.Auth)
	ruleAdmin := mid.Authorize(cfg.Auth, auth.RuleAdminOnly)

	app.Handle(http.MethodGet, version, "/hack", Hack)
	app.Handle(http.MethodGet, version, "/hackauth", Hack, authen, ruleAdmin)
}
