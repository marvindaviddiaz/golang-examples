package main

import (
	"flag"
	"strings"
	"net"
	"bufio"
	"os"
	"fmt"
	"io"
	"log"
	"encoding/base64"
)

func main() {
	op := flag.String("type", "", "Server (s) or client (c) ?")
	address := flag.String("addr", ":8000", "Address? host:port")
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		runUDPServer(*address)
	case "C":
		runUDPClient(*address)
	}
}

func runUDPClient(address string) error {
	conn, err := net.Dial("udp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What message would you like to send?")
	for scanner.Scan() {

		fmt.Println("Writting", scanner.Text())

		bytes := append(scanner.Bytes())
		dst := make([]byte, base64.StdEncoding.EncodedLen(len(bytes)))
		base64.StdEncoding.Encode(dst, bytes)

		conn.Write(append(dst, '\r'))

		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		} else if err == io.EOF {
			log.Println("Connection is closed")
			return nil
		}
		fmt.Println(string(buffer))
		fmt.Println("What message would you like to send?")
	}

	return scanner.Err()
}

func runUDPServer(address string) error {
	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		return err
	}
	defer pc.Close()
	buffer := make([]byte, 1024)
	log.Println("Listening....")
	for {
		n, addr, _ := pc.ReadFrom(buffer)

		dst := make([]byte, base64.StdEncoding.DecodedLen(n))
		base64.StdEncoding.Decode(dst, buffer[:n])

		log.Printf("Address: %s,  Received: %s \n", addr, string(dst))
		_, err := pc.WriteTo([]byte("Message Received"), addr)
		if err != nil {
			log.Fatal("Could not write back on connection", err)
		}
	}
}