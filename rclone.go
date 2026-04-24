// Sync files and directories to and from local and remote object stores
//
// Nick Craig-Wood <nick@craig-wood.com>
package main

import (
	_ "github.com/Somebodybro/rclone-WA-proto/backend/all" // import all backends
	"github.com/Somebodybro/rclone-WA-proto/cmd"
	_ "github.com/Somebodybro/rclone-WA-proto/cmd/all"    // import all commands
	_ "github.com/Somebodybro/rclone-WA-proto/lib/plugin" // import plugins
)

func main() {
	cmd.Main()
}
