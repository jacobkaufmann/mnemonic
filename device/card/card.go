package card

import (
	"fmt"
	"strings"
	"time"
)

// A Card represents a flashcard with a question and an answer.
type Card struct {
	Question string    `json:"question"`
	Answer   string    `json:"answer"`
	History  []*Record `json:"history"`
	Created  time.Time `json:"created"`
}

// New returns a new card.
func New(q, a string) *Card {
	history := []*Record{}
	return &Card{q, a, history, time.Now().UTC()}
}

// NewFromMap returns a new card defined by map m and type t.
func NewFromMap(m map[string]interface{}, t Type) *Card {
	var q, a strings.Builder
	for k, v := range m {
		if t.Question[k] {
			q.WriteString(fmt.Sprintf("%s: %v\n", k, v))
		}
		if t.Answer[k] {
			a.WriteString(fmt.Sprintf("%s: %v\n", k, v))
		}
	}

	return New(q.String(), a.String())
}

// A Template defines a set of fields to display.
type Template map[string]bool

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

// A Record represents a historical record of a response to a card. The record
// indicates whether the card was correctly answered and the time of the
// response.
type Record struct {
	Correct bool      `json:"correct"`
	Created time.Time `json:"created"`
}

// NewRecord returns a new record.
func NewRecord(correct bool) *Record {
	return &Record{correct, time.Now().UTC()}
}

// Query returns the question and answer for c.
func (c *Card) Query() (q, a string) {
	return c.Question, c.Answer
}

// AddRecord appends a new record to the history of c indicating whether or not
// the card was answered correctly.
func (c *Card) AddRecord(correct bool) {
	c.History = append(c.History, NewRecord(correct))
}

// ClearHistory resets the performance history of c.
func (c *Card) ClearHistory() {
	c.History = []*Record{}
}

// LastAttempt returns a bool indicating whether c was correctly answered on
// the most recent attempt according to the history.
func (c *Card) LastAttempt() *Record {
	if len(c.History) == 0 {
		return nil
	}
	return c.History[len(c.History)-1]
}

// NumAttempts returns the number of times an answer has been recorded for c.
func (c *Card) NumAttempts() int {
	return len(c.History)
}
