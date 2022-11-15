// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

package elder

import (
	"context"
	"errors"

	"github.com/svengreb/wand"
)

// UnwrapCtx is a helper function to unwrap a elder wand from context.
func UnwrapCtx(ctx context.Context) (*Elder, error) {
	if val, ok := ctx.Value(wand.GetCtxKey()).(*Elder); ok {
		return val, nil
	}

	return nil, errors.New("no elder wand in context")
}
