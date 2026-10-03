package main

func encode(text string) []byte {
	return []byte(text)
}

func decode(tokens []byte) string {
	return string(tokens)
}

type ByteStats struct {
	Counts		[256]int
	TotalBytes	int 
}

type BigramStats struct {
	Counts				[256][256]int
	TransitionTotals 	[256]int 
}

func analyzeBytes(data []byte) ByteStats {
	var stats ByteStats
	stats.TotalBytes = len(data)

	for _, b := range data {
		stats.Counts[b]++
	}

	return stats

}


func analyzeBigrams(data []byte) BigramStats {
	var stats BigramStats
	
	for i := 0; i < len(data) - 1; i++ {
		current := data[i]
		next := data[i+1]

		stats.Counts[current][next]++
		stats.TransitionTotals[current]++
	}

	return stats
}

func (s ByteStats) Probability(b byte) float64 {
	if s.TotalBytes == 0 {
		return 0
	}

	return float64(s.Counts[b]) / float64(s.TotalBytes)
}

func (s BigramStats) Probability(current byte, next byte) float64 {
	if s.TransitionTotals[current] == 0 {
		return 0
	}

	return float64(s.Counts[current][next]) / float64(s.TransitionTotals[current])
	
}