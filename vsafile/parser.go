package vsafile

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
)

func bytesFrom(reader readerWithIndex, length int) ([]byte, error) {
	bytes := make([]byte, length)

	bytesRead, err := io.ReadFull(reader, bytes)
	if err != nil {
		return nil, err
	}
	if bytesRead != length {
		return nil, fmt.Errorf("expected to read %d byte(s), read %d", length, bytesRead)
	}

	return bytes, nil
}

func integerFrom(reader readerWithIndex, length int) (int, error) {
	switch length {
	case 1, 2, 4, 8:
		// 👍🏻
	default:
		return 0, fmt.Errorf("length must be 1, 2, 4 or 8, was %d", length)
	}

	bytesRead, err := bytesFrom(reader, length)
	if err != nil {
		return 0, err
	}

	i := 0

	switch length {
	case 1:
		i = int(bytesRead[0])
	case 2:
		i = int(binary.LittleEndian.Uint16(bytesRead))
	case 4:
		i = int(binary.LittleEndian.Uint32(bytesRead))
	case 8:
		i = int(binary.LittleEndian.Uint64(bytesRead))
	}

	return i, nil
}

func stringFrom(reader readerWithIndex, lengthCount int) (string, error) {
	length, err := integerFrom(reader, lengthCount)
	if err != nil {
		return "", err
	}

	stringBytes, err := bytesFrom(reader, length)
	if err != nil {
		return "", err
	}

	return string(stringBytes), nil
}

func validateSignature(readSignature []byte) error {
	knownSignature := []byte{0x0a, 0xd7, 0xa3, 0x70, 0x3d, 0x0a, 0x18, 0x40, 0x01, 0x00, 0x00, 0x00}

	for i := 0; i < 12; i++ {
		if (readSignature)[i] != knownSignature[i] {
			return fmt.Errorf("invalid file signature")
		}
	}

	return nil
}

func unknownOneFrom(reader readerWithIndex) ([]byte, error) {
	bytesRead, err := bytesFrom(reader, 12)
	if err != nil {
		return nil, err
	}

	err = validateSignature(bytesRead)
	if err != nil {
		return nil, err
	}

	return bytesRead, nil
}

func levelFrom(reader readerWithIndex) (string, error) {
	l, err := stringFrom(reader, 1)
	if err != nil {
		return "", err
	}

	return l, nil
}

func optionsFrom(reader readerWithIndex) ([]byte, error) {
	o, err := stringFrom(reader, 1)
	if err != nil {
		return nil, err
	}

	return []byte(o), nil
}

func emailFrom(reader readerWithIndex) (string, error) {
	e, err := stringFrom(reader, 1)

	if err != nil {
		return "", err
	}

	return e, nil
}

func eventCountFrom(reader readerWithIndex) (int, error) {
	c, err := integerFrom(reader, 4)
	if err != nil {
		return 0, err
	}

	return c, nil
}

func unknownTwoFrom(reader readerWithIndex) ([]byte, error) {
	bytesRead, err := bytesFrom(reader, 4)

	if err != nil {
		return nil, err
	}

	return bytesRead, nil
}

func eventTypeFrom(reader readerWithIndex) (string, error) {
	et, err := stringFrom(reader, 2)
	if err != nil {
		return "", err
	}
	return et, nil
}

func headerFrom(reader readerWithIndex) (*header, error) {
	u1, err := unknownOneFrom(reader)
	if err != nil {
		return nil, err
	}

	lev, err := levelFrom(reader)
	if err != nil {
		return nil, err
	}

	opts, err := optionsFrom(reader)
	if err != nil {
		return nil, err
	}

	email, err := emailFrom(reader)
	if err != nil {
		return nil, err
	}

	ec, err := eventCountFrom(reader)
	if err != nil {
		return nil, err
	}

	u2, err := unknownTwoFrom(reader)
	if err != nil {
		return nil, err
	}

	oet := ""
	det, err := eventTypeFrom(reader)
	if err != nil {
		return nil, err
	}
	switch det {
	case "CEventBarLinear":
		oet = "CEventBarPulse"
	case "CEventBar":
		oet = "CEventBarLinear"
	}

	return &header{
		unknownOne:       u1,
		level:            lev,
		options:          opts,
		email:            email,
		eventCount:       ec,
		unknownTwo:       u2,
		defaultEventType: det,
		otherEventType:   oet,
	}, nil
}

func newEventFrom(reader readerWithIndex, eventNumber int, currentEventType string) (*event, error) {
	track, err := integerFrom(reader, 2)
	if err != nil {
		return nil, err
	}

	startTime, err := integerFrom(reader, 4)
	if err != nil {
		return nil, err
	}

	endTime, err := integerFrom(reader, 4)
	if err != nil {
		return nil, err
	}

	startPosition, err := integerFrom(reader, 4)
	if err != nil {
		return nil, err
	}

	endPosition, err := integerFrom(reader, 4)
	if err != nil {
		return nil, err
	}

	lengthOfUnknownFour, err := integerFrom(reader, 1)
	if err != nil {
		return nil, err
	}

	data := []byte{}
	switch currentEventType {
	case "CEventBarLinear":
		data, err = bytesFrom(reader, 12)
	case "CEventBarPulse":
		data, err = bytesFrom(reader, 16)
	}
	if err != nil {
		return nil, err
	}

	unknownFour, err := bytesFrom(reader, lengthOfUnknownFour)
	if err != nil {
		return nil, err
	}

	continuation, err := bytesFrom(reader, 2)
	if err != nil {
		return nil, err
	}

	return &event{
		eventNumber:         eventNumber,
		_type:               currentEventType,
		track:               track,
		startTime:           startTime,
		endTime:             endTime,
		startPosition:       startPosition,
		endPosition:         endPosition,
		lengthOfUnknownFour: lengthOfUnknownFour,
		data:                data,
		unknownFour:         unknownFour,
		continuation:        hex.EncodeToString(continuation),
	}, nil
}

func eventsFrom(reader readerWithIndex, h header) ([]event, error) {
	events := make([]event, h.eventCount)

	currentEventType := h.defaultEventType

	i := 0

loop:
	for {
		// fmt.Printf("event: %d begins at index %d\n", i, reader.Index())

		e, err := newEventFrom(reader, i, currentEventType)
		if err != nil {
			return nil, err
		}
		events[i] = *e

		// fmt.Printf("event: %d type %s ends at index %d\n", i, e._type, reader.Index()-3)
		// fmt.Printf("event: %d continuation %s ends at index %d\n", i, e.continuation, reader.Index()-1)

		switch e.continuation {
		case "0100":
			// fmt.Printf("event: %d continuation: 0100 final event\n", i)
			break loop
		case "0180":
			// fmt.Printf("event %d continuation: 0180 next event is default type\n", i)
			currentEventType = h.defaultEventType
		case "3087":
			// currentEventType = h.otherEventType // maybe?
			return nil, fmt.Errorf("event: %d continuation: %s Nelson documented other event", i, e.continuation)
		case "ffff":
			_, err := bytesFrom(reader, 2) // discard two unknownFive bytes
			if err != nil {
				return nil, err
			}
			currentEventType, err = stringFrom(reader, 2)
			if err != nil {
				return nil, err
			}
			// fmt.Printf("event: %d continuation: ffff new current event type: %s\n", i, currentEventType)
		default:
			return nil, fmt.Errorf("event: %d continuation: %s is UNKNOWN", i, e.continuation)
		}
		// fmt.Println()
		i++
	}

	if i != h.eventCount-1 {
		return nil, fmt.Errorf("expected %d event(s), got %d", h.eventCount, i)
	}

	return events, nil
}

// NewFileFrom creates a new File from a ByteReader
func NewFileFrom(reader io.Reader) (*File, error) {
	cr := newCountingReader(reader)

	f := File{}

	h, err := headerFrom(cr)
	if err != nil {
		return nil, err
	}
	f.header = *h

	es, err := eventsFrom(cr, f.header)
	if err != nil {
		return nil, err
	}
	f.events = es

	return &f, nil
}
