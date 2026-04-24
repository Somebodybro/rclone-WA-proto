package filen

import (
	"testing"

	"github.com/Somebodybro/rclone-WA-proto/fstest/fstests"
)

func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestFilen:",
		NilObject:  (*Object)(nil),
	})
}
