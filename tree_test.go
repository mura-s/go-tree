package main

import (
	"bytes"
	"testing"
)

var (
	opts = &Options{
		AllFiles:     false,
		MaxDeepLevel: DefaultMaxDeepLevel,
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
		AllFiles:     true,
		MaxDeepLevel: DefaultMaxDeepLevel,
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
		AllFiles:     false,
		MaxDeepLevel: 1,
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
		c.opts.Out = &b
		tree, err := MakeTree("testdata", c.opts)
		if err != nil {
			t.Fatalf("failed to make tree %v", err)
		}
		tree.Print()

		if b.String() != c.expected {
			t.Errorf("case %d:\ngot:\n%v\nwant:\n%v", i, b.String(), c.expected)
		}
	}
}
