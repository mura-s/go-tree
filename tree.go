package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Options store options for the tree command.
type Options struct {
	allFiles     bool
	maxDeepLevel int

	out io.Writer // writes the result to out.
}

// Tree represents a directory tree.
type Tree struct {
	root *node
}

// node represents a file or a directory in a tree.
type node struct {
	baseName string
	isDir    bool
	subNodes []*node
}

// MakeTree traverses the directory tree rooted at the given path and returns the tree.
func MakeTree(rootPath string, opts *Options) (*Tree, error) {
	rootFile, err := os.Open(rootPath)
	if err != nil {
		return nil, err
	}
	defer rootFile.Close()

	rootFI, err := rootFile.Stat()
	if err != nil {
		return nil, err
	}
	if !rootFI.IsDir() {
		return nil, errors.New("root path must be a directory path")
	}

	root, err := traverse(rootPath, 0, opts)
	if err != nil {
		return nil, err
	}
	return &Tree{root: root}, nil
}

func traverse(path string, depth int, opts *Options) (*node, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	n := &node{
		baseName: fi.Name(),
		isDir:    fi.IsDir(),
	}
	// return if the node is a file
	if !n.isDir {
		return n, nil
	}
	// maxDeepLevel option
	if depth >= opts.maxDeepLevel {
		return n, nil
	}

	// traverse sub directories
	names, err := file.Readdirnames(0)
	if err != nil {
		return nil, err
	}
	sort.Strings(names)

	for _, name := range names {
		// allFiles option
		if !opts.allFiles && strings.HasPrefix(name, ".") {
			continue
		}

		subNode, err := traverse(filepath.Join(path, name), depth+1, opts)
		if err != nil {
			return nil, err
		}
		n.subNodes = append(n.subNodes, subNode)
	}

	return n, nil
}

// Print the structure of the tree.
func (t *Tree) Print(opts *Options) {
	t.root.print("", "", opts)
}

func (n *node) print(indent, prefix string, opts *Options) {
	fmt.Fprintf(opts.out, "%s%s\n", prefix, n.baseName)

	for i, subNode := range n.subNodes {
		if i == len(n.subNodes)-1 {
			subNode.print(indent+"    ", indent+"└── ", opts)
		} else {
			subNode.print(indent+"│   ", indent+"├── ", opts)
		}
	}
}
