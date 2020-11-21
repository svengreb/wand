// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package error provides types and utilities to handle errors.
package error

// ErrString is a string type for implementing constant errors.
type ErrString string

func (e ErrString) Error() string {
	return string(e)
}
