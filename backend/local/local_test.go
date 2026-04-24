// Test Local filesystem interface
package local_test

import (
	"testing"

	"github.com/Somebodybro/rclone-WA-proto/backend/local"
	"github.com/Somebodybro/rclone-WA-proto/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName:  "",
		NilObject:   (*local.Object)(nil),
		QuickTestOK: true,
	})
}
