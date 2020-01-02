package strand

var rnaTranscriptionLookup = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

// ToRNA function
func ToRNA(dna string) string {
	var result string

	for _, item := range dna {
		result += rnaTranscriptionLookup[string(item)]
	}

	return result
}
