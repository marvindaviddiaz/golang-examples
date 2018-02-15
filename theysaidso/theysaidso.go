package theysaidso

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func PrintQuoteOfTheDay() {
	response, error := http.Get("http://api.theysaidso.com/qod.json")
	if error != nil {
		fmt.Print(error)
		return
	}
	defer response.Body.Close()
	contents, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Print(error)
		return
	}
	fmt.Print(string(contents))
}