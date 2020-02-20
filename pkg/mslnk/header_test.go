package mslnk_test

import (
	"bytes"
	"github.com/gotopkg/mslnk/pkg/mslnk"
	"testing"
)

func TestHeader(t *testing.T) {
	println("Testing Header creation...")
	var expected []byte
	h := mslnk.Header()
	expected = []byte{76, 0, 0, 0, 1, 20, 2, 0, 0, 0, 0, 0, 192, 0, 0, 0, 0, 0, 0, 70, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if !bytes.Equal(h.Bytes(), expected) {
		t.Errorf("Initialized header\n%v\ndoesn't match template\n%v\n", h.Bytes(), expected)
	}
	println("Testing Header LinkFlags and FileAttributes...")
	h.LinkFlags["HasLinkTargetIDList"] = true
	h.LinkFlags["HasLinkInfo"] = true
	h.LinkFlags["HasRelativePath"] = true
	h.LinkFlags["HasWorkingDir"] = true
	h.LinkFlags["IsUnicode"] = true
	h.LinkFlags["EnableTargetMetadata"] = true
	h.FileAttributes["FILE_ATTRIBUTE_ARCHIVE"] = true
	h.Update()
	expected = []byte{76, 0, 0, 0, 1, 20, 2, 0, 0, 0, 0, 0, 192, 0, 0, 0, 0, 0, 0, 70, 155, 0, 8, 0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if !bytes.Equal(h.Bytes(), expected) {
		t.Errorf("Initialized header\n%v\ndoesn't match template\n%v\n", h.Bytes(), expected)
	}
}

func TestHeader_DecodeFlags(t *testing.T) {
	println("Testing Header DecodeFlags...")
	h := mslnk.Header()
	h.Data.LinkFlags = [4]byte{155, 0, 8, 0}
	h.Data.FileAttributes = [4]byte{32, 0, 0, 0}
	h.DecodeFlags()
	// list of expected flags should be in order
	expected := []string{"HasLinkTargetIDList", "HasLinkInfo", "HasRelativePath", "HasWorkingDir", "IsUnicode", "EnableTargetMetadata", "FILE_ATTRIBUTE_ARCHIVE", ""}
	var i int = 0
	for _, k := range mslnk.LinkFlags {
		if h.LinkFlags[k] {
			if expected[i] == k {
				i += 1
			} else {
				t.Errorf("Flag %s set when it shouldn't be. Expecting %s first.\n", k, expected[i])
			}
		} else if expected[i] == k {
			t.Errorf("Flag %s not set when it should be.", k)
			i += 1
		}
	}
	for _, k := range mslnk.FileAttributesFlags {
		if h.FileAttributes[k] {
			if expected[i] == k {
				i += 1
			} else {
				t.Errorf("Flag %s set when it shouldn't be. Expecting %s first.\n", k, expected[i])
			}
		} else if expected[i] == k {
			t.Errorf("Flag %s not set when it should be.", k)
			i += 1
		}
	}
}
