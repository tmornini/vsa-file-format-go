// Package vsafile provides functionality for working with VSA files.
package vsafile

import (
	"encoding/hex"
	"fmt"
	"time"
)

// File VSA file
type File struct {
	header         header
	events         events
	parseStartTime time.Time
	parseEndTime   time.Time
}

// EventsPerSecond returns parsing duration and rate
func (f File) EventsPerSecond() (time.Duration, float32) {
	duration := f.parseEndTime.Sub(f.parseStartTime)
	eventsPerSecond := float32(f.header.eventCount) / float32(duration.Seconds())
	return duration, eventsPerSecond
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
	otherEventType   string
}

func (h header) String() string {
	return "HEADER\n" +
		"        unknownOne: " + hex.EncodeToString(h.unknownOne) + "\n" +
		"             level: " + fmt.Sprint(h.level) + "\n" +
		"           options: " + hex.EncodeToString(h.options) + " (non-ASCII string, hex encoded for display)\n" +
		"             email: " + fmt.Sprint(h.email) + "\n" +
		"       event count: " + fmt.Sprintf("%d", h.eventCount) + "\n" +
		"        unknownTwo: " + hex.EncodeToString(h.unknownTwo) + "\n" +
		"  defaultEventType: " + fmt.Sprint(h.defaultEventType) + "\n" +
		"  otherEventType: " + fmt.Sprint(h.otherEventType) + "\n"
}

type event struct {
	eventNumber   int
	_type         string
	track         int
	startTime     int
	endTime       int
	startPosition int
	endPosition   int
	text          string
	data          []byte
	unknownFour   []byte
	continuation  string
}

func (e event) String() string {
	return "            eventNumber: " + fmt.Sprint(e.eventNumber) + "\n" +
		"                  _type: " + e._type + "\n" +
		"                  track: " + fmt.Sprint(e.track) + "\n" +
		"             start time: " + fmt.Sprint(e.startTime) + "\n" +
		"               end time: " + fmt.Sprint(e.endTime) + "\n" +
		"          startPosition: " + fmt.Sprint(e.startPosition) + "\n" +
		"            endPosition: " + fmt.Sprint(e.endPosition) + "\n" +
		"                   text: " + e.text + "\n" +
		"                   data: " + hex.EncodeToString(e.data) + "\n" +
		"           continuation: " + fmt.Sprint(e.continuation) + "\n"
}

type events []event

func (es events) String() string {
	s := "EVENTS:\n"
	for _, e := range es {
		s += e.String() + "\n"
	}
	return s
}

type audioFile struct {
}
type audioFiles []audioFile

type videoFile struct{}
type videoFiles []videoFile

type unknnownSix struct{}

type trackSetting struct{}
type trackSettings []trackSetting
