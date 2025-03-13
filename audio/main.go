package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mp := AnchorMap()
	fmt.Println(mp)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
