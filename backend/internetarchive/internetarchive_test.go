// Test internetarchive filesystem interface
package internetarchive_test

import (
	"testing"

	"github.com/Somebodybro/rclone-WA-proto/backend/internetarchive"
	"github.com/Somebodybro/rclone-WA-proto/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestIA:lesmi-rclone-test/",
		NilObject:  (*internetarchive.Object)(nil),
	})
}
