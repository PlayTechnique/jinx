package jinkiesengine

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreateDirectories(t *testing.T) {
	testDir, _ := os.MkdirTemp("", "")

	directories := createDirectories(testDir)

	for _, dir := range directories {
		assert.DirExists(t, dir)
	}
}
