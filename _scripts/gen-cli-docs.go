//go:build ignore
// +build ignore

package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"

	"github.com/go-delve/delve/pkg/terminal"
)

const defaultUsageDir = "./Documentation/cli"

func main() {
	usageDir := defaultUsageDir
	if len(os.Args) > 1 {
		usageDir = os.Args[1]
	}
	usageDir = filepath.Clean(usageDir)
	fh, err := os.Create(os.ExpandEnv(filepath.Join(usageDir, "README.md")))
	if err != nil {
		log.Fatalf("could not create README.md: %v", err)
	}
	defer fh.Close()

	w := bufio.NewWriter(fh)
	defer w.Flush()

	commands := terminal.DebugCommands(nil)
	commands.WriteMarkdown(w)
}
