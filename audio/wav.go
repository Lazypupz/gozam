package main

import (
	"flag"
	"fmt"
	"log"
	"math"

	"github.com/r9y9/gossp"
	"github.com/r9y9/gossp/io"
	"github.com/r9y9/gossp/stft"
	"github.com/r9y9/gossp/window"
)

/*
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
*/
func createSpec() {
	test_recording := flag.String("i", "../wav/fixed.wav", "fixed.wav")
	flag.Parse()

	w, err := io.ReadWav(*test_recording)
	if err != nil {
		log.Fatal(err)
	}

	data := w.GetMonoData()
	s := &stft.STFT{
		FrameShift: int(float64(w.SampleRate) / 100.0),
		FrameLen:   2048,
		Window:     window.CreateHanning(2048),
	}
	spectrogram, _ := gossp.SplitSpectrogram(s.STFT(data))
	PrintMatrixAsGnuplotFormat(spectrogram)
}

func PrintMatrixAsGnuplotFormat(matrix [][]float64) {
	fmt.Println("#", len(matrix[0]), len(matrix)/2)
	for i, vec := range matrix {
		for j, val := range vec[:1024] {
			fmt.Println(i, j, math.Log(val))
		}
		fmt.Println("")
	}
}

func tft() {
	t := flag.String("i", "../wav/Recording.wav", "Recording")
	flag.Parse()

	fmt.Println("FilePath:", *t)
	//test file path
}

func main() {
	createSpec()
}
