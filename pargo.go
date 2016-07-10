package pargo

import (
	"bytes"
	"errors"
	"io"
)

type Parser interface {
	Parse(input io.Reader) (remaining io.Reader, err error)
}

type Tag struct {
	Tag  []byte
	Done func(parsed []byte)
}

func (t Tag) Parse(input io.Reader) (remaining io.Reader, err error) {
	if t.Tag == nil {
		return nil, errors.New("no tag was defined")
	}
	buf := make([]byte, len(t.Tag))
	if _, err = io.ReadFull(input, buf); err != nil {
		return
	}
	if bytes.Compare(buf, t.Tag) != 0 {
		return nil, errors.New("could not find tag")
	}
	if t.Done != nil {
		t.Done(buf)
	}
	return input, nil
}

type Alt struct {
	Parsers []Parser
}

func (a Alt) Parse(input io.Reader) (remaining io.Reader, err error) {
	for _, p := range a.Parsers {
		copy := &bytes.Buffer{}
		i := io.TeeReader(input, copy)
		rem, err := p.Parse(i)
		if err == nil {
			return rem, nil
		}
		input = io.MultiReader(copy, input)
	}
	return nil, errors.New("no parsers matched")
}
