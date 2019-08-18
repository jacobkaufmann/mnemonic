package card

// A Card represents a flashcard with a question and an answer.
type Card struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// New returns a new card.
func New(q, a string) *Card {
	return &Card{q, a}
}

// A Type contains templates for mapping fields to the question and answer
// fields of a Card.
type Type struct {
	Name     string   `json:"name"`
	Question Template `json:"questionTemplate"`
	Answer   Template `json:"answerTemplate"`
}

// NewType returns a new card type.
func NewType(name string, q, a Template) *Type {
	return &Type{name, q, a}
}

// A Template is a mapping which determines a set of fields to display.
type Template map[string]interface{}

// Query returns the question and answer for c.
func (c *Card) Query() (q, a string) {
	return c.Question, c.Answer
}
