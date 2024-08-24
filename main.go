package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var printFileContents bool = false
var printDirs bool = false
var patternSeparator string = "/"

func main() {
	printFileContents = slices.Contains(os.Args, "-pf") || slices.Contains(os.Args, "-f") || slices.Contains(os.Args, "--printfiles")
	printDirs = slices.Contains(os.Args, "-pd") || slices.Contains(os.Args, "-p") || slices.Contains(os.Args, "--printdirs")
	if slices.Contains(os.Args, "-w") || slices.Contains(os.Args, "--windows") {
		patternSeparator = "\\"
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input path search pattern\n> ")
	scanner.Scan()
	path := strings.TrimSpace(scanner.Text())
	if path == "" {
		fmt.Println("path search pattern cannot be empty")
		return
	}
	followPath("", strings.Split(path, patternSeparator))

	// tested paths:

	// =testdir / dir% || ?.txt / dir% || ?.txt / ?.txt

	// =testdir / dir% || #o || #s / dir% || #o || #s / dir% || #o || #s

	// =testdir / dir% || #o || #s / dir% || #o || #s / ?
}
