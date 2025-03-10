package shazam

import (
	"math"
)

const (
	maxFreq    = -8.50
	targetZone = 5 // offset in the arary idek? ?? ? ? ??
)

func MaxFilter(spec [][]float64) ([]float64, float64, float64) {
	anc := make([]float64, 1024)
	freqntime := make([][]int, len(spec))

	for i, vec := range spec {
		time := make([]int, len(spec))
		time = append(time, i)
		freqntime = append(freqntime, time)

		for j, val := range vec[:1024] {
			freq := make([]int, len(spec))
			freq = append()
			newVal := math.Log(val)

			if newVal > maxFreq {
				anc = append(anc, newVal) // fix this and the whole thing
			}

		}

	}
	return anc
}

// need to implement target zone before the anchorMap function

func anchorMap(anc []float64) map[int]float64 {
	ancMap := make(map[int]float64)
	//ancMap
}
