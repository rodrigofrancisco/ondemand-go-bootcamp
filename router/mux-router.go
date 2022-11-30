package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) SERVE(port string) {
	http.ListenAndServe(port, muxDispatcher)
}

func (*muxRouter) GET(url string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(url, f).Methods("GET")
}

func (*muxRouter) POST(url string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(url, f).Methods("POST")
}
