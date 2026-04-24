package filelu_test

import (
	"testing"

	"github.com/Somebodybro/rclone-WA-proto/fstest/fstests"
)

// TestIntegration runs integration tests for the FileLu backend
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName:      "TestFileLu:",
		NilObject:       nil,
		SkipInvalidUTF8: true,
	})
}
