package routers

import (
	"net"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sstutz/pastestuff/server/assets"
	. "github.com/sstutz/pastestuff/server/config"
)

type Router struct {
	*mux.Router
}

func ListenAndServe() {
	r := &Router{mux.NewRouter()}

	RegisterPublicRoutes(r)

	fs := http.FileServer(assets.LoadAssets().Http())

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	s := net.JoinHostPort(Config.Server.Address, strconv.Itoa(Config.Server.Port))

	http.ListenAndServe(s, r)
}

func (r *Router) Get(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.NewRoute().Methods("GET").Path(path).HandlerFunc(f)
}
