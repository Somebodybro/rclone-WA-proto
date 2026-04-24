// Test Sia filesystem interface
package sia_test

import (
	"testing"

	"github.com/Somebodybro/rclone-WA-proto/backend/sia"

	"github.com/Somebodybro/rclone-WA-proto/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestSia:",
		NilObject:  (*sia.Object)(nil),
	})
}
