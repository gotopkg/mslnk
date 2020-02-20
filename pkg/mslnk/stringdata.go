package mslnk

import (
	"bytes"
	"encoding/binary"
)

type StringData map[string][]byte

// check if proper flags are set on the header
func (s *StringData) Update(h *header) {
	if (*s)[StringDataOptions[0]] != nil { // name
		h.LinkFlags["HasName"] = true
	}
	if (*s)[StringDataOptions[3]] != nil { // arguments
		h.LinkFlags["HasArguments"] = true
	}
	for _, v := range []byte{1, 2, 4} {
		if (*s)[StringDataOptions[v]] != nil {
			h.LinkFlags["Has"+StringDataOptions[v]] = true
		}
	}
	h.Update()
}

func (s *StringData) Bytes() []byte {
	var buffer bytes.Buffer
	for _, k := range StringDataOptions {
		binary.Write(&buffer, binary.LittleEndian, (*s)[k])
	}
	return buffer.Bytes()
}

func StringDataStruct(s string) []byte {
	r := make([]byte, len(s)+2)
	r[0] = byte(len(s) & 255)
	if len(s) > 255 {
		r[1] = byte(len(s) >> 8)
	}
	for i, v := range []byte(s) {
		r[i+2] = v
	}
	return r
}

var StringDataOptions = []string{
	"NameString",
	"RelativePath",
	"WorkingDir",
	"CommandLineArguments",
	"IconLocation",
}
