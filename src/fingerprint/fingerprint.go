package fingerprint

type keyFreq struct {
	peakFreq   float64
	targetFreq float64
}

func createFreqArray(spec [][]float64) []float64 {
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
	return freqArray
}

func CreateHash(freqArray []float64) map[keyFreq]float64 {
	// hash looks like (freqeuncy1,  freq2, time difference) -> song_id)

	sampleRate := 44100 // Sample rate in Hz
	hopSize := 1024

	for i := 0; i < len(freqArray); i++ {
		time_dif := 0.0
		freq1 := freqArray[i]
		freq2 := freqArray[i+1]
		for time, _ := range freqArray {
			timeindex := float64(time*hopSize) / float64(sampleRate)
			time_dif = timeindex

		}
		hashTable := make(map[keyFreq]float64)
		hashTable[keyFreq{peakFreq: freq1, targetFreq: freq2}] = time_dif
		return hashTable
	}

}
