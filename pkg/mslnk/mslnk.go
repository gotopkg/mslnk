package mslnk

import (
	"strings"
)

func LinkFile(target string, name string) error {
	target = strings.TrimSpace(target)
	var drive string
	if len(target) < 3 || !(target[1] == ':' && target[2] == '\\') {
		drive = "C:\\"
	} else {
		drive = target[:3]
		target = target[3:]
	}

	r := ShellLink{
		ShellLinkHeader: Header(),
		LinkTargetIDList: LinkTargetIDList{
			ItemIDList: []ItemID{
				ItemIDCLSID(ItemIDMagic["MY_COMPUTER"]),
				ItemIDDrive(drive),
				ItemIDFile(target),
			},
		},
	}

	r.LinkTargetIDList.Size()
	r.ShellLinkHeader.LinkFlags["HasLinkTargetIDList"] = true
	r.ShellLinkHeader.LinkFlags["ForceNoLinkInfo"] = true
	r.ShellLinkHeader.FileAttributes["FILE_ATTRIBUTE_NORMAL"] = true
	r.ShellLinkHeader.Update()

	return r.Save(name)
}
