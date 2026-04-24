// Test Files filesystem interface
package filescom_test

import (
	"testing"

	"github.com/Somebodybro/rclone-WA-proto/backend/filescom"
	"github.com/Somebodybro/rclone-WA-proto/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestFilesCom:",
		NilObject:  (*filescom.Object)(nil),
	})
}
