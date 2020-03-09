package main

import (
	"errors"
)

const (
	// ImportIdent is import directive identifier
	ImportIdent = "import"
	// Root is the List key for root, starting point of resolving
	Root = "_"
)

var (
	// ErrRootNotFound is returned when root is not found
	ErrRootNotFound = errors.New("Root is not found")
	// ErrVariableNotFound is returned when variable is not found
	ErrVariableNotFound = errors.New("No such variable ident")
)

// List is all Variable list.
// map key is the key to access Variable
type List map[string]Variables

// Resolve resolves th given list into Variables aling with the key "_"
// Internally, this uses ResolveKey with key = Root
func Resolve(list List) (Variables, error) {
	return ResolveKey(list, Root)
}

// ResolveKey resolves the given list into Variables along with the given key as root.
func ResolveKey(list List, key string) (Variables, error) {
	v, err := list.Get(key)
	if err != nil {
		return nil, ErrRootNotFound
	}
	tree := list.GenTree(v.ToNode(key, nil))
	return tree.Reduce(), nil
}

// Get provides the way to access the list with key.
// On `map`, v, ok := map[xxx] is used, but error is needed for lib design.
// Get returns error when the key is not found.
func (list List) Get(key string) (Variables, error) {
	v, ok := list[key]
	if !ok {
		return nil, ErrVariableNotFound
	}
	return v, nil
}

// ContainsImports checks Variable imports or not.
func (v Variables) ContainsImports() bool {
	if v == nil {
		return false
	}

	im, ok := v[ImportIdent]
	if !ok {
		return false
	}

	if _, ok := im.([]string); !ok {
		return false
	}

	return true
}

// ToImports resolve Variable into import keys.
// If no deps, empty string slice is returned.
func (v Variables) ToImports() []string {
	if !v.ContainsImports() {
		return []string{}
	}

	// by v.ContainsImports, it is gualanteed to have ImportDelim key in v
	ret, _ := v[ImportIdent].([]string)
	return ret
}

// Node is the deps tree node.
type Node struct {
	Key     string
	Content Variables
	Deps    []string

	parent   *Node
	children []*Node
	resolved bool
}

// ToNode converts Variable to Node.
// if node is root, then parent must be nil
func (v Variables) ToNode(key string, parent *Node) *Node {
	return &Node{
		Key:      key,
		Content:  v,
		Deps:     v.ToImports(),
		parent:   parent,
		children: []*Node{},
		resolved: false,
	}
}

// resolve resolves 1 incremental depth for node.
// resolve will changes node.children and node.resolved.
// Each child have n as parent.
func (n *Node) link(list List) {
	// if already resolved, then no more resolve
	if n.resolved {
		return
	}

	for _, dep := range n.Deps {
		v, err := list.Get(dep)

		// if dep is not found in list, then continue
		if err != nil {
			continue
		}

		n.children = append(n.children, v.ToNode(dep, n))
	}
	n.resolved = true
}

type imported map[string]bool

// Depth is nodes slice for the same depth.
type Depth struct {
	Depth int
	Nodes []*Node
}

// resolve returns next depth.
// resolve must be called from root to leaves, left to right
func resolve(list List, im imported, depth *Depth) *Depth {
	nodes := []*Node{}
	for _, node := range depth.Nodes {
		// if node is already imported, then continue
		if b, ok := im[node.Key]; ok && b {
			continue
		}

		if !node.resolved {
			node.link(list)
		}

		nodes = append(nodes, node.children...)
		im[node.Key] = true
	}

	if len(nodes) == 0 {
		return nil
	}

	return &Depth{
		Depth: depth.Depth + 1,
		Nodes: nodes,
	}
}

// Tree is tree of sum of depth.
type Tree []*Depth

func (list List) GenTree(root *Node) Tree {
	im := make(imported)
	zeroDepth := &Depth{
		Depth: 0,
		Nodes: []*Node{root},
	}

	ret := Tree{zeroDepth}

	for {
		depth := resolve(list, im, zeroDepth)
		if depth == nil {
			break
		}
		ret = append(ret, depth)
	}

	return ret
}

func (tree Tree) Reduce() Variables {
	variables := make(Variables)

	for _, depth := range tree {
		for _, node := range depth.Nodes {
			for k, v := range node.Content {
				if k == ImportIdent {
					continue
				}
				if _, ok := variables[k]; ok {
					continue
				}
				variables[k] = v
			}
		}
	}

	return variables
}
