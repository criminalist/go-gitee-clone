// +build bindata

// Copyright 2016 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package public

import (
	"github.com/go-macaron/bindata"
	"gopkg.in/macaron.v1"
)

// Static implements the macaron static handler for serving assets.
func Static(opts *Options) macaron.Handler {
	return macaron.Static(
		opts.Directory,
		macaron.StaticOptions{
			SkipLogging: opts.SkipLogging,
			FileSystem: bindata.Static(bindata.Options{
				Asset:      Asset,
				AssetDir:   AssetDir,
				AssetInfo:  AssetInfo,
				AssetNames: AssetNames,
				Prefix:     "",
			}),
		},
	)
}
