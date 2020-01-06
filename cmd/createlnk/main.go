package main

import (
	"github.com/gotopkg/mslnk/pkg/mslnk"
	"os"
)

func main() {
	var target, linkname string

	if len(os.Args) < 3 {
		print("usage:\tcreatelnk [target] [name]\n\n" +
			"\ttarget:\tpath to link target\n" +
			"\tname:\tpath to output file\n\n")
		return
	}

	target = os.Args[1]
	linkname = os.Args[2]

	println("Writing a link to", target, "as", linkname, "...")
	if err := mslnk.LinkFile(target, linkname); err != nil {
		panic(err)
	}
	println("Done.")
}
