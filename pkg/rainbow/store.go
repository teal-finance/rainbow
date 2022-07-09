// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package rainbow

type Store interface {
	InsertOptions(options []Option) error
	Options(args StoreArgs) ([]Option, error)
}

type StoreArgs struct {
	Assets    []string
	Expiries  []string
	Providers []string
}
