package gen

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

	// recurrent import
	for k, val := range v {
		value, _ := val.(Variables) // since Variables is recurrent system, it must be always ok
		if !value.ContainsImports() {
			continue
		}
		// found named import
		v[k] = list.GenTree(value.ToNode("")).Reduce()
	}

	tree := list.GenTree(v.ToNode(key))
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

	switch im.(type) {
	case []string:
		return true
	case []interface{}:
		return true
	default:
		return false
	}
}

// ToImports resolve Variable into import keys.
// If no deps, empty string slice is returned.
func (v Variables) ToImports() []string {
	if !v.ContainsImports() {
		return []string{}
	}

	// by v.ContainsImports, it is gualanteed to have ImportDelim key in v
	switch t := v[ImportIdent].(type) {
	case []interface{}:
		ret := make([]string, len(t), cap(t))
		for i, v := range t {
			ret[i] = v.(string)
		}
		return ret
	case []string:
		return t
	default:
		return []string{}
	}
}
