// Command outsystems-cli is a command-line interface for the OutSystems
// LifeTime REST API v2.
package main

import "github.com/jinjie/outsystems-cli/cmd"

// Build metadata, set at release time via -ldflags by GoReleaser:
//
//	-ldflags "-X main.version={{.Version}} -X main.commit={{.Commit}} -X main.date={{.Date}}"
//
// For plain `go build` / `go run` these remain at their zero values and the
// CLI reports the "dev" defaults.
var (
	version = ""
	commit  = ""
	date    = ""
)

func main() {
	cmd.SetVersionInfo(version, commit, date)
	cmd.Execute()
}
