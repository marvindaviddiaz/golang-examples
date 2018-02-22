package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type CrewMember struct {
	ID int `json:"id"`
	Name string `json:"name"`
	SecurityClearance int `json:"securityClearance"`
	AccessCodes []string `json:"accessCodes"`
	dummy1 int  // ignored because of un-exportable
	Dummy2 int `json:"-"` //ignored because of "-"
}

func Main() {

	/** MARSHALL */

	cm := CrewMember{1, "Jaro", 10, []string{"ADA", "LAL"}, 0, 0}
	b, err := json.Marshal(&cm) // Marshall supports values and references(pointers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("MARSHAL")
	fmt.Println(string(b))

	// ENCODERS = To write the result into a file
	f, err := os.Create("/tmp/test.json")
	defer f.Close()
	json.NewEncoder(f).Encode(cm)
	fmt.Println("Json enconded into", f.Name())

	/** UN-MARSHALL */

	bytes := []byte(`{"id":1,"name":"Jaro","securityClearance":10,"accessCodes":["ADA","LAL"]}`)
	member := new(CrewMember)
	json.Unmarshal(bytes, member)
	fmt.Println("UNMARSHAL")
	fmt.Println(*member)

	// DECODERS = To read json from a file
	f, err = os.Open("/tmp/test.json")
	defer f.Close()
	json.NewDecoder(f).Decode(member)
	fmt.Println("Json deconded from", f.Name())
	fmt.Println(*member)
}
