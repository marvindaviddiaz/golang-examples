package files

import (
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"bufio"
)

func Main() {

	//f1, err := os.OpenFile("/tmp/test.txt", os.O_APPEND | os.O_RDWR, 0666)

	// Open
	f1, err := os.Create("/tmp/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write
	writeBuffer := bufio.NewWriter(f1)
	for i := 1; i <=5; i++ {
		writeBuffer.WriteString(fmt.Sprintln("Added line", i))
	}
	writeBuffer.Flush()
	f1.Close()

	f1, err = os.Open("/tmp/test.txt")
	b, err := ioutil.ReadAll(f1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)

	// Delete
	os.Remove(f1.Name())

}