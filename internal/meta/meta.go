package meta

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

const (
	App          = "cartman"
	Short        = "Lightweight Certificate Authority"
	VersionMajor = 25
	VersionMinor = 5
	VersionPatch = 0
)

var Commit string

func VersionString() string {
	return fmt.Sprintf("%s %02d.%02d.%d %s/%s", App, VersionMajor, VersionMinor, VersionPatch, runtime.GOOS, runtime.GOARCH)
}

func init() {

	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				Commit = setting.Value
			}
		}
	}
}
