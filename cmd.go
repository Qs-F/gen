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
	w := flag.Bool("w", false, "write out")
	info := flag.Bool("i", false, "show info")
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
	if *info {
		fmt.Fprintf(os.Stderr, ">>> Project file list\n")
		fmt.Println(list)
	}

	mdExp := mde.New("layout", "content", "__content__", list)

	out, err := gen.Expand(g.BasePath, g.SrcPath, g.DstPath, list, mdExp)
	if err != nil {
		logrus.Error(err)
		return
	}
	if *info {
		fmt.Fprintf(os.Stderr, ">>> Files outputed\n")
		fmt.Println(out)
	}

	if !*w {
		return
	}
	err = gen.Write(g.BasePath, out)
	if err != nil {
		logrus.Error(err)
		return
	}
}
