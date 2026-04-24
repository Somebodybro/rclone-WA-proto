// Package cmdtest creates a testable interface to rclone main
//
// The interface is used to perform end-to-end test of
// commands, flags, environment variables etc.
package cmdtest

// The rest of this file is a 1:1 copy from rclone.go

import (
	_ "github.com/Somebodybro/rclone-WA-proto/backend/all" // import all backends
	"github.com/Somebodybro/rclone-WA-proto/cmd"
	_ "github.com/Somebodybro/rclone-WA-proto/cmd/all"    // import all commands
	_ "github.com/Somebodybro/rclone-WA-proto/lib/plugin" // import plugins
)

func main() {
	cmd.Main()
}
