# Generator directory

## Intentionally mostly-empty Go file

This directory contains a mostly-empty `generate.go` file to:

1. keep the generator tool as a dependency:

    ```go
    import _ "github.com/Khan/genqlient"
    ```

2. and provide a `go:generate` directive.

No other source code should be found in this directory.

## Targets

* use a go-based generator tool (e.g. `genqlient`)
* everyone use the same version of that tool
* track the tool's version in `go.mod` file

## The `import` statement

The import statement allows the `go` command to
precisely record the tool version in `go.mod`.

```go
import _ "github.com/Khan/genqlient"
```

## The `//go:build generator` build tag

The `//go:build generator` build tag ensures the dependency
on the generator tool is not compiled into the binary.

## A place for the `go:generate` directives

The `go generate ./...` runs commands described by directives
within existing files. Those commands can run any process,
but the intent is to create/update Go source files.

Go generate is never run automatically by
`go build`, `go get`, `go test`, and so on.
It must be run explicitly.

Go generate scans the file for directives,
which are lines of the form:

```go
//go:generate command argument...
```

Note: no space within "`//go:generate`".

The `command` is the generator tool to be run,
corresponding to an executable file that can be run locally.
The `command` can be `go run github.com/Khan/genqlient`.

## See also

Original inspiration:
<https://github.com/StevenACoffman/teamboard/tree/main/pkg/tools>

For an example, see
["Go Modules by Example" walkthrough](https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md).

A historical discussion of the approach is in
[this comment](https://github.com/golang/go/issues/25922#issuecomment-412992431).
