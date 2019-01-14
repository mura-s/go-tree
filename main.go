package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

var (
	allFiles     = flag.Bool("a", false, "")
	maxDeepLevel = flag.Int("L", math.MaxInt32, "")

	usage = fmt.Sprintf(`Usage:
  %s [Options...] [--] <directory_path>
----- Options -----
-a          All files which include dotfiles are listed.
-h,-help    Print usage and exit.
-L level    Max display depth of the directory tree.
`, os.Args[0])
)

func main() {
	flag.Usage = func() { fmt.Fprint(os.Stderr, usage) }
	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(2)
	}
	path := flag.Args()[0]
	opts := &Options{
		allFiles:     *allFiles,
		maxDeepLevel: *maxDeepLevel,
		out:          os.Stdout,
	}

	tree, err := MakeTree(path, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to make tree: %v", err)
		os.Exit(1)
	}
	tree.Print(opts)
}
