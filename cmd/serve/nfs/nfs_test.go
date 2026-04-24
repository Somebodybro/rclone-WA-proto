//go:build unix

// The serving is tested in cmd/nfsmount - here we test anything else
package nfs

import (
	"testing"

	_ "github.com/Somebodybro/rclone-WA-proto/backend/local"
	"github.com/Somebodybro/rclone-WA-proto/cmd/serve/servetest"
	"github.com/Somebodybro/rclone-WA-proto/fs/rc"
)

func TestRc(t *testing.T) {
	servetest.TestRc(t, rc.Params{
		"type":           "nfs",
		"vfs_cache_mode": "off",
	})
}
