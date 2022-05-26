// This file is inspired from
// https://github.com/StevenACoffman/teamboard/main/pkg/tools

//go:build generate

package generate

import _ "github.com/Khan/genqlient"

// Generate Go code to send GraphQL query.
//go:generate go run github.com/Khan/genqlient ../opyn/generate-graphql-config.yaml
//go:generate go run github.com/Khan/genqlient ../thales/generate-graphql-config.yaml

// There is nothing here intentionally (see README.md).
