package mslnk_test

import (
	"github.com/gotopkg/mslnk/pkg/mslnk"
	"testing"
)

func TestSetLinkInfoFlag(t *testing.T) {
	println("Testing SetLinkInfoFlag function...")
	li := mslnk.LinkInfo{LinkInfoHeaderSize: mslnk.LinkInfoHeaderSize[0]}
	li.SetLinkInfoFlag(1, true)
	if li.LinkInfoFlags[0] != 2 {
		t.Error("SetLinkInfoFlag failed, bit wasn't set properly.")
	}
	li.SetLinkInfoFlag(0, true)
	if li.LinkInfoFlags[0] != 3 {
		t.Error("SetLinkInfoFlag failed, bit wasn't set properly.")
	}
	li.SetLinkInfoFlag(1, false)
	if li.LinkInfoFlags[0] != 1 {
		t.Error("SetLinkInfoFlag failed, bit wasn't set properly.")
	}
	li.SetLinkInfoFlag(1, false)
	if li.LinkInfoFlags[0] != 1 {
		t.Error("SetLinkInfoFlag failed, bit wasn't set properly.")
	}
}
