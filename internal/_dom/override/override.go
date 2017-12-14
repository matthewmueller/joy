// Package override TODO
package override

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/dom/index"
)

type task struct {
	Type string           `json:"type,omitempty"`
	Data map[string]*data `json:"data,omitempty"`
}

type data struct {
	Properties *[]struct {
		Name     string `json:"name,omitempty"`
		Type     string `json:"type,omitempty"`
		Readonly bool   `json:"readonly,omitempty"`
		Comment  string `json:"comment,omitempty"`
	} `json:"properties,omitempty"`
}

// Override fn
func Override(idx index.Index) (index.Index, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return idx, fmt.Errorf("error getting file")
	}

	dir := path.Dir(file)
	override, err := ioutil.ReadFile(path.Join(dir, "..", "inputs", "override.json"))
	if err != nil {
		return idx, err
	}

	var tasks []task
	if err := json.Unmarshal(override, &tasks); err != nil {
		return idx, err
	}

	for _, task := range tasks {
		switch task.Type {
		case "replace":
			if err := replace(idx, task.Data); err != nil {
				return idx, err
			}
		}
	}

	return idx, nil
}

func replace(idx index.Index, changes map[string]*data) error {
	for id, r := range changes {
		def := idx.Find(id)
		if def == nil {
			continue
		}

		log.Infof("got def=%s", def.ID())
		_ = r
		// def, isset := idx[id]
		// if !isset {
		// 	continue
		// }

		// def.
	}
	return nil
}
