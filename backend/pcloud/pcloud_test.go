// Test Pcloud filesystem interface
package pcloud_test

import (
	"testing"

	"github.com/Somebodybro/rclone-WA-proto/backend/pcloud"
	"github.com/Somebodybro/rclone-WA-proto/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestPcloud:",
		NilObject:  (*pcloud.Object)(nil),
	})
}
