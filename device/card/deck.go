package card

// A Deck is a collection of cards.
type Deck struct {
	Name     string  `json:"name"`
	Cards    []*Card `json:"cards"`
	CardType Type    `json:"cardType"`
}
