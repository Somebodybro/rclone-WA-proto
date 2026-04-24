//go:build noselfupdate

package selfupdate

import (
	"github.com/Somebodybro/rclone-WA-proto/lib/buildinfo"
)

func init() {
	buildinfo.Tags = append(buildinfo.Tags, "noselfupdate")
}
