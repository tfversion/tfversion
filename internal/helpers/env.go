package helpers

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func WarnIfNotInPath(checkPath string) {
	path := os.Getenv("PATH")
	if !strings.Contains(path, checkPath) {
		fmt.Printf("%s not found in your shell PATH\n", color.CyanString(checkPath))
		fmt.Printf("Please run %s to make this version available in your shell\n", color.CyanString("`export PATH=%s:$PATH`", checkPath))
		fmt.Printf("Additionally, consider adding this line to your shell profile (e.g., .bashrc, .zshrc or fish config) for persistence.\n")
	}
}
