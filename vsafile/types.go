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

type header struct {
	unknownOne       []byte
	level            string
	options          []byte
	email            string
	eventCount       int
	unknownTwo       []byte
	defaultEventType string
}

func (h header) String() string {
	return "HEADER\n" +
		"        unknownOne: " + hex.EncodeToString(h.unknownOne) + "\n" +
		"             level: " + fmt.Sprint(h.level) + "\n" +
		"           options: " + hex.EncodeToString(h.options) + " (non-ASCII string, hex encoded for display)\n" +
		"             email: " + fmt.Sprint(h.email) + "\n" +
		"       event count: " + fmt.Sprintf("%d", h.eventCount) + "\n" +
		"        unknownTwo: " + hex.EncodeToString(h.unknownTwo) + "\n" +
		"  defaultEventType: " + fmt.Sprint(h.defaultEventType) + "\n"
}

type event struct {
	eventNumber   int
	_type         string
	track         int
	startTime     int
	endTime       int
	startPosition int
	endPosition   int
	unknownThree  []byte
	data          []byte
	continuation  string
}

func (e event) String() string {
	return "      eventNumber: " + fmt.Sprint(e.eventNumber) + "\n" +
		"            _type: " + fmt.Sprint(e._type) + "\n" +
		"            track: " + fmt.Sprint(e.track) + "\n" +
		"       start time: " + fmt.Sprint(e.startTime) + "\n" +
		"         end time: " + fmt.Sprint(e.endTime) + "\n" +
		"    startPosition: " + fmt.Sprint(e.startPosition) + "\n" +
		"      endPosition: " + fmt.Sprint(e.endPosition) + "\n" +
		"     unknownThree: " + hex.EncodeToString(e.unknownThree) + "\n" +
		"             data: " + hex.EncodeToString(e.data) + "\n" +
		"     continuation: " + fmt.Sprint(e.continuation) + "\n"
}

type events []event

func (es events) String() string {
	s := "EVENTS:\n"
	for _, e := range es {
		s += e.String() + "\n"
	}
	return s
}
