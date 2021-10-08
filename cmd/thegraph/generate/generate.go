// +build generate

package generate

import (
	_ "github.com/Khan/genqlient"
)

//go:generate go run github.com/Khan/genqlient ../generate-config.yaml

// There is nothing here intentionally

// See: https://github.com/StevenACoffman/teamboard/main/pkg/tools/README.md

// There is one file in this directory, and it is mostly empty, but very important for two reasons:
//
// 1. To require a dependency on a tool, using a build constraint comment (a build tag)
//    at the top of the file, to ensure it is not compiled into a binary
//
// 2. To provide a `go:generate` directives.
//
// No other actual source code should be found in this directory.
//
//
// 1. How can I track tool dependencies for a module?
// --------------------------------------------------
//
// If you:
//
// * want to use a go-based tool (e.g. `stringer`) while working on a module, and
// * want to ensure that everyone is using the same version of that tool
//   while tracking the tool's version in your module's `go.mod` file
//
// then one currently recommended approach is to add a `tools.go` file
// to your module that includes import statements for the tools of interest
// (such as `import _ "golang.org/x/tools/cmd/stringer"`),
// along with a `// +build tools` build constraint.
// The import statements allow the `go` command to precisely record
// the version information for your tools in your module's `go.mod`,
// while the `// +build tools` build constraint prevents
// your normal builds from actually importing your tools.
//
// For a concrete example of how to do this, please see this
// ["Go Modules by Example" walkthrough](https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md).
//
// A historical discussion of the approach along of how to do this is in
// [this comment in #25922](https://github.com/golang/go/issues/25922#issuecomment-412992431).
//
// The brief rationale from #25922:
//
// > I think the tools.go file is, in fact, the best practice
// > for tool dependencies, certainly for Go 1.11.
// >
// > I like it because it does not introduce new mechanisms.
// >
// > It simply reuses existing ones.
//
//
// 2. A place for go:generate directives
// -------------------------------------
//
// The `go generate ./...` runs commands described by directives
// within existing files. Those commands can run any process,
// but the intent is to create or update Go source files.
//
// Go generate is never run automatically by
// go build, go get, go test, and so on.
// It must be run explicitly.
//
// Go generate scans the file for directives,
// which are lines of the form:
//
//     //go:generate command argument...
//
// (note: no leading spaces and no space in "//go") where command
// is the generator to be run, corresponding to an executable file
// that can be run locally. It must either be in the shell path
// (gofmt), a fully qualified path (/usr/you/bin/mytool), or a
// command alias.
