# mslnk
> Create Microsoft .LNK files with Go 

mslnk is a Golang module (with CLI) implementing Microsoft's Shell Link Binary File Format (MS-SHLLINK) without external dependencies

## Introduction

Sometimes you may need to link somewhere with a [shortcut file](https://en.wikipedia.org/wiki/Shortcut_(computing)) which contains additional details and can be moved around freely. That's the purpose of shell link files on Window$. They are something like [XDG Desktop Entries](https://specifications.freedesktop.org/desktop-entry-spec/desktop-entry-spec-latest.html) (*.desktop files), but obfuscated in a binary file format and documented in a worse manner.

**This project's goal is to make the creation of those files painless** by bringing you simple command line tools and exposing an easy-to-understand library API for your Go programs.
Creating a list of relevant resources would also be a desirable side effect.

### Usage
#### command line
##### createlnk
```
$ ./createlnk 
usage:	createlnk [target] [name]

	target:	path to link target
	name:	path to output file
```

#### programming
```
import "github.com/gotopkg/mslnk/pkg/mslnk"

func main() {
	mslnk.LinkFile("C:\\path-to\\a-file.txt", "mylink.lnk")
}
```
Despite what the function name might suggest, linking to a folder that way should work too.
Basic functions reside in pkg/mslnk/mslnk.go
If you are familiar with the format specification, you can manipulate exposed structures directly.

### Installation
It's as simple as `go get` and `go build`. All you need is a Go compiler!
```
$ go get -u github.com/gotopkg/mslnk
$ go build github.com/gotopkg/mslnk/cmd/createlnk
```

## Implementation
mslnk is not yet a full implementation, but it already gets the job done.

### Resources
#### Specification
Implementation is based on [microsoft's specification](https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-shllink/16cb4ca1-9339-4d0c-a68d-bf1d6cc0f943), revision 5.0.
To be precise, I was using [this pdf](https://winprotocoldoc.blob.core.windows.net/productionwindowsarchives/MS-SHLLINK/%5bMS-SHLLINK%5d.pdf) during development.

#### Libraries
##### Go
Some structs were copied from [github.com/parsiya/golnk](https://github.com/parsiya/golnk/) with slight modifications. I wanted to fork it at first, but 
the code was of no use for me, so I wrote the rest from scratch. It is a nice project if you want to parse and read .lnk files instead of writing them.

##### Java
The other useful project is [Dmitrii Shamrikov's mslinks](https://github.com/DmitriiShamrikov/mslinks) Java library. I used it as a reference for the implementation of the [_undocumented ItemID Data field_](https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-shllink/6ac3b286-6640-4cf3-85f2-4085c02e6a71).

##### Python
I am also aware of the [pylnk](https://github.com/strayge/pylnk) project for Python, but have not looked into it.

##### C, bash
There are some scripts, which do what our createlnk does, but in a hacky way.
Most notably: https://www.mamachine.org/mslink/index.en.html

## Contribute
### Resources
If you know of any additional documents I might have overlooked, or resources which might be helpful, please let me know. You are welcome to list it on this page.

