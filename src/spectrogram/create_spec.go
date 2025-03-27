package spectrogram

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-audio/wav"
	"github.com/r9y9/gossp/stft"
	"github.com/r9y9/gossp/window"
)

func CreateSpectrogram(inputFile string) [][]complex128 {
	test_recording := flag.String("i", "../recording.wav", inputFile)
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
		FrameLen:   1024,
		Window:     window.CreateHanning(1024),
	}
	spectrogram := s.STFT(data)

	return spectrogram
}
