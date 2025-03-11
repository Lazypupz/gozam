package shazam

func genHash()

/*	rand.Shuffle(6, func([i]let int, [j]let int) {
		let := []string{
			"a", "b", "c", "d", "e", "f",
		}
	})

}*/

func cliTargetZone() []float64 {
	//hash := make(map[string]float64)
	newAnc := make([]float64, 512)
	spec := createSpec()
	anc := MaxFilter(spec)
	for _, i := range anc {
		result := i - i + 1
		newAnc = append(newAnc, result)
	}
	return newAnc

}
