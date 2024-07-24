package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type julienschmidtRouter struct{}

var (
	julienschmidtDispatcher = httprouter.New()
)

func NewJulienschmidtRouter() Router {
	return &julienschmidtRouter{}
}

func (*julienschmidtRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	julienschmidtDispatcher.HandlerFunc("GET", uri, f)
}
func (*julienschmidtRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	julienschmidtDispatcher.HandlerFunc("POST", uri, f)
}
func (*julienschmidtRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	julienschmidtDispatcher.HandlerFunc("PUT", uri, f)
}
func (*julienschmidtRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	julienschmidtDispatcher.HandlerFunc("DELETE", uri, f)
}
func (*julienschmidtRouter) SERVE(port string) {
	fmt.Printf("Julien Schmidt HTTP server running on port %v", port)
	http.ListenAndServe(port, julienschmidtDispatcher)
}
