package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

var currentDir = &dir{}

type dir struct {
	size   int
	parent *dir
}

var dirsFlat = make(map[*dir]int)

func main() {
	cmds := strings.Split(strings.TrimSpace(input), "\n")

	ans := 0
	ans2 := 0

	root := currentDir

	for _, cmd := range cmds {
		if strings.HasPrefix(cmd, "$") {
			// is a command
			if strings.HasPrefix(cmd[2:], "cd") {
				// change dir
				if strings.HasPrefix(cmd[5:], "..") {
					// go up
					currentDir = currentDir.parent
					continue
				}
				// go in
				destinationDir := &dir{
					parent: currentDir,
				}
				currentDir = destinationDir
				continue
			}
		}
		if unicode.IsNumber(rune(cmd[0])) {
			// is a file
			lineSplit := strings.Split(cmd, " ")
			fileSize, _ := strconv.Atoi(lineSplit[0])
			currentDir.size += fileSize
			dirsFlat[currentDir] = currentDir.size

			walkDir := currentDir
			for walkDir.parent != nil {
				walkDir.parent.size += fileSize
				dirsFlat[walkDir.parent] = walkDir.parent.size
				walkDir = walkDir.parent
			}

			continue
		}
	}

	unusedSpace := 70_000_000 - root.size
	neededSpace := 30_000_000 - unusedSpace

	var dirToDelete *dir

	for dir, dirSize := range dirsFlat {
		if dirSize <= 100_000 {
			ans += dirSize
		}
		if dirSize > neededSpace && (dirToDelete == nil || dirSize < dirToDelete.size) {
			dirToDelete = dir
		}
	}

	ans2 = dirToDelete.size

	fmt.Println(ans)
	fmt.Println(ans2)
}
