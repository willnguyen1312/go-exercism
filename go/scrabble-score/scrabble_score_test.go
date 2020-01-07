package scrabble

import "testing"

func TestScore(t *testing.T) {
	for _, test := range scrabbleScoreTests {
		if actual := ScoreSwitch(test.input); actual != test.expected {
			t.Errorf("Score(%q) expected %d, Actual %d", test.input, test.expected, actual)
		}
	}
	for _, test := range scrabbleScoreTests {
		if actual := ScoreMap(test.input); actual != test.expected {
			t.Errorf("Score(%q) expected %d, Actual %d", test.input, test.expected, actual)
		}
	}
}

func BenchmarkScoreSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			ScoreSwitch(test.input)
		}
	}
}

func BenchmarkScoreMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			ScoreMap(test.input)
		}
	}
}
