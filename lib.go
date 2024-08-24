package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func readFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	content := ""
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

func pathCompare(path, pattern string, info *os.DirEntry, wholePattern *[]string) bool {
	pattern = strings.TrimSpace(pattern)
	if pattern == "" {
		fmt.Println("pattern error: pattern is empty")
		return false
	}

	if strings.Contains(pattern, "&&") {
		p := strings.SplitN(pattern, "&&", 2)
		return pathCompare(path, strings.TrimSpace(p[0]), info, wholePattern) && pathCompare(path, strings.TrimSpace(p[1]), info, wholePattern)
	} else if strings.Contains(pattern, "||") {
		p := strings.SplitN(pattern, "||", 2)
		return pathCompare(path, strings.TrimSpace(p[0]), info, wholePattern) || pathCompare(path, strings.TrimSpace(p[1]), info, wholePattern)
	}

	if pattern == "*" {
		return true
	} else if strings.HasPrefix(pattern, "**") {
		if pattern == "**" {
			*wholePattern = append(*wholePattern, "**")
		}

		res, err := strconv.Atoi(pattern[2:])
		if err != nil {
			fmt.Println("pattern error:", err.Error())
		}

		if res > 0 {
			*wholePattern = append(*wholePattern, fmt.Sprintf("**%d", res-1))
		}

		return true
	}

	defaultingHasPrefix := func(s, prefix string) bool {
		if prefix == "" {
			return true
		}
		return strings.HasPrefix(s, prefix)
	}

	processFlag := func(flag string, i *os.DirEntry) bool {
		switch strings.ToLower(strings.TrimSpace(flag)) {
		case "dir":
			return (*i).IsDir()
		}
		return false
	}

	defaultingHasSuffix := func(s, suffix string) bool {
		if suffix == "" {
			return true
		}
		return strings.HasSuffix(s, suffix)
	}

	if pattern[0] == '!' {
		return !pathCompare(path, pattern[1:], info, wholePattern)
	} else if pattern[0] == '=' {
		return path == strings.TrimSpace(pattern[1:])
	} else if pattern[0] == '#' {
		return strings.Contains(path, pattern[1:])
	}

	if strings.Contains(pattern, "%") {
		p := strings.SplitN(pattern, "%", 2)
		flag, sub := strings.TrimSpace(p[0]), strings.TrimSpace(p[1])
		if sub != "" {
			return processFlag(flag, info) && pathCompare(path, sub, info, wholePattern)
		}
		return processFlag(flag, info)
	}

	if strings.Contains(pattern, "?") {
		p := strings.SplitN(pattern, "?", 2)
		prefix, suffix := strings.TrimSpace(p[0]), strings.TrimSpace(p[1])
		return defaultingHasPrefix(path, prefix) && defaultingHasSuffix(path, suffix)
	}

	return false
}

func followPath(cpath string, pathPattern []string) {
	if len(pathPattern) == 0 {
		return
	}

	cpath = path.Clean(cpath)
	files, err := os.ReadDir(cpath)
	if err != nil {
		fmt.Printf("encountered error while reading directory '%s': %s\n", cpath, err.Error())
		return
	}

	for _, f := range files {
		//fmt.Println(f.Name(), pathPattern[0])
		if pathCompare(f.Name(), strings.TrimSpace(pathPattern[0]), &f, &pathPattern) {
			fname := path.Clean(cpath + string(os.PathSeparator) + f.Name())
			if f.IsDir() {
				if printDirs {
					fmt.Println("directory:", fname)
				}
				followPath(fname, pathPattern[1:])
			} else {
				fmt.Println("file:", fname)
				if printFileContents {
					content, err := readFile(fname)
					if err != nil {
						fmt.Printf("error printing contents of file '%s': %s\n", fname, err.Error())
					}
					fmt.Printf("[CONTENT]\n%s\n[END]\n\n", content)
				}
			}
		}
	}
}
