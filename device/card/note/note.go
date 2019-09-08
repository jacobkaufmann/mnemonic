package note

import "github.com/jacobkaufmann/mnemonic/device/card"

// A Note represents a map of information that can be transformed into a set
// of cards.
type Note struct {
	Content map[string]interface{} `json:"content"`
	Type    Type                   `json:"type"`
}

// New returns a new note containing content and of type noteType.
func New(content map[string]interface{}, noteType Type) *Note {
	return &Note{content, noteType}
}

// UpdateField updates field in n to hold value.
func (n *Note) UpdateField(field string, value interface{}) {
	n.Content[field] = value
}

// A Type defines a set of fields a note may have and a set of card types to
// define the ways a note can be represented as a card.
type Type struct {
	Name      string          `json:"name"`
	Fields    map[string]bool `json:"fields"`
	CardTypes []*card.Type    `json:"cardTypes"`
}

// A Notebook is a collection of notes.
type Notebook struct {
	Name  string  `json:"name"`
	Notes []*Note `json:"notes"`
}

// NewNotebook returns a new Notebook containing notes.
func NewNotebook(name string, notes []*Note) *Notebook {
	return &Notebook{name, notes}
}
