package cueutil

import (
	"io"
	"io/fs"
	"path/filepath"

	"cuelang.org/go/cue/load"
)

// An Overlay value may be assigned to the Overlay field of
// a [load.Config].
type Overlay map[string]load.Source

// AddFSPaths adds files from parts of fsys, selected by
// the specified paths, to the Overlay.
// AddFSPaths walks the parts file system, reading each file into a byte slice,
// adding it to the Overlay as a [load.Source] using [load.FromBytes].
// It returns any error occured during the call to [fs.Walkdir].
func (o Overlay) AddFSPaths(fsys fs.FS, paths ...string) error {
	for _, path := range paths {
		err := o.addPath(fsys, path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o Overlay) addPath(fsys fs.FS, path string) error {

	fs.WalkDir(fsys, path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fs.SkipDir
		}
		if d.IsDir() {
			return nil
		}
		if d.Type()&fs.ModeSymlink != 0 {
			return nil
		}
		r, err := fsys.Open(path)
		if err != nil {
			return err
		}
		b, err := io.ReadAll(r)
		r.Close()
		if err != nil {
			return err
		}
		abs, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		o[abs] = load.FromBytes(b)
		return nil
	})
	return nil
}

// AddString adds a file with the specified string content
// to the Overlay. It calls [filepath.Abs] on the filename before
// storing the string into the Overlay, and returns an error
// if that call fails.
func (o Overlay) AddString(filename, contents string) error {
	return o.addSource(filename, load.FromString(contents))
}

// AddBytes adds a file with the specified []byte content
// to the Overlay. It calls [filepath.Abs] on the filename before
// storing the byte slice into the Overlay, and returns an error
// if that call fails.
func (o Overlay) AddBytes(filename string, contents []byte) error {
	return o.addSource(filename, load.FromBytes(contents))
}

func (o Overlay) addSource(filename string, s load.Source) error {
	abs, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	o[abs] = s
	return nil
}
