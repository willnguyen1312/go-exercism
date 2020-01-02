package dna

import (
	"errors"
	"regexp"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	regexDNA := regexp.MustCompile(`A|C|G|T`)

	for _, item := range d {
		if !regexDNA.Match([]byte(string(item))) {
			return nil, errors.New("invalid DNA input")
		}

		if value, found := h[item]; found {
			h[item] = value + 1
		}

	}

	return h, nil
}
