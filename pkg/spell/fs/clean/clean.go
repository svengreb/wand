// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package clean provides a spell incantation to remove directories in a filesystem.
// It implements spell.GoCode and can be used without a cast.Caster.
package clean

import (
	"fmt"
	"os"
	"path/filepath"

	glFS "github.com/svengreb/golib/pkg/io/fs"
	glFilePath "github.com/svengreb/golib/pkg/io/fs/filepath"

	"github.com/svengreb/wand/pkg/app"
	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/spell"
)

// Spell is a spell incantation to remove directories in a filesystem.
// It is of kind spell.KindGoCode and can be used without a cast.Caster.
type Spell struct {
	ac   app.Config
	proj project.Metadata
	opts *Options
}

// Formula returns the spell incantation formula.
// Note that Spell implements spell.GoCode so this method always returns an empty slice!
func (s *Spell) Formula() []string {
	return []string{}
}

// Kind returns the spell incantation kind.
func (s *Spell) Kind() spell.Kind {
	return spell.KindGoCode
}

// Options returns the spell incantation options.
func (s *Spell) Options() interface{} {
	return *s.opts
}

// Clean removes the configured paths.
// It returns an error of type *spell.ErrGoCode for any error that occurs during the execution of the Go code.
func (s *Spell) Clean() ([]string, error) {
	var cleaned []string

	for _, p := range s.opts.paths {
		pAbs := filepath.Join(s.proj.Options().RootDirPathAbs, p)

		if s.opts.limitToAppOutputDir {
			appDir := filepath.Join(s.proj.Options().RootDirPathAbs, s.ac.BaseOutputDir)
			pAbs = filepath.Join(s.proj.Options().RootDirPathAbs, p)

			isSubDir, subDirErr := glFilePath.IsSubDir(appDir, pAbs, false)
			if subDirErr != nil {
				return cleaned, &spell.ErrGoCode{
					Err:  fmt.Errorf("check if %q is a subdirectory of %q: %w", pAbs, appDir, subDirErr),
					Kind: spell.ErrExec,
				}
			}
			if !isSubDir {
				return cleaned, &spell.ErrGoCode{
					Err:  fmt.Errorf("%q is not a subdirectory of %q", pAbs, appDir),
					Kind: spell.ErrExec,
				}
			}
		}

		nodeExists, fsErr := glFS.FileExists(pAbs)
		if fsErr != nil {
			return cleaned, &spell.ErrGoCode{
				Err:  fmt.Errorf("check if %q exists: %w", pAbs, fsErr),
				Kind: spell.ErrExec,
			}
		}
		if nodeExists {
			if err := os.RemoveAll(pAbs); err != nil {
				return cleaned, &spell.ErrGoCode{
					Err:  fmt.Errorf("remove path %q: %w", pAbs, err),
					Kind: spell.ErrExec,
				}
			}
			cleaned = append(cleaned, p)
		}
	}

	return cleaned, nil
}

// New creates a new spell incantation to remove the configured filesystem paths, e.g. output data like artifacts and
// reports from previous development, test, production and distribution builds.
//nolint:gocritic // The app.Config struct is passed as value by design to ensure immutability.
func New(proj project.Metadata, ac app.Config, opts ...Option) (*Spell, error) {
	opt, optErr := NewOptions(opts...)
	if optErr != nil {
		return nil, optErr
	}

	return &Spell{ac: ac, proj: proj, opts: opt}, nil
}
