// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package server

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
)

// loadPolicies loads one or more policies files.
func (s *Server) loadPolicies() (err error) {
	modules := map[string]string{}

	for _, f := range s.OPAFilenames {
		log.Print("OPA: load file ", f)

		content, e := ioutil.ReadFile(f)
		if e != nil {
			return fmt.Errorf("OPA: ReadFile %w", e)
		}

		modules[path.Base(f)] = string(content)
	}

	s.compiler, err = ast.CompileModules(modules)
	if err != nil {
		return fmt.Errorf("OPA: CompileModules %w", err)
	}

	return nil
}

// auth is the HTTP middleware for authorization.
func (s *Server) auth(next http.Handler) http.Handler {
	log.Print("Middleware OPA: ", s.compiler.Modules)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input := map[string]interface{}{
			"method": r.Method,
			"path":   strings.Split(strings.Trim(r.URL.Path, "/"), "/"),
			"token":  r.Header.Get("Authorization"),
		}

		// evaluation
		rg := rego.New(
			rego.Query("data.auth.allow"),
			rego.Compiler(s.compiler),
			rego.Input(input),
		)

		rs, err := rg.Eval(context.Background())
		if err != nil || len(rs) == 0 {
			log.Print("ERROR OPA Eval: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			s.Resp.Error(w, r, "Internal Server Error #1")
			return
		}

		allow, ok := rs[0].Expressions[0].Value.(bool)
		if !ok {
			log.Print("ERROR missing OPA data in ", rs)

			w.WriteHeader(http.StatusInternalServerError)
			s.Resp.Error(w, r, "Internal Server Error #2")

			return
		}

		if !allow {
			log.Print("OPA unauthorize " + r.RemoteAddr + " " + r.RequestURI)

			w.WriteHeader(http.StatusUnauthorized)
			s.Resp.Error(w, r, "Unauthorized")

			return
		}

		next.ServeHTTP(w, r)
	})
}
