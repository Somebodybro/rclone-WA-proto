//go:build !plan9 && !solaris

package iclouddrive_test

import (
	"testing"

	"github.com/Somebodybro/rclone-WA-proto/backend/iclouddrive"
	"github.com/Somebodybro/rclone-WA-proto/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestICloudDrive:",
		NilObject:  (*iclouddrive.Object)(nil),
	})
}
