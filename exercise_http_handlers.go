package main

import (
	"net/http"
	"fmt"
	"log"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(s))
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf("%v", s))
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	//e := http.ListenAndServe("localhost:4000", nil)
	//if e != nil {
	//	log.Fatal(e)
	//}
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
