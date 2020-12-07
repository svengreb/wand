// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package app

import (
	"fmt"
)

// Store is a storage that provides methods to record application configurations.
type Store interface {
	// Add adds a application configuration.
	Add(*Config)

	// Get returns the application configuration for the given name or nil along with an error when not stored.
	Get(string) (*Config, error)
}

// appStore is a storage for application configurations.
type appStore struct {
	data map[string]*Config
}

// Add adds an application configuration.
func (s *appStore) Add(ac *Config) {
	s.data[ac.Name] = ac
}

// Get returns an application configuration.
// It returns an error of type *app.ErrApp when there is no such configuration in the store along with an empty
// application configuration.
func (s *appStore) Get(appName string) (*Config, error) {
	ac, ok := s.data[appName]
	if !ok {
		return nil, &ErrApp{
			Err:  fmt.Errorf("application name %q", appName),
			Kind: ErrNoSuchConfig,
		}
	}
	return ac, nil
}

// NewStore creates a new store for application configurations.
func NewStore() Store {
	return &appStore{
		data: make(map[string]*Config),
	}
}
