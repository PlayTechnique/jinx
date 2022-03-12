//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"jinx/src/jinkiesengine"
)

func getContainerConfigPath() string {
	return containerConfigPath
}

func initializeCobraJinkies() jinkiesengine.Jinkies {
	wire.Build(jinkiesengine.NewJinkies, hydrateFromConfig, getContainerConfigPath)
	return jinkiesengine.Jinkies{}
}
