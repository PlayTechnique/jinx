package jenkins

import (
	"os"
	"regexp"
	"testing"
)
import jinxtypes "jinx/types"

func TestPlugins(t *testing.T) {
	jinxRuntime := jinxtypes.JinxGlobalRuntime{
		ContainerName: "roflcopter",
		PullImages:    false,
	}

	topLevelDir, _ := os.MkdirTemp("", "")
	defer os.RemoveAll(topLevelDir)

	outputFormat := "notvalid.txt"

	err := Plugins(jinxRuntime, topLevelDir, outputFormat)
	match, _ := regexp.MatchString("Valid output formats are .*", err.Error())
	if !match {
		t.Errorf("Invalid outputformat generated unexpected error %v", err)
	}

}
