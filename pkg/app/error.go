// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package app

import "fmt"

// ErrConfigNotFound indicates that an application configuration has not been found.
type ErrConfigNotFound struct {
	Name string
}

func (e *ErrConfigNotFound) Error() string {
	return fmt.Sprintf("no configuration found for application with name %q", e.Name)
}
