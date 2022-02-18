package config

import "github.com/craicoverflow/git-releaser/pkg/project"

type IConfig interface {
	Load() (*Config, error)
}

type Config struct {
	RemoteURL string
	project   *project.Project
}

type file struct{}

func NewFile() IConfig {
	return &file{}
}

func (f *file) Load() (*Config, error) {
	cfg := Config{
		RemoteURL: "https://github.com/craicoverflow/git-release",
	}

	project, err := project.New(cfg.RemoteURL)
	if err != nil {
		return nil, err
	}
	cfg.project = project

	return &cfg, nil
}

func (c *Config) Project() *project.Project {
	return c.project
}