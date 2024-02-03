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
	lengthByte := make([]byte, 1)

	n, err := io.ReadFull(reader, lengthByte)
	if err != nil {
		return nil, err
	}
	if n != 1 {
		return nil, fmt.Errorf("expected to read 1 byte, actually read %d", n)
	}

	stringLength := int(lengthByte[0])

	stringBytes := make([]byte, stringLength)

	n, err = io.ReadFull(reader, stringBytes)
	if err != nil {
		return nil, err
	}
	if n != stringLength {
		return nil, fmt.Errorf("expected to read %d byte, actually read %d", stringLength, n)
	}

	s := string(stringBytes)

	return &s, nil
}

func versionFrom(reader io.Reader) (*version, error) {
	bytesRead, err := bytesFrom(reader, 12)
	if err != nil {
		return nil, err
	}

	return (*version)(bytesRead), nil
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

func otherFrom(reader io.Reader) (*other, error) {
	bytesRead, err := bytesFrom(reader, 4)

	if err != nil {
		return nil, err
	}

	return (*other)(bytesRead), nil
}

func headerFrom(reader io.Reader) (*header, error) {
	h := header{}

	v, err := versionFrom(reader)
	if err != nil {
		return nil, err
	}
	h.version = *v

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

	ot, err := otherFrom(reader)
	if err != nil {
		return nil, err
	}
	h.other = *ot

	return &h, nil
}

// NewFileFrom creates a new File from a ByteReader
func NewFileFrom(reader io.Reader) (*File, error) {
	f := File{}

	h, err := headerFrom(reader)
	if err != nil {
		return nil, err
	}
	f.header = *h

	f.events = []event{}

	return &f, nil
}
