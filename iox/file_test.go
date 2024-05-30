package iox_test

import (
	"github.com/nicktaobo/go_tool/iox"
	"github.com/nicktaobo/go_tool/testx"
	"testing"
)

func TestExistsFile(t *testing.T) {
	lg := testx.Wrap(t)

	lg.Case("give an existing file")
	f := "/Users/sam/workspace/mine/gotools/io/file_test.go"
	lg.Require(iox.File.Exists(f), "should exist")

	lg.Case("give an existing dir, but is not a file")
	f = "/Users/sam/workspace/mine/gotools/io/"
	lg.Require(!iox.File.Exists(f), "should not exist")
}
