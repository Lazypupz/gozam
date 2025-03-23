package main

import (
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
	wav.GetWavData("recording.wav")

}
