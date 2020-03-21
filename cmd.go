package main

import (
	"flag"
	"fmt"
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
	fmt.Println(list)

	mdExp := mde.New("layout", "content", "__content__", list)

	out, err := gen.Expand(g.BasePath, g.SrcPath, g.DstPath, list, mdExp)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(out)

	err = gen.Write(g.BasePath, out)
	if err != nil {
		logrus.Error(err)
		return
	}
}
