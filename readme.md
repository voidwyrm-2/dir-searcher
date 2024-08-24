# Dir Searcher
I needed to find a specific file within a bunch of folder the other day, so I though "why not make an entire file path pattern matcher?"<br>
And now we have this!(it actually works pretty well, I'm very proud of my pattern matching function!)

## NOTE TO SELF, ADD PATTERN SYSTEM EXPLAINATION

## Flags
* `-pf` | `-f` | `--printfiles`: prints the contents of files that match the pattern; default false
* `-pd` | `-p` | `--printdirs`: prints directories that match the pattern; default false
* `-w` | `--windows`: use Windows path separators instead of Unix path separators for patterns(e.g. `?-test\?.txt` instead of `?-test/?.txt`); default false<!--only use this one if you're a weirdo-->

## Installation
**Requirements**
* Go >= 1.22.1

First, install Go if you haven't already from [go.dev](go.dev)

Then, run this command:
```sh
go install "github.com/voidwyrm-2/dir-searcher@latest"
```
**If you have an issue, report it in issues!**

### This repo is licensed under the [MIT license](./LICENSE)