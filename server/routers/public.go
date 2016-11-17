package routers

import "github.com/sstutz/pastestuff/server/routers/handlers"

func RegisterPublicRoutes(r *Router) {
	r.Get("/", handlers.Welcome)
}
