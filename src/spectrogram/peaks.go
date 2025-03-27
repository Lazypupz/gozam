package spectrogram

import (
	"math/cmplx"
)

func ExtractPeaks(spectrogram [][]complex128) [][]int {
	var peaks [][]int

	for i := 1; i < len(spectrogram); i++ {
		for j := 1; j < len(spectrogram[i]); j++ {
			magnitude := cmplx.Abs(spectrogram[i][j])
			if isPeak(spectrogram, i, j, magnitude) {
				peaks = append(peaks, []int{i, j})
			}
		}
	}
	return peaks
}

func isPeak(spectrogram [][]complex128, i, j int, magnitude float64) bool {
	// get number of rows and columns

	rows := len(spectrogram)
	cols := len(spectrogram[0])

	//the current value

	//check neighbors (left, right, up, down, diagonal)

	for di := i - 1; di <= i+1; di++ {
		for dj := j - 1; dj <= j+1; dj++ {
			//check if within bounds
			if di >= 0 && di < rows && dj >= 0 && dj < cols {
				if di == i && dj == j {
					continue
				}

				neightbourMagnitude := cmplx.Abs(spectrogram[i][j])
				if magnitude < neightbourMagnitude {
					return false
				}
			}

		}
	}

	// If no neighbors are larger ...  it's a peak
	return true
}
