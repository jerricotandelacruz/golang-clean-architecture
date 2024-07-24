package main

import (
	"fmt"
	"net/http"

	router "github.com/jerricotandelacruz/golang-clean-architecture/http"
)

var (
	httpRouter router.Router = router.NewJulienschmidtRouter()
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", Sample)

	httpRouter.SERVE(port)
}

func Sample(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
