// Package vsafile provides functionality for working with VSA files.
package vsafile

import (
	"encoding/hex"
	"fmt"
)

// File VSA file
type File struct {
	header header
	events events
}

func (f File) String() string {
	return f.header.String() +
		"\n" +
		f.events.String()
}

// private

type unknownOne []byte
type level string
type options []byte
type email string
type eventCount int
type unknownTwo []byte
type firstEventType string

type header struct {
	unknownOne     unknownOne
	level          level
	options        options
	email          email
	eventCount     eventCount
	unknownTwo     unknownTwo
	firstEventType firstEventType
}

func (h header) String() string {
	return "HEADER\n" +
		"      unknownOne: " + hex.EncodeToString([]byte(h.unknownOne)) + "\n" +
		"           level: " + string(h.level) + "\n" +
		"         options: " + hex.EncodeToString([]byte(h.options)) + " (non-ASCII string, hex encoded for display)\n" +
		"           email: " + string(h.email) + "\n" +
		"     event count: " + fmt.Sprintf("%d", h.eventCount) + "\n" +
		"      unknownTwo: " + hex.EncodeToString([]byte(h.unknownTwo)) + "\n" +
		"  firstEventType: " + string(h.firstEventType) + "\n"
}

type event struct {
	number        int
	_type         string
	track         int
	startTime     int
	endTime       int
	startPosition int
	endPosition   int
	data          []byte
}

func (e event) String() string {
	return fmt.Sprintf("  event: %d\n", e.number) +
		"            _type: " + e._type + "\n" +
		"            track: " + fmt.Sprint(e.track) + "\n" +
		"       start time: " + fmt.Sprint(e.startTime) + "\n" +
		"         end time: " + fmt.Sprint(e.endTime) + "\n" +
		"    startPosition: " + fmt.Sprint(e.startPosition) + "\n" +
		"      endPosition: " + fmt.Sprint(e.endPosition) + "\n" +
		"             data: " + hex.EncodeToString(e.data) + "\n"
}

type events []event

func (es events) String() string {
	s := "EVENTS:\n"
	for _, e := range es {
		s += e.String()
	}
	return s
}
