package spectrogram

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func SpectrogramToImg(spectrogram [][]complex128, outputPath string) ([][]complex128, error) {

	numWindows := len(spectrogram)
	freqBins := len(spectrogram[0])
	img := image.NewGray(image.Rect(0, 0, freqBins, numWindows))

	maxMagnitude := 0.0
	for i := 0; i < numWindows; i++ {
		for j := 0; j < freqBins; j++ {
			magnitude := cmplx.Abs(spectrogram[i][j])
			if magnitude > maxMagnitude {
				maxMagnitude = magnitude
			}
		}
	}

	for i := 0; i < numWindows; i++ {
		for j := 0; j < freqBins; j++ {
			magnitude := cmplx.Abs(spectrogram[i][j])
			intensity := uint8(math.Floor(255 * (magnitude / maxMagnitude)))
			img.SetGray(j, i, color.Gray{Y: intensity})

		}
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return nil, err

	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return nil, err

	}
	return spectrogram, nil
}
