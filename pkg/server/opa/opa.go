// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package opa

// This file manages the Open Policy Agent
// https://www.openpolicyagent.org/docs/edge/integration/#integrating-with-the-go-api

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

	"github.com/teal-finance/rainbow/pkg/server/resperr"
)

type Policy struct {
	Compiler *ast.Compiler
	Resp     resperr.RespErr
}

// LoadPolicies loads one or more Rego files.
func LoadPolicies(filenames []string) (*ast.Compiler, error) {
	modules := map[string]string{}

	for _, f := range filenames {
		log.Print("OPA: load ", f)

		content, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, fmt.Errorf("OPA: ReadFile %w", err)
		}

		modules[path.Base(f)] = string(content)
	}

	return ast.CompileModules(modules)
}

// Auth is the HTTP middleware for authorization.
func (opa Policy) Auth(next http.Handler) http.Handler {
	log.Print("Middleware OPA: ", opa.Compiler.Modules)

	compiler := opa.Compiler
	resp := opa.Resp

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
			w.WriteHeader(http.StatusInternalServerError)
			resp.Error(w, r, "Internal Server Error #1")
			return
		}

		allow, ok := rs[0].Expressions[0].Value.(bool)
		if !ok {
			log.Print("ERROR missing OPA data in ", rs)

			w.WriteHeader(http.StatusInternalServerError)
			resp.Error(w, r, "Internal Server Error #2")

			return
		}

		if !allow {
			log.Print("OPA unauthorize " + r.RemoteAddr + " " + r.RequestURI)

			w.WriteHeader(http.StatusUnauthorized)
			resp.Error(w, r, "Unauthorized")

			return
		}

		next.ServeHTTP(w, r)
	})
}
