package webserver

import (
	"net/http"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Main() {
	log.Println("Starting Hydra web service")

	/* SIMPLE */
	//http.HandleFunc("/", rootHandler)

	/* MORE CONTROL */
	http.ListenAndServe(":8080", newHandler())

	/* MORE CONTROL */

	/*
	server := & http.Server {
		Addr: ":8080",
		Handler: newHandler(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	server.ListenAndServe()
	*/
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received an http Get request on root url")

	fmt.Fprint(w, "Welcome to the Hydra software system")
}


type testHandler struct {
	r int
}
func newHandler() testHandler {
	rand.Seed(time.Now().UnixNano())
	return testHandler {
		r: rand.Intn(200),
	}
}
func (h testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch  r.URL.Path {
	case "/":
		fmt.Fprint(w, "Welcome to the root URL from a custom Handler")
	case "/test":
		fmt.Fprint(w, "Welcome to the /test/rand URL with random number: ", h.r)
	}
	fmt.Println(r.URL.Query())
}