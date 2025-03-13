package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"math"
	"math/cmplx"
	"os"

	"github.com/go-audio/wav"
	"github.com/r9y9/gossp/stft"
	"github.com/r9y9/gossp/window"
)

func createSpec() [][]float64 {
	test_recording := flag.String("i", "../wav/fixed_mono.wav", "fixed_mono.wav")
	flag.Parse()

	file, err := os.Open(*test_recording)
	if err != nil {
		log.Fatal("Error opening WAV file:", err)
	}
	defer file.Close()

	// fuck gossp, os built diff( ithink)
	decoder := wav.NewDecoder(file)
	if !decoder.IsValidFile() {
		log.Fatal("Invalid Wav file")
	}

	buf, err := decoder.FullPCMBuffer()
	if err != nil {
		log.Fatal("error decoding WAv data:", err)
	}

	// Confirm we have audio data
	if len(buf.Data) == 0 {
		log.Fatal("error: Wav file does not contain data")
	}

	fmt.Printf("Sample Rate: %d, Channels: %d, Bit Depth: %d\n",
		decoder.SampleRate, decoder.NumChans, decoder.BitDepth)

	bitdepth := decoder.BitDepth
	if bitdepth < 1 || bitdepth > 32 {
		log.Fatal("nuh uh:", bitdepth)
	}
	data := make([]float64, len(buf.Data))
	for i, sample := range buf.Data {
		data[i] = float64(sample) / float64(int(1)<<(bitdepth-1))
	}

	s := &stft.STFT{
		FrameShift: int(float64(decoder.SampleRate) / 100.0),
		FrameLen:   2048,
		Window:     window.CreateHanning(2048),
	}
	spectrogram := s.STFT(data)
	convertedSpectrogram := make([][]float64, len(spectrogram))
	for i, row := range spectrogram {
		convertedSpectrogram[i] = make([]float64, len(row))
		for j, val := range row {
			// Use the magnitude of the complex number
			convertedSpectrogram[i][j] = real(val)*real(val) + imag(val)*imag(val) // Magnitude squared (can use sqrt() for actual magnitude)
		}
	}

	PrintMatrixAsGnuplotFormat(convertedSpectrogram) // need to fix this shit why am i stuck is this hard
	//nvm gpt came in clutch :)
	return convertedSpectrogram
}

func spectrogramToImg(spectrogram [][]complex128, outputPath string) error {

	numWindows := len(spectrogram)
	freqBins := len(spectrogram[0])
	img := image.NewGray(image.Rect(0, 0, freqBins, numWindows))

	maxMagnitude := 0.0
	for i := 0; i < numWindows; i++ {
		for j := 0; j < numWindows; j++ {
			magnitude := cmplx.Abs(spectrogram[i][j])
			if magnitude > maxMagnitude {
				maxMagnitude = magnitude
			}
		}
	}
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
