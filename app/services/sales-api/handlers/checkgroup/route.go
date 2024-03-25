package checkgroup

import (
	"github.com/draculaas/deploy/core/logger"
	"github.com/draculaas/deploy/core/web"
	"net/http"
)

// Config contains all the mandatory systems required by handlers
type Config struct {
	Build string
	Log   *logger.Logger
}

// Routes adds specific routes for this group.

func Routes(app *web.App, cfg Config) {
	const version = "v1"

	h := New(cfg.Build, cfg.Log)
	app.HandleNoMiddleware(http.MethodGet, version, "/readiness", h.Readiness)
	app.HandleNoMiddleware(http.MethodGet, version, "/liveness", h.Liveness)
}
