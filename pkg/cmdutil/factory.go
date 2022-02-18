package cmdutil

import "github.com/craicoverflow/git-releaser/pkg/config"

type CmdFactory struct {
	Config config.IConfig
}

type BaseOptions struct {
	Config config.IConfig
}
