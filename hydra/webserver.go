package hydra

import (
	"net/http"
	"fmt"
	"github.com/marvindaviddiaz/golang-examples/hydra/hlogger"
)

func Main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web service")

	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Hydra software system")

	hlogger.GetInstance().Println("Received an http Get request on root url")
}