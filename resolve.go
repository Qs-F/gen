package main

import (
	"errors"
)

const (
	ImportDelim = "import"
	Root        = "_"
)

var (
	ErrRootNotFound = errors.New("Root is not found")
)

func (fm Variables) ContainsImports() bool {
	if fm == nil {
		return false
	}

	v, ok := fm[ImportDelim]
	if !ok {
		return false
	}

	if _, ok := v.([]string); !ok {
		return false
	}

	return true
}

type List map[string]Variables

type Node struct {
	Parent  *Node
	Childen []*Node

	Content Variables
}

func (l List) GenTree() (*Node, error) {
	if _, ok := l[Root]; !ok {
		return nil, ErrRootNotFound
	}
	return nil, nil
}

func (n *Node) Link() (Variables, error) {
	return nil, nil
}

func Resolve(m map[string]Variables) (Variables, error) {
	return nil, nil
}
