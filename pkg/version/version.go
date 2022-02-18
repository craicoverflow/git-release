package version

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-git/go-git/v5/plumbing"
)

type Version struct {
	Prefix     string
	Major      int
	Minor      int
	Patch      int
	Prerelease *Prerelease
	Meta       string
}

type Prerelease struct {
	Type    string
	Version int
}

func (p *Prerelease) String() string {
	return fmt.Sprintf("%v%v", p.Type, p.Version)
}

func (v *Version) String() string {
	versionFmt := fmt.Sprintf("%v%v.%v.%v", v.Prefix, v.Major, v.Minor, v.Patch)
	if v.Prerelease != nil {
		versionFmt += "-" + v.Prerelease.String()
	}
	if v.Meta != "" {
		versionFmt += "+" + v.Meta
	}

	return versionFmt
}

func Parse(ref *plumbing.Reference) (*Version, error) {
	tag := ref.Name().Short()

	var version Version
	parts := strings.Split(tag, ".")

	if strings.HasPrefix(parts[0], "v") {
		version.Prefix = "v"
		parts[0] = strings.Trim(parts[0], "v")
	}

	maj, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, errors.New("could not parse major version " + parts[0])
	}
	version.Major = maj

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, errors.New("could not parse minor version " + parts[1])
	}
	version.Minor = minor

	if i := strings.Index(parts[2], "-"); i > -1 {
		patchS := parts[2][:i]
		patch, err := strconv.Atoi(patchS)
		if err != nil {
			return nil, errors.New("could not parse patch version " + patchS)
		}
		version.Patch = patch
		prerel := parts[2][i+1:]
		if i := strings.Index(prerel, "+"); i > -1 {
			meta := prerel[i+1:]
			version.Meta = meta
			prerel = prerel[:i]
		}
		preRelease, err := parsePrerelease(prerel)
		if err != nil {
			return nil, errors.New("could not parse pre-release " + prerel)
		}
		version.Prerelease = preRelease
	}

	fmt.Println(version.String())

	return &version, nil
}

func parsePrerelease(s string) (*Prerelease, error) {
	var l, n []rune
	for _, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
			l = append(l, r)
		case r >= 'a' && r <= 'z':
			l = append(l, r)
		case r >= '0' && r <= '9':
			n = append(n, r)
		}
	}

	parsedVersion, err := strconv.Atoi(string(n))
	if err != nil {
		return nil, err
	}

	return &Prerelease{
		Type:    string(l),
		Version: parsedVersion,
	}, nil
}
