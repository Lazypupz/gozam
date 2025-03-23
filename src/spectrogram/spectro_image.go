package spectrogram

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func SpectrogramToImg(spectrogram [][]complex128, outputPath string) ([][]float64, error) {

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
	realSpec := convert_spec_to_float(spectrogram)

	peaks := make([][]float64, len(spectrogram))
	for i, row := range spectrogram {
		peaks[i] = make([]float64, len(row))

	}

	for i := 1; i < numWindows-1; i++ {
		for j := 1; j < freqBins-1; j++ {
			if isPeak(realSpec, i, j) {
				peaks[i][j] = realSpec[i][j]
			}
			if !isPeak(realSpec, i, j) {
				peaks[i][j] = math.NaN()
			}

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
	return peaks, nil
}
