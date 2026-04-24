// Test Mailru filesystem interface
package mailru_test

import (
	"testing"

	"github.com/Somebodybro/rclone-WA-proto/backend/mailru"
	"github.com/Somebodybro/rclone-WA-proto/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName:               "TestMailru:",
		NilObject:                (*mailru.Object)(nil),
		SkipBadWindowsCharacters: true,
	})
}
