package main

import (
	"os"

	"github.com/Qs-F/gen/lib/gen"
	"github.com/Qs-F/gen/lib/loader/markdown"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
		return
	}

	list, err := gen.GenList(wd, markdown.New())
	if err != nil {
		logrus.Error(err)
		return
	}

	v, err := gen.ResolveKey(list, "testdata/basic/md/basic.md")
	if err != nil {
		logrus.Error(err)
		return
	}
	spew.Dump(v)
}
