package iox_test

import (
	"fmt"
	"github.com/gophero/gotools/iox"
	"github.com/gophero/gotools/testx"
	"os"
	"path/filepath"
	"testing"
)

func TestPathExists(t *testing.T) {
	lg := testx.Wrap(t)

	lg.Case("give an exists path")
	path := "/Users/sam/workspace/mine/gotools"
	lg.Require(iox.Path.PathExists(path), "given path should exist")

	lg.Case("give an none exists path")
	path = "/Users/haha"
	lg.Require(!iox.Path.PathExists(path), "given path should not exist")
}

func TestExecPath(t *testing.T) {
	execpath, err := os.Executable() // 获得程序路径
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(execpath)
	fmt.Println(dir)

	s, _ := os.Getwd()
	println(s)

	println(iox.Path.ExecPath())
	println(iox.Path.CurrentPath())
	println(iox.Path.ProjectPath())
}
