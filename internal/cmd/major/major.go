package major

import (
	"fmt"

	"github.com/craicoverflow/git-releaser/pkg/cmdutil"
	"github.com/craicoverflow/git-releaser/pkg/config"
	"github.com/spf13/cobra"
)

type options struct {
	config config.IConfig
}

func NewCmd(f cmdutil.CmdFactory) *cobra.Command {
	opts := options{
		config: f.Config,
	}

	cmd := cobra.Command{
		Use: "major",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runCmd(opts)
		},
	}

	return &cmd
}

func runCmd(opts options) error {
	cfg, err := opts.config.Load()
	if err != nil {
		return err
	}
	versions := cfg.Project().VersionList()

	for _, v := range versions {
		fmt.Println("Printing version " + v.String())
	}
	return nil
}
