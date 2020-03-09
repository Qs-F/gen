package main

import (
	"os"

	"github.com/Qs-F/gen/lib/gen"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
		return
	}
	g := gen.New(wd, "")
	list, err := g.List(NewMarkdownLoader())
	if err != nil {
		logrus.Error(err)
		return
	}
	v, err := g.ResolveKey(list, "testdata/md/basic.md")
	if err != nil {
		logrus.Error(err)
		return
	}
	spew.Dump(v)
}
