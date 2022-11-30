package router

import "net/http"

type Router interface {
	SERVE(port string)
	GET(url string, f func(w http.ResponseWriter, r *http.Request))
	POST(url string, f func(w http.ResponseWriter, r *http.Request))
}
