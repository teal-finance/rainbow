// Copyright (C) 2020-2021 TealTicks contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
	log.Printf("Load OPA %v files", len(s.OPAFiles))

	modules := map[string]string{}

	for _, f := range s.OPAFiles {
		content, e := ioutil.ReadFile(f)
		if e != nil {
			return fmt.Errorf("OPA read file: %w", e)
		}

		modules[path.Base(f)] = string(content)
	}

	s.opaCompiler, err = ast.CompileModules(modules)
	if err != nil {
		return fmt.Errorf("OPA compile modules: %w", err)
	}

	return nil
}

// auth is the HTTP middleware for authorization.
func (s *Server) auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input := map[string]interface{}{
			"method": r.Method,
			"path":   strings.Split(strings.Trim(r.URL.Path, "/"), "/"),
			"token":  r.Header.Get("Authorization"),
		}

		// evaluation
		rg := rego.New(
			rego.Query("data.auth.allow"),
			rego.Compiler(s.opaCompiler),
			rego.Input(input),
		)

		rs, err := rg.Eval(context.Background())
		if err != nil || len(rs) == 0 {
			log.Print("ERROR OPA Eval: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			s.ReqError(w, r, "Internal Server Error #1")
			return
		}

		allow, ok := rs[0].Expressions[0].Value.(bool)
		if !ok {
			log.Print("ERROR missing OPA data in ", rs)
			w.WriteHeader(http.StatusInternalServerError)
			s.ReqError(w, r, "Internal Server Error #2")
			return
		}

		if !allow {
			log.Print("OPA unauthorize " + r.RemoteAddr + " " + r.RequestURI)
			w.WriteHeader(http.StatusUnauthorized)
			s.ReqError(w, r, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
