package main

import (
	"math"
)

const (
	maxFreq    = -18.2
	targetZone = 10 // offset in the arary idek? ?? ? ? ??
)

func MaxFilter(spec [][]float64) []float64 {
	anc := make([]float64, 1024)

	for _, vec := range spec {

		for _, val := range vec[:1024] {

			newVal := math.Log(val)

			if newVal < maxFreq {
				anc = append(anc, newVal) // fix this and the whole thing

			}

		}

	}
	return anc
}

func AnchorMap() map[int]float64 {
	newAnc := cliTargetZone()
	ancMap := make(map[int]float64)
	for i, val := range newAnc {
		ancMap[i] = val
	}
	return ancMap

}
