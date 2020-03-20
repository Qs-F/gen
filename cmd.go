package main

import (
	"flag"
	"os"

	mde "github.com/Qs-F/gen/lib/expander/markdown"
	"github.com/Qs-F/gen/lib/gen"
	"github.com/Qs-F/gen/lib/loader/html"
	mdl "github.com/Qs-F/gen/lib/loader/markdown"
	"github.com/sirupsen/logrus"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
		return
	}
	base := flag.String("base", wd, "base path")
	src := flag.String("src", wd, "src path")
	dst := flag.String("dst", wd, "dst path")
	flag.Parse()

	g := gen.New("", "", "")
	err = g.Set(*base, *src, *dst)
	if err != nil {
		logrus.Error(err)
		return
	}

	mdLd := mdl.New()
	htmlLd := html.New("content")

	list, err := gen.GenList(g.BasePath, mdLd, htmlLd)
	if err != nil {
		logrus.Error(err)
		return
	}

	mdExp := mde.New("layout", "content", "__content__", list)

	err = list.Expand(g.BasePath, g.SrcPath, g.DstPath, mdExp)
	if err != nil {
		logrus.Error(err)
		return
	}
}
