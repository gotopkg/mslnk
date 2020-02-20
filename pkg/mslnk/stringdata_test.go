package mslnk_test

import (
	"bytes"
	"github.com/gotopkg/mslnk/pkg/mslnk"
	"testing"
)

func TestStringdata(t *testing.T) {
	println("Testing StringData...")
	x := mslnk.StringData{
		"RelativePath":         mslnk.StringDataStruct("testpath"),
		"CommandLineArguments": mslnk.StringDataStruct("testABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABCABC"),
	}
	expected := []byte{8, 0, 116, 101, 115, 116, 112, 97, 116, 104, 3, 1, 116, 101, 115, 116, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67, 65, 66, 67}
	if !bytes.Equal(x.Bytes(), expected) {
		t.Errorf("StringData\n%v\ndoesn't match template\n%v\n", x.Bytes(), expected)
	}
}

func TestStringdata_Update(t *testing.T) {
	println("Testing StringData header update...")
	h := mslnk.Header()
	x := mslnk.StringData{
		"NameString":           mslnk.StringDataStruct("test name"),
		"RelativePath":         mslnk.StringDataStruct("test path"),
		"CommandLineArguments": mslnk.StringDataStruct("test argument"),
	}
	x.Update(&h)
	if !(h.LinkFlags["HasName"] && h.LinkFlags["HasArguments"] && h.LinkFlags["HasRelativePath"]) {
		t.Errorf("Missing header flags. LinkFlags: %v\n", h.LinkFlags)
	}
}
