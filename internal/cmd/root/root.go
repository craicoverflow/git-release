package root

import (
	"github.com/craicoverflow/git-releaser/internal/cmd/major"
	"github.com/craicoverflow/git-releaser/pkg/cmdutil"
	"github.com/craicoverflow/git-releaser/pkg/config"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:           "git-release",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmdFactory := cmdutil.CmdFactory{
		Config: config.NewFile(),
	}

	cmd.AddCommand(
		major.NewCmd(cmdFactory),
	)

	return &cmd
}
