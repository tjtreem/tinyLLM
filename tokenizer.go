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

func analyzeBytes(data []byte) ByteStats {
	var stats ByteStats
	stats.TotalBytes = len(data)

	for _, b := range data {
		stats.Counts[b]++
	}

	return stats

}

func (s ByteStats) Probability(b byte) float64 {
	if s.TotalBytes == 0 {
		return 0
	}

	return float64(s.Counts[b]) / float64(s.TotalBytes)
}