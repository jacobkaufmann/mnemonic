package card

// A Card represents a flashcard with a question and an answer.
type Card struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	History  []bool `json:"history"`
}

// New returns a new card.
func New(q, a string) *Card {
	var history []bool
	return &Card{q, a, history}
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

// Record appends a bool to the history of c indicating whether or not the
// question was answered correctly.
func (c *Card) Record(correct bool) {
	c.History = append(c.History, correct)
}

// ClearHistory resets the performance history of c.
func (c *Card) ClearHistory() {
	c.History = make([]bool, 0)
}

// LastAttempt returns a bool indicating whether c was correctly answered on
// the most recent attempt according to the history.
func (c *Card) LastAttempt() bool {
	if len(c.History) == 0 {
		return false
	}
	return c.History[len(c.History)-1]
}

// NumAttempts returns the number of times an answer has been recorded for c.
func (c *Card) NumAttempts() int {
	return len(c.History)
}
