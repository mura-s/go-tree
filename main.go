package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	allFiles     = flag.Bool("a", false, "")
	maxDeepLevel = flag.Int("L", DefaultMaxDeepLevel, "")

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
	path := flag.Arg(0)
	opts := &Options{
		AllFiles:     *allFiles,
		MaxDeepLevel: *maxDeepLevel,
		Out:          os.Stdout,
	}

	tree, err := MakeTree(path, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to make tree: %v", err)
		os.Exit(1)
	}
	tree.Print()
}
