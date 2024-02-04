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

func integerFrom(reader io.Reader, length int) (*int, error) {
	switch length {
	case 1, 2, 4, 8:
	default:
		return nil, fmt.Errorf("length must be 1, 2, 4 or 8, was %d", length)
	}

	bytesRead, err := bytesFrom(reader, length)
	if err != nil {
		return nil, err
	}

	i := 0

	switch length {
	case 1:
		i = int((*bytesRead)[0])
	case 2:
		i = int(binary.LittleEndian.Uint16(*bytesRead))
	case 4:
		i = int(binary.LittleEndian.Uint32(*bytesRead))
	case 5:
		i = int(binary.LittleEndian.Uint64(*bytesRead))
	}

	return &i, nil
}

func stringFrom(reader io.Reader, lengthCount int) (*string, error) {
	length, err := integerFrom(reader, lengthCount)
	if err != nil {
		return nil, err
	}

	stringBytes, err := bytesFrom(reader, *length)
	if err != nil {
		return nil, err
	}

	s := string(*stringBytes)

	return &s, nil
}

func validateSignature(readSignature *[]byte) error {
	knownSignature := []byte{0x0a, 0xd7, 0xa3, 0x70, 0x3d, 0x0a, 0x18, 0x40, 0x01, 0x00, 0x00, 0x00}

	for i := 0; i < 12; i++ {
		if (*readSignature)[i] != knownSignature[i] {
			return fmt.Errorf("invalid file signature")
		}
	}

	return nil
}

func unknownOneFrom(reader io.Reader) (*unknownOne, error) {
	bytesRead, err := bytesFrom(reader, 12)
	if err != nil {
		return nil, err
	}

	err = validateSignature(bytesRead)
	if err != nil {
		return nil, err
	}

	return (*unknownOne)(bytesRead), nil
}

func levelFrom(reader io.Reader) (*level, error) {
	l, err := stringFrom(reader, 1)
	if err != nil {
		return nil, err
	}

	return (*level)(l), nil
}

func optionsFrom(reader io.Reader) (*options, error) {
	o, err := stringFrom(reader, 1)
	if err != nil {
		return nil, err
	}

	os := options(*o)

	return &os, nil
}

func emailFrom(reader io.Reader) (*email, error) {
	es, err := stringFrom(reader, 1)

	if err != nil {
		return nil, err
	}

	return (*email)(es), nil
}

func eventCountFrom(reader io.Reader) (*eventCount, error) {
	c, err := integerFrom(reader, 4)
	if err != nil {
		return nil, err
	}

	ec := eventCount(*c)

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

// func eventFrom(reader io.Reader) (*event, error) {}

// func eventsFrom(reader io.Reader, h header) ([]event, error) {
// 	events := make([]event, h.eventCount)

// 	for i := 0; i < int(h.eventCount); i++ {
// 		e, err := eventFrom(reader)
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
