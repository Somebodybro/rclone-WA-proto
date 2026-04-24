// Test filesystem interface
package premiumizeme_test

import (
	"testing"

	"github.com/Somebodybro/rclone-WA-proto/backend/premiumizeme"
	"github.com/Somebodybro/rclone-WA-proto/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestPremiumizeMe:",
		NilObject:  (*premiumizeme.Object)(nil),
	})
}
