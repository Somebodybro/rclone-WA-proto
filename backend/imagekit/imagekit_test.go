package imagekit

import (
	"testing"

	"github.com/Somebodybro/rclone-WA-proto/fstest"
	"github.com/Somebodybro/rclone-WA-proto/fstest/fstests"
)

func TestIntegration(t *testing.T) {
	debug := true
	fstest.Verbose = &debug
	fstests.Run(t, &fstests.Opt{
		RemoteName:      "TestImageKit:",
		NilObject:       (*Object)(nil),
		SkipFsCheckWrap: true,
	})
}
