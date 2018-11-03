package internal

import (
	"github.com/gorilla/mux"
	"github.com/jvgrootveld/knowledge-base/internal/api"
	"net/http"
)

const KB_API = "/kbapi"

// Initializes a router, see api package for handlers
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Static (frontend) files
	api.InitStaticApi(r)

	// Knowledge-base api
	kbapi := r.PathPrefix(KB_API).Subrouter()

	api.InitDirectoryApi(kbapi)
	api.InitFileApi(kbapi)

	http.Handle("/", r)

	return r
}
