// Teal.Finance/Server is an opinionated boilerplate API and website server.
// Copyright (C) 2021 Teal.Finance contributors
//
// This file is part of Teal.Finance/Server, licensed under LGPL-3.0-or-later.
//
// Teal.Finance/Server is free software: you can redistribute it
// and/or modify it under the terms of the GNU Lesser General Public License
// either version 3 of the License, or (at your option) any later version.
//
// Teal.Finance/Server is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty
// of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.

// package opa manages the Open Policy Agent.
// see https://www.openpolicyagent.org/docs/edge/integration/#integrating-with-the-go-api
package opa

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/teal-finance/server/reserr"
)

type Policy struct {
	Compiler *ast.Compiler
	ResErr   reserr.ResErr
}

// New creates a new Policy by loading rego files.
func New(filenames []string, reserr reserr.ResErr) (Policy, error) {
	compiler, err := Load(filenames)

	return Policy{compiler, reserr}, err
}

// Ready returns true if the Policy compiler contains rules.
func (opa Policy) Ready() bool {
	return opa.Compiler != nil
}

// Load check the Rego filenames and loads them to build the OPA compiler.
func Load(filenames []string) (*ast.Compiler, error) {
	trimmedFN := make([]string, 0, len(filenames))

	for _, fn := range filenames {
		f := strings.TrimSpace(fn)
		if f == "" {
			log.Printf("OPA: skip empty filename %q", fn)

			continue
		}

		trimmedFN = append(trimmedFN, f)
	}

	if len(trimmedFN) == 0 {
		return nil, nil
	}

	modules := map[string]string{}

	for _, f := range filenames {
		log.Printf("OPA: load %q", f)

		content, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, fmt.Errorf("OPA: ReadFile %w", err)
		}

		modules[path.Base(f)] = string(content)
	}

	return ast.CompileModules(modules)
}

// Auth is the HTTP middleware to check endpoint authorization.
func (opa Policy) Auth(next http.Handler) http.Handler {
	log.Print("Middleware OPA: ", opa.Compiler.Modules)

	compiler := opa.Compiler
	resErr := opa.ResErr

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input := map[string]interface{}{
			"method": r.Method,
			"path":   strings.Split(strings.Trim(r.URL.Path, "/"), "/"),
			"token":  r.Header.Get("Authorization"),
		}

		// evaluation
		rg := rego.New(
			rego.Query("data.auth.allow"),
			rego.Compiler(compiler),
			rego.Input(input),
		)

		rs, err := rg.Eval(context.Background())
		if err != nil || len(rs) == 0 {
			log.Print("ERROR OPA Eval: ", err)
			resErr.Write(w, r, http.StatusInternalServerError, "Internal Server Error #1")

			return
		}

		allow, ok := rs[0].Expressions[0].Value.(bool)
		if !ok {
			log.Print("ERROR missing OPA data in ", rs)
			resErr.Write(w, r, http.StatusInternalServerError, "Internal Server Error #2")

			return
		}

		if !allow {
			log.Print("OPA unauthorize " + r.RemoteAddr + " " + r.RequestURI)
			resErr.Write(w, r, http.StatusUnauthorized, "Unauthorized")

			return
		}

		next.ServeHTTP(w, r)
	})
}
