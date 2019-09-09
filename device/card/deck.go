package card

import (
	"math/rand"
	"time"
)

// A Deck is a collection of cards.
type Deck struct {
	Name  string  `json:"name"`
	Cards []*Card `json:"cards"`
}

// NewDeck returns a new deck called name.
func NewDeck(name string) *Deck {
	cards := []*Card{}
	return &Deck{name, cards}
}

// AddCard adds a card to the deck.
func (d *Deck) AddCard(card *Card) {
	d.Cards = append(d.Cards, card)
}

// Clear resets the deck to an empty slice.
func (d *Deck) Clear() {
	d.Cards = []*Card{}
}

// Filter represents a function that determines which cards from a deck
// should be included for study. If the filter returns true, the card
// should be included.
type Filter func(card *Card) (keep bool)

// FilterKeepAll is a filter which returns true for any card.
var FilterKeepAll = func(*Card) bool { return true }

// Study returns a sequence of cards from d for practice. The sequence of cards
// may optionally be filtered according to filter.
func (d *Deck) Study(filter Filter, shuffle bool) []*Card {
	cards := d.Sequence(shuffle)
	if filter == nil {
		return cards
	}

	var filtered []*Card
	for _, c := range cards {
		if filter(c) {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

// Sequence returns a sequence of cards from d. The sequence of cards
// may optionally be shuffled.
func (d *Deck) Sequence(shuffle bool) []*Card {
	if !shuffle {
		return d.Cards
	}
	return shuffleCards(d.Cards)
}

// shuffleCards is a helper function to shuffle a slice of cards using a
// psuedo-random number generator.
func shuffleCards(cards []*Card) []*Card {
	shuffled := make([]*Card, len(cards))
	copy(shuffled, cards)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}
