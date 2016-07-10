package pargo

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestTag_Parse_ok(t *testing.T) {
	found := false
	tag := Tag{
		Tag: []byte("fart"),
		Done: func(parsed []byte) {
			found = true
		},
	}
	remaining, err := tag.Parse(bytes.NewBufferString("fart cheese"))
	if err != nil {
		t.Fatalf("expected no error, but got: %s", err)
	}
	if !found {
		t.Fatalf("expected tag to be found, but wasn't")
	}
	remBytes, err := ioutil.ReadAll(remaining)
	if err != nil {
		t.Fatalf("expected no error, but got: %s", err)
	}
	if !reflect.DeepEqual([]byte(" cheese"), remBytes) {
		t.Fatalf("expected rem to be ' cheese', but got '%s'", remBytes)
	}
}

func TestAlt_Parse_ok(t *testing.T) {
	found := false
	alt := Alt{
		Parsers: []Parser{
			Tag{Tag: []byte("bacon")},
			Tag{Tag: []byte("cheese"), Done: func(parsed []byte) { found = true }},
			Tag{Tag: []byte("pickle")},
		},
	}
	remaining, err := alt.Parse(bytes.NewBufferString("cheeseburger"))
	if err != nil {
		t.Fatalf("expected no error, but got: %s", err)
	}
	if !found {
		t.Fatalf("expected cheese to be found, but wasn't")
	}
	remBytes, err := ioutil.ReadAll(remaining)
	if err != nil {
		t.Fatalf("expected no error, but got: %s", err)
	}
	if !reflect.DeepEqual([]byte("burger"), remBytes) {
		t.Fatalf("expected rem to be 'burger', but got '%s'", remBytes)
	}
}
