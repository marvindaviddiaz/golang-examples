package xml

import (
	"fmt"
	"os"
	"encoding/xml"
)

type CrewMember struct  {
	XMLName xml.Name `xml:"root"`
	ID int `xml:"id,omitempty"`
	Name string `xml:"name,attr"`
	SecurityClearance int `xml:"securityClearance"`
	AccessCodes []string `xml:"accessCodes>code"`
	dummy1 int  // ignored because of un-exportable
	Dummy2 int `xml:"-"` //ignored because of "-"
}

func Main() {

	/** MARSHALL */

	cm := CrewMember{ID:1, Name:"Jaro", SecurityClearance:10, AccessCodes:[]string{"ADA", "LAL"}}
	//b, err := xml.Marshal(&cm) // Marshall supports values and references(pointers)
	b, err := xml.MarshalIndent(&cm, "", "	")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("MARSHAL")
	fmt.Println(xml.Header, string(b))

	// ENCODERS = To write the result into a file
	f, err := os.Create("/tmp/test.xml")
	defer f.Close()
	xml.NewEncoder(f).Encode(cm)
	fmt.Println("XML enconded into", f.Name())

	/** UN-MARSHALL */

	bytes := []byte(`<?xml version="1.0" encoding="UTF-8"?>
 						<root name="Jaro">
							<id>1</id>
							<securityClearance>10</securityClearance>
							<accessCodes>
								<code>ADA</code>
								<code>LAL</code>
							</accessCodes>
						</root>`)
	member := new(CrewMember)
	xml.Unmarshal(bytes, member)
	fmt.Println("UNMARSHAL")
	fmt.Println(*member)

	// DECODERS = To read json from a file
	f, err = os.Open("/tmp/test.xml")
	defer f.Close()
	member2 := new(CrewMember)
	xml.NewDecoder(f).Decode(member2)
	fmt.Println("XML decoded from", f.Name())
	fmt.Println(*member2)
}
