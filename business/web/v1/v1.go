package v1

import (
	"github.com/draculaas/deploy/business/web/v1/auth"
	"github.com/draculaas/deploy/business/web/v1/mid"
	"github.com/draculaas/deploy/core/logger"
	"github.com/draculaas/deploy/core/web"
	"os"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Build    string
	Shutdown chan os.Signal
	Log      *logger.Logger
	Auth     *auth.Auth
}

// RouteAdder defines behavior that sets the routes to bind for an instance
// of the service.
type RouteAdder interface {
	Add(app *web.App, cfg APIMuxConfig)
}

// APIMux ...Create a http.Handler with all app routes defined
func APIMux(cfg APIMuxConfig, routeAdder RouteAdder) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Metrics(), mid.Panics())

	routeAdder.Add(app, cfg)

	return app
}
