package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

// Variables is the type expressing front matters
type Variables map[string]interface{}

type Gen struct {
	BasePath   string
	LayoutPath string
}
