// Package pkg provides ways to package programs in some linux packaging
// formats.
package pkg

import (
	"context"
	"io"
)

// Packager represents any packager implementation
type Packager interface {
	Package(ctx context.Context, info Info, w io.Writer) error
}

// Info contains information about the package
type Info struct {
	Name        string            `yaml:"name,omitempty"`
	Arch        string            `yaml:"arch,omitempty"`
	Version     string            `yaml:"version,omitempty"`
	Section     string            `yaml:"section,omitempty"`
	Priority    string            `yaml:"priority,omitempty"`
	Replaces    string            `yaml:"replaces,omitempty"`
	Provides    string            `yaml:"provides,omitempty"`
	Depends     []string          `yaml:"depends,omitempty"`
	Conflicts   []string          `yaml:"conflicts,omitempty"`
	Maintainer  string            `yaml:"maintainer,omitempty"`
	Description string            `yaml:"description,omitempty"`
	Vendor      string            `yaml:"vendor,omitempty"`
	Homepage    string            `yaml:"homepage,omitempty"`
	Files       map[string]string `yaml:"files,omitempty"`
}
