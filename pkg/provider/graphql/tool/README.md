# Generator directory

See: <https://github.com/StevenACoffman/teamboard/tree/main/pkg/tools>

## Purpose

This directory is mostly empty, but required:

1. To keep the generator tool as a dependency:

    ```go
    import _ "github.com/Khan/genqlient"
    ```

2. To provide a `go:generate` directive.

No other actual source code should be found in this directory.

## Track the generator tool version

If you:

* want to use a go-based tool (e.g. `genqlient`)
* want to everyone use the same version of that tool
* want to track the tool's version in `go.mod` file

then the recommended approach is to add a `tool/generate.go` file
containing a `//go:build tool` build tag and an import statement such as:

```go
import _ "github.com/Khan/genqlient"
```

The import statement allows the `go` command to
precisely record the tool version in `go.mod`
while the build tag `//go:build tool` ensures
it is not compiled into your binary.

For an example, see
["Go Modules by Example" walkthrough](https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md).

A historical discussion of the approach is in
[this comment](https://github.com/golang/go/issues/25922#issuecomment-412992431).

## A place for `go:generate` directives

The `go generate ./...` runs commands described by directives
within existing files. Those commands can run any process,
but the intent is to create or update Go source files.

Go generate is never run automatically by
`go build`, `go get`, `go test`, and so on.
It must be run explicitly.

Go generate scans the file for directives,
which are lines of the form:

```go
//go:generate command argument...
```

Note: no leading spaces and no space in "`//go:generate`"

The `command` is the generator tool to be run,
corresponding to an executable file that can be run locally.
