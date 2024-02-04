package vsafile

import (
	"encoding/binary"
	"fmt"
	"io"
)

func bytesFrom(reader io.Reader, length int) (*[]byte, error) {
	bytes := make([]byte, length)

	bytesRead, err := io.ReadFull(reader, bytes)
	if err != nil {
		return nil, err
	}
	if bytesRead != length {
		return nil, fmt.Errorf("expected to read %d byte(s), read %d", length, bytesRead)
	}

	return &bytes, nil
}

func stringFrom(reader io.Reader) (*string, error) {
	lengthBytes, err := bytesFrom(reader, 1)

	if err != nil {
		return nil, err
	}

	bytes := *lengthBytes

	stringLength := int(bytes[0])

	stringBytes, err := bytesFrom(reader, stringLength)
	if err != nil {
		return nil, err
	}

	s := string(*stringBytes)

	return &s, nil
}

func unknownOneFrom(reader io.Reader) (*unknownOne, error) {
	bytesRead, err := bytesFrom(reader, 12)
	if err != nil {
		return nil, err
	}

	return (*unknownOne)(bytesRead), nil
}

func levelFrom(reader io.Reader) (*level, error) {
	l, err := stringFrom(reader)
	if err != nil {
		return nil, err
	}

	return (*level)(l), nil
}

func optionsFrom(reader io.Reader) (*options, error) {
	o, err := stringFrom(reader)
	if err != nil {
		return nil, err
	}

	os := options(*o)

	return &os, nil
}

func emailFrom(reader io.Reader) (*email, error) {
	es, err := stringFrom(reader)

	if err != nil {
		return nil, err
	}

	return (*email)(es), nil
}

func eventCountFrom(reader io.Reader) (*eventCount, error) {
	ecBytes, err := bytesFrom(reader, 4)
	if err != nil {
		return nil, err
	}

	c := binary.LittleEndian.Uint32(*ecBytes)

	ec := eventCount(c)

	return &ec, nil
}

func unknownTwoFrom(reader io.Reader) (*unknownTwo, error) {
	bytesRead, err := bytesFrom(reader, 4)

	if err != nil {
		return nil, err
	}

	return (*unknownTwo)(bytesRead), nil
}

func headerFrom(reader io.Reader) (*header, error) {
	h := header{}

	v, err := unknownOneFrom(reader)
	if err != nil {
		return nil, err
	}
	h.unknownOne = *v

	l, err := levelFrom(reader)
	if err != nil {
		return nil, err
	}
	h.level = *l

	os, err := optionsFrom(reader)
	if err != nil {
		return nil, err
	}
	h.options = *os

	e, err := emailFrom(reader)
	if err != nil {
		return nil, err
	}
	h.email = *e

	ec, err := eventCountFrom(reader)
	if err != nil {
		return nil, err
	}
	h.eventCount = *ec

	ot, err := unknownTwoFrom(reader)
	if err != nil {
		return nil, err
	}
	h.unknownTwo = *ot

	return &h, nil
}

// EVENTS
//     SIZE(2) DEFAULTEVENTTYPE(S)
//     SIZE(2) COUNT(S)
//         EVENT
//             TRACK(2)
//             STARTTIME(4)
//             ENDTIME(4)
//             STARPOSITION(4)
//             ENDPOSITION(4)
//             SIZE(1) TEXT(S)            <--\
//                 "CEventBarLinear"         |
//                     DATA(12)              | <-\
//                 "CEventBarPulse"          |   | or
//                     DATA(16) Pulse        | <-|
//             NEXTEVENTTYPE(2)              |   |
//                 FF FF = New Event Type >--/   |
//                 01 80 = Default event type >--/
//                 30 87 = Other event type ? perhaps observed but not understood ?
//                 00 00 = Last event

// AUDIO FILES
//     COUNT(4)
//         Audio Files
//             SIZE(1) FILE(S)
//             SIZE(1) AUDIO_DEVICE(S)

// VIDEO FILES
//     COUNT(4)
//         Video Files
//             SIZE(1) FILE(S)
//             SIZE(1) AUDIO(S)
//             SIZE(1) MONITOR(S)
//             FULLSCREEN(1)
//             XOFFSET(4)
//             YOFFSET(4)

// unknown3(12)

// TRACK SETTINGS
//     COUNT(4)
//         Tracks
//             SIZE(1) TEXT(S)
//             ADDR(4)
//             CNTR(1)
//             unknown4(11)
//             +VAL(4)
//             -VAL(4)
//             NEUT(4)
//             ENBL(1)
//             unknown5(1)
//             unknown6(2)
//             SIZE(1) PORT(S)
//             unknown7(12)

func eventFrom(reader io.Reader) (*event, error) {

}

func eventsFrom(reader io.Reader, h header) ([]event, error) {
	events := make([]event, h.eventCount)

	for i := 0; i < int(h.eventCount); i++ {
		e, err := eventFrom(reader)
		if err != nil {
			return nil, err
		}
		events[i] = *e
	}

	return events, nil
}

// NewFileFrom creates a new File from a ByteReader
func NewFileFrom(reader io.Reader) (*File, error) {
	f := File{}

	h, err := headerFrom(reader)
	if err != nil {
		return nil, err
	}
	f.header = *h

	es, err := eventsFrom(reader, f.header)
	if err != nil {
		return nil, err
	}
	f.events = es

	return &f, nil
}
