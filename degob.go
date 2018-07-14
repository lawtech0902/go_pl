package main

import (
	"os"
	"bufio"
	"encoding/gob"
	"log"
	"fmt"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type Vcard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

var content string
var vc Vcard

func main() {
	file, _ := os.Open("vcard.gob")
	defer file.Close()
	inputReader := bufio.NewReader(file)
	dec := gob.NewDecoder(inputReader)
	err := dec.Decode(&vc)
	if err != nil {
		log.Println("Error in decoding gob")
	}
	fmt.Println(vc)
}
