package project

import (
	"github.com/craicoverflow/git-releaser/pkg/version"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

type Project struct {
	versionList []*version.Version
	remoteURL   string
}

func New(remoteURL string) (*Project, error) {
	project := Project{
		remoteURL: remoteURL,
	}

	remote, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: remoteURL,
	})
	if err != nil {
		return nil, err
	}

	tags, err := remote.Tags()
	if err != nil {
		return nil, err
	}

	_ = tags.ForEach(func(r *plumbing.Reference) error {
		parsedVersion, err := version.Parse(r)
		if err == nil {
			project.versionList = append(project.versionList, parsedVersion)
		}
		return nil
	})

	return &project, nil
}

func (p *Project) VersionList() []*version.Version {
	return p.versionList
}
