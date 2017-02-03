package lisp

import (
	"time"
	"bytes"
)

type Entry struct {
	expression string
	date int64
}

func NewEntry(line string) Entry {
	return Entry{expression: line, date: time.Now().Unix()}
}

type Timeline struct {
	entries []Entry
}

func (t *Timeline) NewLine(line string) {
	t.entries = append(t.entries, NewEntry(line)) 
}

func (t *Timeline) LastLine() Entry {
	pos := len(t.entries) - 1
	return t.entries[pos]
}

func (t *Timeline) GetTimeline() []Entry {
	return t.entries
}

func (t *Timeline) toString() string {
	var buffer bytes.Buffer
	for _,entry := range(t.entries){
		buffer.WriteString(entry.expression)
	}

	return buffer.String()
}

func NewTimeline() Timeline {
	return Timeline{entries: make([]Entry, 0)}
}