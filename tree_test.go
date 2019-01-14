package main

import (
	"bytes"
	"math"
	"testing"
)

var (
	opts = &Options{
		allFiles:     false,
		maxDeepLevel: math.MaxInt32,
	}
	expected = `testdata
├── dir1
│   ├── dir11
│   │   └── file3
│   ├── dir12
│   │   └── file4
│   └── file2
└── file1
`

	allFilesOpts = &Options{
		allFiles:     true,
		maxDeepLevel: math.MaxInt32,
	}
	expectedForAllFilesOpts = `testdata
├── .dotdir1
│   └── .dotfile1
├── dir1
│   ├── dir11
│   │   └── file3
│   ├── dir12
│   │   └── file4
│   └── file2
└── file1
`

	maxDeepLevelOpts = &Options{
		allFiles:     false,
		maxDeepLevel: 1,
	}
	expectedForMaxDeepLevelOpts = `testdata
├── dir1
└── file1
`
)

func TestTreeCommand(t *testing.T) {
	cases := []struct {
		opts     *Options
		expected string
	}{
		{opts: opts, expected: expected},
		{opts: allFilesOpts, expected: expectedForAllFilesOpts},
		{opts: maxDeepLevelOpts, expected: expectedForMaxDeepLevelOpts},
	}

	for i, c := range cases {
		var b bytes.Buffer
		c.opts.out = &b
		tree, err := MakeTree("testdata", c.opts)
		if err != nil {
			t.Fatalf("failed to make tree %v", err)
		}
		tree.Print(c.opts)

		if b.String() != c.expected {
			t.Errorf("case %d:\ngot:\n%v\nwant:\n%v", i, b.String(), c.expected)
		}
	}
}
