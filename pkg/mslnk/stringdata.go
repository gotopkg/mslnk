package mslnk

import (
	"bytes"
	"encoding/binary"
)

type stringdata map[string][]byte

// check if proper flags are set on the header
func (s *stringdata) Update(h *header) {
	var hasString bool = false
	if (*s)[StringData[0]] != nil { // name
		h.LinkFlags["HasName"] = true
		hasString = true
	}
	if (*s)[StringData[3]] != nil { // arguments
		h.LinkFlags["HasArguments"] = true
		hasString = true
	}
	for _, v := range []byte{1, 2, 4} {
		if (*s)[StringData[v]] != nil {
			h.LinkFlags["Has"+StringData[v]] = true
			hasString = true
		}
	}
	h.LinkFlags["IsUnicode"] = hasString
	h.Update()
}

func (s *stringdata) Bytes() []byte {
	var buffer bytes.Buffer
	for _, k := range StringData {
		binary.Write(&buffer, binary.LittleEndian, (*s)[k])
	}
	return buffer.Bytes()
}

func StringDataStruct(s string) []byte {
	r := make([]byte, len(s)+2)
	r[0] = byte(len(s) & 255)
	if len(s) > 255 {
		r[1] = byte(len(s) - 255)
	}
	for i, v := range []byte(s) {
		r[i+2] = v
	}
	return r
}

var StringData = []string{
	"NameString",
	"RelativePath",
	"WorkingDir",
	"CommandLineArguments",
	"IconLocation",
}
