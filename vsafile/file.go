// Package vsafile provides functionality for working with VSA files.
package vsafile

import (
	"encoding/hex"
	"fmt"
)

type unknownOne []byte

func (v unknownOne) String() string {
	return "unknownOne:     " + hex.EncodeToString([]byte(v)) + "\n"
}

type level string

func (l level) String() string {
	return "level:       " + string(l) + "\n"
}

type options []byte

func (o options) String() string {
	return "options:     " + hex.EncodeToString([]byte(o)) + "\n"
}

type email string

func (e email) String() string {
	return "email:       " + string(e) + "\n"
}

type eventCount uint32

func (ec eventCount) String() string {
	return fmt.Sprintf("event count: %d\n", ec)
}

type unknownTwo []byte

func (o unknownTwo) String() string {
	return "unknownTwo:       " + hex.EncodeToString([]byte(o)) + "\n"
}

type header struct {
	unknownOne unknownOne
	level      level
	options    options
	email      email
	eventCount eventCount
	unknownTwo unknownTwo
}

func (h header) String() string {
	return h.unknownOne.String() +
		h.level.String() +
		h.options.String() +
		h.email.String() +
		h.eventCount.String() +
		h.unknownTwo.String() + "\n"
}

type event struct {
	kind          string
	track         int16
	startTime     int32
	endTime       int32
	startPosition int32
	endPosition   int32
	text          string
	data          []byte
}

func (e event) String() string {
	return "event: ?\n"
}

type events []event

func (es events) String() string {
	s := "events:\n"
	for _, e := range es {
		s += e.String()
	}
	return s
}

// File VSA file
type File struct {
	header header
	events events
}

func (f File) String() string {
	return f.header.String() + f.events.String()
}
