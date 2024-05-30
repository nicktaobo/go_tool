package iox_test

import (
	"fmt"
	"github.com/nicktaobo/go_tool/iox"
	"strings"
	"testing"
)

func TestWalkAllFiles(t *testing.T) {
	fs := iox.WalkDir("/Users/hank/Downloads")
	fmt.Println(strings.Join(fs, "\n"))
}
