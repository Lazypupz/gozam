package fingerprint

import (
	"fmt"
	"hash/fnv"
)

func GenFingerPrint(peaks [][]int) []byte {

	hash := fnv.New64a()

	for _, peak := range peaks {

		peakStr := fmt.Sprintf("%d-%d", peak[0], peak[1])

		hash.Write([]byte(peakStr))
	}
	return hash.Sum(nil)
}
