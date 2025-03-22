package spectrogram

func convert_spec_to_float(spectrogram [][]complex128) [][]float64 {

	convertedSpectrogram := make([][]float64, len(spectrogram))
	for i, row := range spectrogram {
		convertedSpectrogram[i] = make([]float64, len(row))
		for j, col := range row {
			convertedSpectrogram[i][j] = real(col)*real(col) + imag(col)*imag(col)
		}
	}
	return convertedSpectrogram
}

func isPeak(data [][]float64, i, j int) bool {
	// get number of rows and columns
	rows := len(data)
	cols := len(data[0])

	//the current value
	val := data[i][j]

	//check neighbors (left, right, up, down, diagonal)

	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}

			ni, nj := i+di, j+dj

			//check if within bounds
			if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
				if val <= data[ni][nj] {
					// if any neighbour is larger ... not a peak
					return false
				}
			}
		}
	}

	// If no neighbors are larger ...  it's a peak
	return true
}
