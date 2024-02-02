package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type version [12]byte

func (v version) String() string {
	s := "version: "

	for _, vb := range v {
		s += fmt.Sprintf("%x", vb)
	}

	return s + "\n"
}

type level byte

func (l level) String() string {
	return "level: " + fmt.Sprintf("%x\n", l)
}

type options byte

func (o options) String() string {
	return "options: " + fmt.Sprintf("%x\n", o)
}

type email string

func (e email) String() string {
	return "email: " + string(e)
}

type eventCount uint32

func (ec eventCount) String() string {
	return fmt.Sprintf("event count: %i\n", ec)
}

type other [4]byte

func (o other) String() string {
	s := "other: "

	for _, ob := range o {
		s += fmt.Sprintf("%x", ob)
	}

	return s + "\n"
}

type event struct{}

func (e event) String() string {
	return "event: ?\n"
}

// VSAHeader VSA file header
type header struct {
	version    *version
	level      *level
	options    *options
	email      *email
	eventCount *eventCount
	other      *other
}

func (h header) String() string {
	return h.version.String() +
		h.level.String() +
		h.options.String() +
		h.email.String() +
		h.eventCount.String() +
		h.other.String()
}

// File VSA file
type File struct {
	header header
	events []event
}

func (f File) String() string {
	s := f.header.String()
	for _, e := range f.events {
		s += e.String()
	}
	return s
}

func versionFrom(br io.ByteReader) (*version, error) {
	v := version{}

	for i := 0; i < 12; i++ {
		var err error

		v[i], err = br.ReadByte()
		if err != nil {
			return nil, err
		}
	}

	return &v, nil
}

func levelFrom(br io.ByteReader) (*level, error) {
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	l := level(b)
	return &l, nil
}

func optionsFrom(br io.ByteReader) (*options, error) {
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	op := options(b)
	return &op, nil
}

func emailFrom(br io.ByteReader) (*email, error) {
	e := ""

	for {
		b, err := br.ReadByte()
		if err != nil {
			return nil, err
		}
		if b == 0 {
			break
		}
		e += string(b)
	}

	em := email(e)

	return &em, nil
}

func eventCountFrom(br io.ByteReader) (*eventCount, error) {
	ecBytes := [4]byte{}

	for i := 0; i < 4; i++ {
		var err error

		ecBytes[i], err = br.ReadByte()
		if err != nil {
			return nil, err
		}
	}

	ec := eventCount(
		int(ecBytes[0]) +
			int(ecBytes[1])*256 +
			int(ecBytes[2])*256*256 +
			int(ecBytes[3])*256*256*256)

	return &ec, nil
}

func otherFrom(br io.ByteReader) (*other, error) {
	otBytes := [4]byte{}

	for i := 0; i < 4; i++ {
		var err error

		otBytes[i], err = br.ReadByte()
		if err != nil {
			return nil, err
		}
	}
	ot := other(otBytes)
	return &ot, nil
}

func newHeaderFrom(br io.ByteReader) (*header, error) {
	header := header{}

	var err error

	header.version, err = versionFrom(br)
	if err != nil {
		return nil, err
	}
	header.level, err = levelFrom(br)
	if err != nil {
		return nil, err
	}
	header.options, err = optionsFrom(br)
	if err != nil {
		return nil, err
	}
	header.email, err = emailFrom(br)
	if err != nil {
		return nil, err
	}
	header.eventCount, err = eventCountFrom(br)
	if err != nil {
		return nil, err
	}
	header.other, err = otherFrom(br)

	return &header, nil
}

func newFileFrom(br io.ByteReader) (*File, error) {
	vf := File{}

	h, err := newHeaderFrom(br)
	if err != nil {
		return nil, err
	}
	vf.header = *h

	vf.events = []event{}

	return &vf, nil
}

func main() {
	br := bufio.NewReader(os.Stdin)

	file, err := newFileFrom(br)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("%+v", file)

	os.Exit(0)
}
