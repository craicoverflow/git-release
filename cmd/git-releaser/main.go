package main

import (
	"os"

	"github.com/craicoverflow/git-releaser/internal/cmd/root"
)

func main() {
	rootCmd := root.NewCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
