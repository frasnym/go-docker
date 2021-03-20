package main

import (
	"fmt"
	"net/http"
	"os"

	router "github.com/frasnym/go-docker/delivery/http"
)

var (
	httpRouter router.Router = router.NewIrisRouter()
)

func main() {
	httpRouter.GET("/", func(rw http.ResponseWriter, r *http.Request) {})

	port, present := os.LookupEnv("PORT")
	if !present {
		port = "8000"
	}
	httpRouter.SERVE(fmt.Sprintf(":%s", port))
}
