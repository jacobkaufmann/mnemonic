package note

import "github.com/jacobkaufmann/mnemonic/device/card"

// A Note represents a structured map of information that can be transformed
// into a defined set of cards.
type Note struct {
	Content map[string]interface{} `json:"content"`
	Type    Type                   `json:"type"`
}

// A Type defines a set of fields a note may have and a set of card types to
// define the ways a note can be represented as a card.
type Type struct {
	Name      string          `json:"name"`
	Fields    map[string]bool `json:"fields"`
	CardTypes []*card.Type    `json:"cardTypes"`
}
