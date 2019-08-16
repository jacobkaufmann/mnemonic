package card

import (
	"math/rand"
	"time"
)

// gen is a pseudo-random number generator.
var gen = rand.New(rand.NewSource(time.Now().UnixNano()))

// A Card represents a flashcard. If a card is reversible, either side may
// be presented as a prompt. Otherwise, front is always treated as the question
// and back is always treated as the answer.
type Card struct {
	Front      string `json:"front"`
	Back       string `json:"back"`
	Reversible bool   `json:"reversible"`
}

// New returns a new card.
func New(front, back string, reversible bool) *Card {
	return &Card{
		Front:      front,
		Back:       back,
		Reversible: reversible,
	}
}

// Query returns a question and answer for c. If reversible, return either
// combination with probability one half.
func (c *Card) Query() (q, a string) {
	if !c.Reversible {
		return c.Front, c.Back
	}

	n := gen.Float64()
	if n < 0.5 {
		return c.Front, c.Back
	}
	return c.Back, c.Front
}
