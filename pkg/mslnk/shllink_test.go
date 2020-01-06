package mslnk_test

import (
	"bytes"
	"github.com/gotopkg/mslnk/pkg/mslnk"
	"testing"
)

func TestShellLink_Bytes(t *testing.T) {
	println("Testing ShellLink.Bytes...")
	r := mslnk.ShellLink{
		ShellLinkHeader: mslnk.Header(),
		LinkTargetIDList: mslnk.LinkTargetIDList{
			ItemIDList: []mslnk.ItemID{
				mslnk.ItemIDCLSID(mslnk.ItemIDMagic["MY_COMPUTER"]),
				mslnk.ItemIDDrive("C:\\"),
				mslnk.ItemIDFile("test\\this.txt"),
			},
		},
	}
	r.LinkTargetIDList.Size()
	r.ShellLinkHeader.LinkFlags["HasLinkTargetIDList"] = true
	r.ShellLinkHeader.LinkFlags["ForceNoLinkInfo"] = true
	r.ShellLinkHeader.FileAttributes["FILE_ATTRIBUTE_NORMAL"] = true
	r.ShellLinkHeader.Update()
	expected := []byte{76, 0, 0, 0, 1, 20, 2, 0, 0, 0, 0, 0, 192, 0, 0, 0, 0, 0, 0, 70, 1, 1, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 78, 0, 20, 0, 31, 80, 224, 79, 208, 32, 234, 58, 105, 16, 162, 216, 8, 0, 43, 48, 48, 157, 28, 0, 47, 67, 58, 92, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 28, 0, 50, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 116, 101, 115, 116, 92, 116, 104, 105, 115, 46, 116, 120, 116, 0, 0, 0}
	if !bytes.Equal(r.Bytes(), expected) {
		t.Errorf("Link bytes\n%v\ndon't match expected values\n%v\n", r.Bytes(), expected)
	}
}
