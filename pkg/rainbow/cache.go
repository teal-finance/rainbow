// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

import (
	"fmt"
)

type Cache struct {
	options     []Option
	callPut     CallPut
	initialized bool
}

func NewCache() Cache {
	return Cache{
		initialized: false,
		options:     []Option{},
		callPut:     CallPut{Rows: []Row{}},
	}
}

func (c Cache) Empty() bool {
	return !c.initialized
}

func (c *Cache) Refresh(options []Option) {
	c.initialized = true
	c.options = options
	c.callPut = buildCallPut(options)
}

func (c *Cache) Options() ([]Option, error) {
	if c.Empty() {
		return nil, fmt.Errorf("Cache is empty")
	}

	return c.options, nil
}

func (c *Cache) CallPut() (CallPut, error) {
	if c.Empty() {
		return c.callPut, fmt.Errorf("Cache is empty")
	}

	return c.callPut, nil
}
