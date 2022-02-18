package root

import (
	"fmt"

	"github.com/spf13/cobra"
)


func NewCmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "git-releaser",
		SilenceErrors: true,
		SilenceUsage: true,
		Run: func (cmd *cobra.Command, _ []string)  {
			fmt.Println("hello world");
		},
	}

	return &cmd
}