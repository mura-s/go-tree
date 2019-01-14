# go-tree

An implementation of the tree command written in Go.

```
$ cd testdata
$ ../go-tree .
.
├── dir1
│   ├── dir11
│   │   └── file3
│   ├── dir12
│   │   └── file4
│   └── file2
└── file1
```

## Test

```
$ go test
```

## Build

```
$ go build -o go-tree
```

## Usage

```
$ ./go-tree -h
Usage:
  ./go-tree [Options...] [--] <directory_path>
----- Options -----
-a          All files which include dotfiles are listed.
-h,-help    Print usage and exit.
-L level    Max display depth of the directory tree.
```
