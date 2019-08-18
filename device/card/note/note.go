package note

import "github.com/jacobkaufmann/mnemonic/device/card"

// A Note represents a set of fields. Fields are units of information, and
// the note type defines how those fields are organized.
type Note struct {
	Fields map[string]interface{} `json:"fields"`
	Type   Type                   `json:"type"`
}

// A Type defines a set of fields a note may have and a set of card types
// which define the ways a note can be represented as a card.
type Type struct {
	Name       string       `json:"name"`
	FieldNames []string     `json:"fieldNames"`
	CardTypes  []*card.Type `json:"cardTypes"`
}
