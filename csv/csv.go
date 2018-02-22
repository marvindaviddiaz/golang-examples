package csv

import (
	"os"
	"log"
	"encoding/csv"
	"io"
	"fmt"
)

func Main() {
	write()
	read()
}

func write() {
	records := [][] string {
		{"Jaro", "5", "ALA,OIO"},
		{"Hala", "4", "A8D,B0O"},
		{"Kay", "3", "H8J,D3N"},
	}
	file, err := os.Create("/tmp/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	//w.Comma = ';'
	w.WriteAll(records)
	w.Flush()

	fmt.Println("File writen:", file.Name())
}


func read() {
	file, err := os.Open("/tmp/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comment = '#'
	//r.Comma = ';'

	fmt.Println("Reading from:", file.Name())

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			if pe, ok := err.(*csv.ParseError); ok {
				fmt.Println("bad column:", pe.Column)
				fmt.Println("bad line:", pe.Line)
				fmt.Println("Error reported:", pe.Err)
				if pe.Err == csv.ErrFieldCount {
					continue
				}
			}
			log.Fatal(err)
		}
		fmt.Println("CSV Row:", record)
	}
}
