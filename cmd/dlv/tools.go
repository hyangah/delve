package main

// List packages used by _scripts

//go:generate go run ../../_scripts/gen-cli-docs.go ../../Documentation/cli
//go:generate go run ../../_scripts/gen-usage-docs.go ../../Documentation/usage

import (
	_ "github.com/spf13/cobra/doc"
)
