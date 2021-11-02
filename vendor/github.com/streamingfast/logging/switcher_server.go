package logging

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type logChangeReq struct {
	Inputs string `json:"inputs"`
	Level  string `json:"level"`
}

type switcherServerHandler struct {
	registry *registry
}

func (h *switcherServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	in := logChangeReq{}
	if err := decoder.Decode(&in); err != nil {
		http.Error(w, fmt.Sprintf("cannot unmarshal request: %s", err), 400)
		return
	}

	if in.Inputs == "" {
		http.Error(w, "inputs not defined, should be comma-separated list of words or a regular expressions", 400)
		return
	}

	level := strings.ToUpper(in.Level)
	if level != "WARN" && level != "WARNING" && level != "INFO" && level != "DEBUG" && level != "TRACE" {
		http.Error(w, fmt.Sprintf("invalid level value %q", in.Level), 400)
		return
	}

	spec := newLogLevelSpecFromMap(map[string]string{
		strings.ToUpper(in.Level): in.Inputs,
	})

	globalRegistry.overrideFromSpec(spec, globalRegistry.factory)

	w.Write([]byte("ok"))
}
