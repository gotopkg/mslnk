package mslnk

import (
	"bytes"
	"testing"
)

func TestStringdata(t *testing.T) {
	println("Testing stringdata...")
	x := stringdata{
		"RelativePath":         StringDataStruct("testpath"),
		"CommandLineArguments": StringDataStruct("test"),
	}
	expected := []byte{8, 0, 116, 101, 115, 116, 112, 97, 116, 104, 4, 0, 116, 101, 115, 116}
	if !bytes.Equal(x.Bytes(), expected) {
		t.Errorf("stringdata\n%v\ndoesn't match template\n%v\n", x.Bytes(), expected)
	}
}
