package models

import (
	"errors"
	"time"
)

var ErrorNoRecord = errors.New("models: no matching records found")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
