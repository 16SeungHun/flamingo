package interfaces

import (
	"context"
	"net/url"

	"flamingo.me/flamingo/core/auth/application"
	"flamingo.me/flamingo/framework/web"
	"flamingo.me/flamingo/framework/web/responder"
	"github.com/satori/go.uuid"
)

type (
	LoginControllerInterface interface {
		Get(context.Context, *web.Request) web.Response
	}

	// LoginController handles the login redirect
	LoginController struct {
		responder.RedirectAware
		authManager *application.AuthManager
	}
)

// Inject LoginController dependencies
func (l *LoginController) Inject(
	redirectAware responder.RedirectAware,
	authManager *application.AuthManager,
) {
	l.RedirectAware = redirectAware
	l.authManager = authManager
}

// Get handler for logins (redirect)
func (l *LoginController) Get(c context.Context, request *web.Request) web.Response {
	redirecturl, ok := request.Param1("redirecturl")
	if !ok || redirecturl == "" {
		redirecturl = request.Request().Referer()
	}

	if refURL, err := url.Parse(redirecturl); err != nil || refURL.Host != request.Request().Host {
		u, _ := l.authManager.URL(c, "")
		redirecturl = u.String()
	}

	state := uuid.NewV4().String()
	request.Session().Store("auth.state", state)
	request.Session().Store("auth.redirect", redirecturl)

	return l.RedirectURL(l.authManager.OAuth2Config(c).AuthCodeURL(state))
}
