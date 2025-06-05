package version

import "runtime/debug"

var (
	Version   = "dev"
	BuildTime = "unknown"
	Commit    = "none"
)

func GetVersionInfo() string {
	info := "Version: " + Version + "\nCommit: " + Commit + "\nBuild Time: " + BuildTime

	if bi, ok := debug.ReadBuildInfo(); ok {
		info += "\nGo Version: " + bi.GoVersion
	}

	return info
}
