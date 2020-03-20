package gen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Expander interface {
	Ext() (from, to string)
	Expand(p []byte, v Variables) ([]byte, error)
}

func ExpandEach(list List, exp Expander, file string, content []byte) (dst string, w []byte, err error) {
	v, err := ResolveKey(list, file)
	if err != nil {
		return "", nil, err
	}

	w, err = exp.Expand(content, v)
	if err != nil {
		return "", nil, err
	}

	_, ext := exp.Ext()
	dst = strings.TrimRight(file, filepath.Ext(file)) + ext
	return dst, w, nil
}
