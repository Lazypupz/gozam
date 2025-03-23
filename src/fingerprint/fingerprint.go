package fingerprint

func CreateHash(spec [][]float64) {
	// hash looks like (freqeuncy1,  freq2, time difference) -> song_id)

	sampleRate := 44100 // Sample rate in Hz
	hopSize := 512

	type Song_id struct {
		freq1    float64
		freq2    float64
		time_dif float64
	}

	id := Song_id{}

	freqArray := make([]float64, len(spec))
	for _, val := range spec {
		for j := range val {
			if val[j] == 0 {
				continue
			} else {
				freqArray = append(freqArray, val[j])
			}

		}
	}

	for i := 0; i < len(freqArray); i++ {

		id.freq1 = freqArray[i]
		id.freq2 = freqArray[i+1]
		for time, _ := range spec {
			timeindex := float64(time*hopSize) / float64(sampleRate)
			id.time_dif = timeindex
		}
	}
}
