//go:build tools

// Package tools is inspired from
// https://github.com/StevenACoffman/teamboard/tree/main/pkg/tools
package tools

import (
	_ "github.com/Khan/genqlient"
	// disable golangci-lint because fails when indirect deps are upgraded
	// _ "github.com/golangci/golangci-lint/v2/cmd/golangci-lint"
	// _ "github.com/segmentio/golines"
	// _ "mvdan.cc/gofumpt"
)

// There is nothing here intentionally (see README.md).
