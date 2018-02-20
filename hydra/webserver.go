package hydra

import (
	"net/http"
	"fmt"
)

func Main() {
	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Hydra software system")
}