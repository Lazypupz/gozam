package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-audio/wav"
)

func wavSpec() {
	f, err := os.Open("../wav/bib.wav")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	d := wav.NewDecoder(f)
	if !d.IsValidFile() {
		log.Fatal("get scammed buddy")
	}
	fmt.Println(d.SampleRate)
	fmt.Println(d.SampleBitDepth())

}

func main() {
	wavSpec()
}
