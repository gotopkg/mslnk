package mslnk

import (
	"bytes"
	"testing"
)

func TestStringdata(t *testing.T) {
	println("Testing StringData...")
	x := StringData{
		"RelativePath":         StringDataStruct("testpath"),
		"CommandLineArguments": StringDataStruct("testABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABC"),
	}
	expected := []byte{8, 0, 116, 101, 115, 116, 112, 97, 116, 104, 3, 1, 116, 101, 115, 116, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67}
	if !bytes.Equal(x.Bytes(), expected) {
		t.Errorf("StringData\n%v\ndoesn't match template\n%v\n", x.Bytes(), expected)
	}
}

func TestStringdata_Update(t *testing.T) {
	println("Testing StringData header update...")
	h := Header()
	x := StringData{
		"NameString":           StringDataStruct("test name"),
		"RelativePath":         StringDataStruct("test path"),
		"CommandLineArguments": StringDataStruct("test argument"),
	}
	x.Update(&h)
	if !(h.LinkFlags["HasName"] && h.LinkFlags["HasArguments"] && h.LinkFlags["HasRelativePath"]) {
		t.Errorf("Missing header flags. LinkFlags: %v\n", h.LinkFlags)
	}
}
