// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package plugin

import (
	"github.com/hashicorp/go-plugin"
)

// Starts the serving of a Mattermost plugin over net/rpc. gRPC is not yet supported.
//
// Call this when your plugin is ready to start.
func ClientMain(pluginImplementation interface{}) {
	if impl, ok := pluginImplementation.(interface {
		SetAPI(api API)
	}); !ok {
		panic("Plugin implementation given must embed plugin.MattermostPlugin")
	} else {
		impl.SetAPI(nil)
	}

	pluginMap := map[string]plugin.Plugin{
		"hooks": &hooksPlugin{hooks: pluginImplementation},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshake,
		Plugins:         pluginMap,
	})
}

type MattermostPlugin struct {
	// API exposes the plugin api, and becomes available just prior to the OnActive hook.
	API API
}

// SetAPI persists the given API interface to the plugin. It is invoked just prior to the
// OnActivate hook, exposing the API for use by the plugin.
func (p *MattermostPlugin) SetAPI(api API) {
	p.API = api
}