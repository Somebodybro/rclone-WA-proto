// Test pixeldrain filesystem interface
package pixeldrain_test

import (
	"testing"

	"github.com/Somebodybro/rclone-WA-proto/backend/pixeldrain"
	"github.com/Somebodybro/rclone-WA-proto/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName:      "TestPixeldrain:",
		NilObject:       (*pixeldrain.Object)(nil),
		SkipInvalidUTF8: true, // Pixeldrain throws an error on invalid utf-8
	})
}
