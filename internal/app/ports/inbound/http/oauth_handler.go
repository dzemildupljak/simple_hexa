package httphdl

import (
	"fmt"
	"github.com/dzemildupljak/simple_hexa/internal/app/application"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"net/http"
)

type OauthHttpHandler struct {
	oauthService application.OAuthService
}

func NewOauthHttpHandler(service application.OAuthService) *OauthHttpHandler {
	return &OauthHttpHandler{
		oauthService: service,
	}
}

func (h *OauthHttpHandler) RegisterHandlers(router *mux.Router, nrApp *newrelic.Application) {
	// Create a sub router for api/v1/oauth
	oAuthRouter := setupOauthVersionedRouter(router, "v1")
	h.setupOauthRoutes(oAuthRouter)
}

func setupOauthVersionedRouter(router *mux.Router, version string) *mux.Router {
	return router.PathPrefix(fmt.Sprintf("/api/%s/oauth", version)).Subrouter()
}

func (h *OauthHttpHandler) setupOauthRoutes(oAuthRouter *mux.Router) {
	//oAuthRouter.HandleFunc("", h.GetAllOauthsHandler).Methods("GET")
}

func (h *OauthHttpHandler) OauthLoginHandler(w http.ResponseWriter, r *http.Request) {}
