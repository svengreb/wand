// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package wand

import (
	"context"

	"github.com/svengreb/wand/pkg/app"
	"github.com/svengreb/wand/pkg/project"
)

// Wand manages a project and its applications and stores their metadata.
// Applications are registered using a unique name and the stored metadata can be received based on this name.
type Wand interface {
	// GetAppConfig returns an application configuration.
	GetAppConfig(appName string) (app.Config, error)

	// GetProjectMetadata returns the project metadata.
	GetProjectMetadata() project.Metadata

	// RegisterApp registers a new application.
	RegisterApp(name, displayName, pathRel string) error
}

// ctxKey is the context key used to wrap a Wand.
type ctxKey struct{}

// GetCtxKey returns the key used to wrap a Wand.
func GetCtxKey() interface{} {
	return ctxKey{}
}

// WrapCtx wraps the given Wand into the parent context.
// Use GetCtxKey to receive the key used to wrap the Wand.
func WrapCtx(parentCtx context.Context, wand Wand) context.Context {
	if parentCtx == nil {
		parentCtx = context.Background()
	}
	return context.WithValue(parentCtx, ctxKey{}, wand)
}
