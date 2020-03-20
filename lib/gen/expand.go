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

func (list List) Expand(basePath, srcPath, dstPath string, expanders ...Expander) error {
	dst := filepath.Join(basePath, dstPath)
	exps := make(map[string]Expander)
	for _, expander := range expanders {
		from, _ := expander.Ext()
		exps[from] = expander
	}

	for file := range list {
		// check file is in src
		if !strings.HasPrefix(file, srcPath) {
			continue
		}

		// check expander is available for file
		cont := false
		var expander Expander
		for ext, exp := range exps {
			if ext == filepath.Ext(file) {
				cont = true
				expander = exp
				break
			}
		}
		if !cont {
			continue
		}

		// get file content
		content, perm, err := open(filepath.Join(basePath, srcPath, file))
		if err != nil {
			return err
		}

		// expand file
		to, b, err := ExpandEach(list, expander, file, content)
		if err != nil {
			return err
		}

		// write
		err = ioutil.WriteFile(filepath.Join(dst, to), b, perm)
		if err != nil {
			return err
		}
	}

	return nil
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

func open(path string) ([]byte, os.FileMode, error) {
	f, err := os.Open(filepath.Join(path))
	if err != nil {
		return nil, os.ModePerm, err
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, os.ModePerm, err
	}
	perm := fi.Mode().Perm()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, os.ModePerm, err
	}
	f.Close()
	return content, perm, nil
}
