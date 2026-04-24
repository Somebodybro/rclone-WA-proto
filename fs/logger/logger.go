// Package logger implements testing for the sync (and bisync) logger
package logger

import (
	_ "github.com/Somebodybro/rclone-WA-proto/backend/all" // import all backends
	"github.com/Somebodybro/rclone-WA-proto/cmd"
	_ "github.com/Somebodybro/rclone-WA-proto/cmd/all"    // import all commands
	_ "github.com/Somebodybro/rclone-WA-proto/lib/plugin" // import plugins
)

// Main enables the testscript package. See:
// https://bitfieldconsulting.com/golang/cli-testing
// https://pkg.go.dev/github.com/rogpeppe/go-internal@v1.11.0/testscript
func Main() {
	cmd.Main()
}
