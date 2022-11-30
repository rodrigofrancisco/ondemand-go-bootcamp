package main

import (
	"fmt"
	"net/http"
	router "rodrigofrancisco/go-bootcamp/router"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})

	httpRouter.SERVE(port)

}
