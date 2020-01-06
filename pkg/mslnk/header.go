package mslnk

import (
	"bytes"
	"encoding/binary"
)

type header struct {
	Data           ShellLinkHeader
	LinkFlags      map[string]bool
	FileAttributes map[string]bool
}

func Header() header {
	r := header{LinkFlags: make(map[string]bool), FileAttributes: make(map[string]bool)}
	r.Data.HeaderSize = HEADERSIZE
	r.Data.LinkCLSID = CLSID
	r.Data.ShowCommand = ShowCommand["SW_SHOWNORMAL"]
	return r
}

// writing flags into .Data bytes
func (h *header) Update() {
	var byteIndex, bitIndex int8
	byteIndex = -1
	bitIndex = 8
	for _, k := range LinkFlags {
		if bitIndex >= 8 {
			byteIndex += 1
			bitIndex = 0
			h.Data.LinkFlags[byteIndex] = 0
		}
		if h.LinkFlags[k] {
			h.Data.LinkFlags[byteIndex] += (1 << bitIndex)
		}
		bitIndex += 1
	}
	byteIndex = -1
	bitIndex = 8
	for _, k := range FileAttributesFlags {
		if bitIndex >= 8 {
			byteIndex += 1
			bitIndex = 0
			h.Data.FileAttributes[byteIndex] = 0
		}
		if h.FileAttributes[k] {
			h.Data.FileAttributes[byteIndex] += (1 << bitIndex)
		}
		bitIndex += 1
	}
}

func (h *header) Bytes() []byte {
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.LittleEndian, h.Data)
	return buffer.Bytes()
}

type ShellLinkHeader struct {
	HeaderSize     [4]byte
	LinkCLSID      [16]byte
	LinkFlags      [4]byte
	FileAttributes [4]byte
	CreationTime   [8]byte
	AccessTime     [8]byte
	WriteTime      [8]byte
	FileSize       [4]byte
	IconIndex      [4]byte
	ShowCommand    [4]byte
	HotKey         [2]byte
	Reserved1      [2]byte
	Reserved2      [4]byte
	Reserved3      [4]byte
}

var (
	HEADERSIZE = [4]byte{0x4C, 0x00, 0x00, 0x00}
	CLSID      = [16]byte{0x01, 0x14, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}

	LinkFlags = []string{
		"HasLinkTargetIDList",         // bit00 - ShellLinkHeader is followed by a LinkTargetIDList structure.
		"HasLinkInfo",                 // bit01 - LinkInfo in file.
		"HasName",                     // bit02 - NAME_String in file.
		"HasRelativePath",             // bit03 - RELATIVE_PATH in file.
		"HasWorkingDir",               // bit04 - WORKING_DIR in file.
		"HasArguments",                // bit05 - COMMAND_LINE_ARGUMENTS
		"HasIconLocation",             // bit06 - ICON_LOCATION
		"IsUnicode",                   // bit07 - Strings are in unicode
		"ForceNoLinkInfo",             // bit08 - LinkInfo is ignored
		"HasExpString",                // bit09 - The shell link is saved with an EnvironmentVariableDataBlock
		"RunInSeparateProcess",        // bit10 - Target runs in a 16-bit virtual machine
		"Unused1",                     // bit11 - ignore
		"HasDarwinID",                 // bit12 - The shell link is saved with a DarwinDataBlock
		"RunAsUser",                   // bit13 - The application is run as a different user when the target of the shell link is activated.
		"HasExpIcon",                  // bit14 - The shell link is saved with an IconEnvironmentDataBlock
		"NoPidlAlias",                 // bit15 - The file system location is represented in the shell namespace when the path to an item is parsed into an IDList.
		"Unused2",                     // bit16 - ignore
		"RunWithShimLayer",            // bit17 - The shell link is saved with a ShimDataBlock.
		"ForceNoLinkTrack",            // bit18 - The TrackerDataBlock is ignored.
		"EnableTargetMetadata",        // bit19 - The shell link attempts to collect target properties and store them in the PropertyStoreDataBlock (section 2.5.7) when the link target is set.
		"DisableLinkPathTracking",     // bit20 - The EnvironmentVariableDataBlock is ignored.
		"DisableKnownFolderTracking",  // bit21 - The SpecialFolderDataBlock (section 2.5.9) and the KnownFolderDataBlock (section 2.5.6) are ignored when loading the shell link. If this bit is set, these extra data blocks SHOULD NOT be saved when saving the shell link.
		"DisableKnownFolderAlias",     // bit22 - If the link has a KnownFolderDataBlock (section 2.5.6), the unaliased form of the known folder IDList SHOULD be used when translating the target IDList at the time that the link is loaded.
		"AllowLinkToLink",             // bit23 - Creating a link that references another link is enabled. Otherwise, specifying a link as the target IDList SHOULD NOT be allowed.
		"UnaliasOnSave",               // bit24 - When saving a link for which the target IDList is under a known folder, either the unaliased form of that known folder or the target IDList SHOULD be used.
		"PreferEnvironmentPath",       // bit25 - The target IDList SHOULD NOT be stored; instead, the path specified in the EnvironmentVariableDataBlock (section 2.5.4) SHOULD be used to refer to the target.
		"KeepLocalIDListForUNCTarget", // bit26 - When the target is a UNC name that refers to a location on a local machine, the local path IDList in the PropertyStoreDataBlock (section 2.5.7) SHOULD be stored, so it can be used when the link is loaded on the local machine.
	}

	FileAttributesFlags = []string{
		"FILE_ATTRIBUTE_READONLY",
		"FILE_ATTRIBUTE_HIDDEN",
		"FILE_ATTRIBUTE_SYSTEM",
		"Reserved1",
		"FILE_ATTRIBUTE_DIRECTORY",
		"FILE_ATTRIBUTE_ARCHIVE",
		"Reserved2",
		"FILE_ATTRIBUTE_NORMAL",
		"FILE_ATTRIBUTE_TEMPORARY",
		"FILE_ATTRIBUTE_SPARSE_FILE",
		"FILE_ATTRIBUTE_REPARSE_POINT",
		"FILE_ATTRIBUTE_COMPRESSED",
		"FILE_ATTRIBUTE_OFFLINE",
		"FILE_ATTRIBUTE_NOT_CONTENT_INDEXED",
		"FILE_ATTRIBUTE_ENCRYPTED",
	}

	ShowCommand = map[string][4]byte{
		"SW_SHOWNORMAL":      {0x01, 0x00, 0x00, 0x00},
		"SW_SHOWMAXIMIZED":   {0x03, 0x00, 0x00, 0x00},
		"SW_SHOWMINNOACTIVE": {0x07, 0x00, 0x00, 0x00},
	}
)
