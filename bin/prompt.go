// From Drew DeVault's dotfiles repository
// https://git.sr.ht/~sircmpwn/dotfiles/tree/ce826c78ce3cd5fcef5d9f0d64d44b67f1098c3b/item/bin/prompt.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	cwd, _ := os.Getwd()
	home := os.Getenv("HOME")
	var parts []string
	if strings.HasPrefix(cwd, home) {
		cwd = "~" + cwd[len(home):]
	}
	parts = strings.Split(cwd, "/")
	for i, part := range parts {
		if i == len(parts)-1 {
			fmt.Printf("%s\n", part)
		} else {
			if len(part) != 0 {
				fmt.Printf("%c/", part[0])
			} else {
				fmt.Printf("/")
			}
		}
	}
}
