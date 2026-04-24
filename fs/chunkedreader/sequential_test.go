package chunkedreader

import (
	"testing"

	_ "github.com/Somebodybro/rclone-WA-proto/backend/local"
	"github.com/Somebodybro/rclone-WA-proto/fstest/mockobject"
)

func TestSequential(t *testing.T) {
	content := makeContent(t, 1024)

	for _, mode := range mockobject.SeekModes {
		t.Run(mode.String(), testRead(content, mode, 0))
	}
}

func TestSequentialErrorAfterClose(t *testing.T) {
	testErrorAfterClose(t, 0)
}
