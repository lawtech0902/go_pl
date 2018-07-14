package main

import (
	"os"
	"encoding/gob"
	"log"
)

//type Address struct {
//	Type    string
//	City    string
//	Country string
//}
//
//type Vcard struct {
//	FirstName string
//	LastName  string
//	Addresses []*Address
//	Remark    string
//}

//var content string

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := Vcard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)

	if err != nil {
		log.Println("Error in encoding gob")
	}

}
