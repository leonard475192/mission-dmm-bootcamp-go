package statuses

import (
	"net/http"

	"yatter-backend-go/app/app"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi"
)

// Implementation of handler
type handler struct {
	app *app.App
}

// Create Handler for `/v1/accounts/`
func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()

	h := &handler{app: app}
	r.Get("/{id}", h.Get)

	loginGroup := r.Group(nil)
	loginGroup.Use(auth.Middleware(app))
	loginGroup.Post("/", h.Create)
	loginGroup.Delete("/{id}", h.Delete)

	return r
}
