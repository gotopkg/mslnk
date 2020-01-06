package mslnk

// TODO: LinkInfo implementation

type LinkInfo struct {
	LinkInfoSize                    [4]byte // size in bytes of this whole struct
	LinkInfoHeaderSize              [4]byte
	LinkInfoFlags                   [4]byte
	VolumeIDOffset                  [4]byte
	LocalBasePathOffset             [4]byte
	CommonNetworkRelativeLinkOffset [4]byte
	CommonPathSuffixOffset          [4]byte

	// optional
	LocalBasePathOffsetUnicode    []byte // 4
	CommonPathSuffixOffsetUnicode []byte // 4

	// variable (and optional)
	VolumeID                  []byte // VolumeID
	LocalBasePath             []byte
	CommonNetworkRelativeLink []byte // unimplemented
	CommonPathSuffix          []byte
	LocalBasePathUnicode      []byte
	CommonPathSuffixUnicode   []byte
}

func (x *LinkInfo) Bytes() []byte {
	return nil
}

func (x *LinkInfo) SetLinkInfoFlag(flag byte, value bool) {
	// if the bit at index 'flag' is different than 'value'
	if ((x.LinkInfoFlags[0]>>flag)%2 == 1) != value {
		if value { // and value was true
			// we add that bit
			x.LinkInfoFlags[0] += 1 << flag
		} else {
			// otherwise we subtract it
			x.LinkInfoFlags[0] -= 1 << flag
		}
	}
}

type VolumeID struct {
	VolumeIDSize      [4]byte
	DriveType         [4]byte
	DriveSerialNumber [4]byte
	VolumeLabelOffset [4]byte

	// optional
	VolumeLabelOffsetUnicode []byte // 4

	// variable
	Data []byte
}

var (
	LinkInfoHeaderSize = [2][4]byte{
		{0x1C, 0x00, 0x00, 0x00},
		{0x24, 0x00, 0x00, 0x00},
	}

	LinkInfoFlags = []string{
		"VolumeIDAndLocalBasePath",
		"CommonNetworkRelativeLinkAndPathSuffix",
	}

	DriveType = map[string][4]byte{
		"DRIVE_UNKNOWN":     {0x00: 0x00, 0x00, 0x00},
		"DRIVE_NO_ROOT_DIR": {0x01, 0x00, 0x00, 0x00},
		"DRIVE_REMOVABLE":   {0x02, 0x00, 0x00, 0x00},
		"DRIVE_FIXED":       {0x03, 0x00, 0x00, 0x00},
		"DRIVE_REMOTE":      {0x04, 0x00, 0x00, 0x00},
		"DRIVE_CDROM":       {0x05, 0x00, 0x00, 0x00},
		"DRIVE_RAMDISK":     {0x06, 0x00, 0x00, 0x00},
	}
)
