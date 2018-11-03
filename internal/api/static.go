package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitStaticApi(r *mux.Router) {
	r.Handle("/", http.FileServer(http.Dir("./static")))

	cssHandler := http.FileServer(http.Dir("./static/css"))
	jsHandler := http.FileServer(http.Dir("./static/js"))
	imagesHandler := http.FileServer(http.Dir("./static/images"))
	vendorHandler := http.FileServer(http.Dir("./static/vendor"))

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", cssHandler))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", jsHandler))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", imagesHandler))
	r.PathPrefix("/vendor/").Handler(http.StripPrefix("/vendor/", vendorHandler))
}
