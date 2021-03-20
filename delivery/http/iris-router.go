package router

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"
)

type irisRouter struct{}

var (
	irisDispatcher = iris.New()
)

func NewIrisRouter() Router {
	return &irisRouter{}
}

type Book struct {
	Title string `json:"title"`
}

func (*irisRouter) GET(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	irisDispatcher.Get(uri, func(ctx iris.Context) {
		books := []Book{
			{"Mastering Concurrency in Go"},
			{"Go Design Patterns"},
			{"Black Hat Go"},
		}

		ctx.JSON(books)
	})
}

func (*irisRouter) POST(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	return
}

func (*irisRouter) SERVE(port string) {
	fmt.Printf("Iris HTTP server running on port %v\n", port)
	irisDispatcher.Listen(port)
}
