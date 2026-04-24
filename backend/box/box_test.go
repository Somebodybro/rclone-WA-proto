// Test Box filesystem interface
package box_test

import (
	"testing"

	"github.com/Somebodybro/rclone-WA-proto/backend/box"
	"github.com/Somebodybro/rclone-WA-proto/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestBox:",
		NilObject:  (*box.Object)(nil),
	})
}
