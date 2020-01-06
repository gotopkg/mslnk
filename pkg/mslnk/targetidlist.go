package mslnk

// TODO: targetidlist_test

import (
	"bytes"
)

type ItemID struct {
	Size [2]byte
	Data []byte
}

func ItemIDFile(s string) ItemID {
	r := ItemID{Data: make([]byte, len(s)+13)}
	r.Size[0] = byte((len(r.Data) + 2) & 255)
	if (len(r.Data) + 2) > 255 {
		r.Size[1] = byte((len(r.Data) + 2) - 255)
	}
	r.Data[0] = ItemIDType["file"]
	r.Data[10] = ItemIDMagic["file_attr"][0]
	for i, v := range []byte(s) {
		r.Data[i+12] = v
	}
	return r
}

// TODO, FIXME: test ItemIDDirectory
func ItemIDDirectory(s string) ItemID {
	r := ItemID{Data: make([]byte, len(s)+13)}
	r.Size[0] = byte((len(r.Data) + 2) & 255)
	if (len(r.Data) + 2) > 255 {
		r.Size[1] = byte((len(r.Data) + 2) - 255)
	}
	r.Data[0] = ItemIDType["directory"]
	r.Data[10] = ItemIDMagic["directory_attr"][0]
	for i, v := range []byte(s) {
		r.Data[i+12] = v
	}
	return r
}

func ItemIDDrive(s string) ItemID {
	r := ItemID{
		Size: [2]byte{0x1c, 0x00},
		Data: make([]byte, 0x1c-2),
	}
	r.Data[0] = ItemIDType["drive"]
	for i, v := range []byte(s) {
		r.Data[i+1] = v
	}
	return r
}

func ItemIDCLSID(b []byte) ItemID {
	r := make([]byte, len(b)+3)
	r[0] = byte(len(r) & 255)
	if len(r) > 255 {
		r[1] = byte(len(r) - 255)
	}
	r[2] = ItemIDType["clsid"]
	for i, v := range b {
		r[i+3] = v
	}
	return ItemID{
		Size: [2]byte{r[0], r[1]},
		Data: r[2:],
	}
}

type LinkTargetIDList struct {
	IDListSize [2]byte

	// 'IDlist' from spec
	ItemIDList []ItemID
	TerminalID [2]byte
}

func (x *LinkTargetIDList) Size() uint16 {
	var r uint16 = 2
	for _, v := range x.ItemIDList {
		r += uint16(len(v.Data) + len(v.Size))
	}
	x.IDListSize[0] = byte(r & 255)
	if r > 255 {
		x.IDListSize[1] = byte(r - 255)
	} else {
		x.IDListSize[1] = 0
	}
	return r
}

func (x *LinkTargetIDList) Bytes() []byte {
	var buffer bytes.Buffer
	buffer.Write(x.IDListSize[:])
	for _, v := range x.ItemIDList {
		buffer.Write(v.Size[:])
		buffer.Write(v.Data)
	}
	buffer.Write(x.TerminalID[:])
	return buffer.Bytes()
}

// based on https://github.com/DmitriiShamrikov/mslinks/blob/master/src/mslinks/data/ItemID.java
var (
	ItemIDType = map[string]byte{
		"file":      0x32,
		"directory": 0x31,
		"drive":     0x2f,
		"clsid":     0x1f,
	}

	ItemIDMagic = map[string][]byte{
		"MY_COMPUTER":    {0x50, 0xe0, 0x4f, 0xd0, 0x20, 0xea, 0x3a, 0x69, 0x10, 0xa2, 0xd8, 0x08, 0x00, 0x2b, 0x30, 0x30, 0x9d},
		"file_attr":      {0x20, 0x00},
		"directory_attr": {0x10, 0x00},
	}
)
