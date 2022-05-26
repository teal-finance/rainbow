# Directory about External Tools

## Intentionally mostly-empty Go files

This directory contains two mostly-empty Go file to:

1. keep the external tools as dependencies:

    ```go
    import _ "github.com/Khan/genqlient"
    ```

2. and provide a `go:generate` directive.

No other source code should be found in this directory.

## Targets

* use go-based tools (e.g. `genqlient`)
* everyone use the same version of that tools
* track the tools versions in `go.mod` file

## The `go:generate` directives

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
corresponding to an executable file that can be run locally,
such as `go run github.com/Khan/genqlient`.

## The `import` statement

The import statement allows the `go` command to
precisely record the tool version in `go.mod`.

```go
import _ "github.com/Khan/genqlient"
```

## The `//go:build tools` build tag

The `//go:build tools` build tag ensures the dependencies
on the external tools are not compiled into the binary.

But the `//go:build tools` build tag cannot be
in the Go file containing the `//go:generate`.

## See also

Original inspiration:
<https://github.com/StevenACoffman/teamboard/tree/main/pkg/tools>

For an example, see
["Go Modules by Example" walkthrough](https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md).

A historical discussion of the approach is in
[this comment](https://github.com/golang/go/issues/25922#issuecomment-412992431).
