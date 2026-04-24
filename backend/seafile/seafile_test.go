// Test Seafile filesystem interface
package seafile_test

import (
	"testing"

	"github.com/Somebodybro/rclone-WA-proto/backend/seafile"
	"github.com/Somebodybro/rclone-WA-proto/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestSeafile:",
		NilObject:  (*seafile.Object)(nil),
	})
}
