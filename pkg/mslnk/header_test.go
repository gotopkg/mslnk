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
