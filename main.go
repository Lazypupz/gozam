package main

import (
	"fmt"

	"github.com/Lazypupz/gozam/src/fingerprint"
	"github.com/Lazypupz/gozam/src/spectrogram"
	"github.com/Lazypupz/gozam/src/wav"
)

func main() {
	/*spec := spectrogram.CreateSpec()
	outputPath := "spectrogram.png"
	peaks, err := spectrogram.SpectrogramToImg(spec, outputPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("spectrogram image saved to", outputPath)

	for i, val := range peaks {
		for j := range val {
			fmt.Println(peaks[i][j])
		}
	}
	*/

	//wav.Record("recording.wav")
	audio_file := "record.wav"
	output := wav.GetWavData(audio_file)
	wav.ReformatWav(audio_file, uint(output))
	spec := spectrogram.CreateSpec(audio_file)
	peaks, err := spectrogram.SpectrogramToImg(spec, "image.png")
	if err != nil {
		fmt.Print("Error", err)
	}
	fingerprint.CreateHash(peaks)
	//need to fix createhash

}
