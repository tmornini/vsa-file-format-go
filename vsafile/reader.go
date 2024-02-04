package vsafile

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
)

func bytesFrom(reader io.Reader, length int) ([]byte, error) {
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

func integerFrom(reader io.Reader, length int) (int, error) {
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
	case 5:
		i = int(binary.LittleEndian.Uint64(bytesRead))
	}

	return i, nil
}

func stringFrom(reader io.Reader, lengthCount int) (string, error) {
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

func unknownOneFrom(reader io.Reader) ([]byte, error) {
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

func levelFrom(reader io.Reader) (string, error) {
	l, err := stringFrom(reader, 1)
	if err != nil {
		return "", err
	}

	return l, nil
}

func optionsFrom(reader io.Reader) ([]byte, error) {
	o, err := stringFrom(reader, 1)
	if err != nil {
		return nil, err
	}

	return []byte(o), nil
}

func emailFrom(reader io.Reader) (string, error) {
	e, err := stringFrom(reader, 1)

	if err != nil {
		return "", err
	}

	return e, nil
}

func eventCountFrom(reader io.Reader) (int, error) {
	c, err := integerFrom(reader, 4)
	if err != nil {
		return 0, err
	}

	return c, nil
}

func unknownTwoFrom(reader io.Reader) ([]byte, error) {
	bytesRead, err := bytesFrom(reader, 4)

	if err != nil {
		return nil, err
	}

	return bytesRead, nil
}

func headerFrom(reader io.Reader) (*header, error) {
	h := &header{}

	v, err := unknownOneFrom(reader)
	if err != nil {
		return nil, err
	}
	h.unknownOne = v

	l, err := levelFrom(reader)
	if err != nil {
		return nil, err
	}
	h.level = l

	os, err := optionsFrom(reader)
	if err != nil {
		return nil, err
	}
	h.options = os

	e, err := emailFrom(reader)
	if err != nil {
		return nil, err
	}
	h.email = e

	ec, err := eventCountFrom(reader)
	if err != nil {
		return nil, err
	}
	h.eventCount = ec

	ot, err := unknownTwoFrom(reader)
	if err != nil {
		return nil, err
	}
	h.unknownTwo = ot

	fe, err := stringFrom(reader, 2)
	if err != nil {
		return nil, err
	}
	h.firstEventType = fe

	return h, nil
}

func newEventFrom(reader io.Reader, eventNumber int, kind string) (*event, error) {
	e := event{eventNumber: eventNumber}

	track, err := integerFrom(reader, 2)
	if err != nil {
		return nil, err
	}
	e.track = track

	startTime, err := integerFrom(reader, 4)
	if err != nil {
		return nil, err
	}
	e.startTime = startTime

	endTime, err := integerFrom(reader, 4)
	if err != nil {
		return nil, err
	}
	e.endTime = endTime

	startPosition, err := integerFrom(reader, 4)
	if err != nil {
		return nil, err
	}
	e.startPosition = startPosition

	endPosition, err := integerFrom(reader, 4)
	if err != nil {
		return nil, err
	}
	e.endPosition = endPosition

	data := []byte{}

	switch kind {
	case "CEventBarLinear":
		data, err = bytesFrom(reader, 12)
	case "CEventBar":
		data, err = bytesFrom(reader, 16)
	}
	if err != nil {
		return nil, err
	}
	e.data = data

	continuation, err := bytesFrom(reader, 2)
	if err != nil {
		return nil, err
	}
	e.continuation = hex.EncodeToString(continuation)

	return &e, nil
}

// func eventsFrom(reader io.Reader, h header) ([]event, error) {
// 	events := make([]event, h.eventCount)

// 	for i := 0; i < int(h.eventCount); i++ {
// 		e, err := newEventFrom(reader)
// 		if err != nil {
// 			return nil, err
// 		}
// 		events[i] = *e
// 	}

// 	return events, nil
// }

// NewFileFrom creates a new File from a ByteReader
func NewFileFrom(reader io.Reader) (*File, error) {
	f := File{}

	h, err := headerFrom(reader)
	if err != nil {
		return nil, err
	}
	f.header = *h

	// es, err := eventsFrom(reader, f.header)
	// if err != nil {
	// 	return nil, err
	// }
	// f.events = es

	return &f, nil
}
