/*
Flamingo Package that provides a healthcheck endpoint under the default route /status/healthcheck
Usage:
 Register your own Status via Dingo:
 injector.BindMap((*healthcheck.Status)(nil), "yourServiceName").To(yourServiceNameApi.Status{})

*/
package healthcheck

import (
	"flamingo.me/flamingo/core/healthcheck/interfaces/controllers"
	"flamingo.me/flamingo/framework/dingo"
	"flamingo.me/flamingo/framework/router"
)

type Module struct{}

func (m *Module) Configure(injector *dingo.Injector) {
	router.Bind(injector, new(routes))
}

type routes struct {
	healthcheck *controllers.Healthcheck
}

func (r *routes) Inject(healthcheck *controllers.Healthcheck) {
	r.healthcheck = healthcheck
}

func (r *routes) Routes(registry *router.Registry) {
	registry.HandleGet("healthcheck", r.healthcheck.Get)
	registry.Route("/status/healthcheck", "healthcheck")
}
