//go:build linux

package mount

import (
	"testing"

	"github.com/Somebodybro/rclone-WA-proto/vfs/vfscommon"
	"github.com/Somebodybro/rclone-WA-proto/vfs/vfstest"
)

func TestMount(t *testing.T) {
	vfstest.RunTests(t, false, vfscommon.CacheModeOff, true, mount)
}
