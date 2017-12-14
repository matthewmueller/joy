package parser

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"sort"
	"strings"

	"github.com/matthewmueller/joy/internal/dom/def"
	"github.com/matthewmueller/joy/internal/dom/defs"
	"github.com/matthewmueller/joy/internal/dom/index"
	"github.com/matthewmueller/joy/internal/dom/raw"
	"github.com/pkg/errors"
)

// docs
type docs struct {
	Interfaces []*struct {
		Name    string `json:"name"`
		Members struct {
			Properties []struct {
				Name    string `json:"name"`
				Comment string `json:"comment"`
			} `json:"property"`
			Methods []struct {
				Name    string `json:"name"`
				Comment string `json:"comment"`
			} `json:"method"`
		} `json:"members"`
	} `json:"interfaces"`
}

// Parse the data
func Parse(api string, doc string) (index.Index, error) {
	var a raw.API
	if e := xml.Unmarshal([]byte(api), &a); e != nil {
		return nil, errors.Wrap(e, "error parsing xml")
	}

	var d docs
	if e := json.Unmarshal([]byte(doc), &d); e != nil {
		return nil, errors.Wrap(e, "error parsing json")
	}

	methodmap := map[string]int{}
	propmap := map[string]int{}
	docmap := map[string]int{}
	for i, iface := range d.Interfaces {
		if _, isset := docmap[iface.Name]; isset {
			return nil, fmt.Errorf("duplicate doc %s", iface.Name)
		}
		docmap[iface.Name] = i

		for i, method := range iface.Members.Methods {
			methodmap[iface.Name+" "+method.Name] = i
		}
		for i, prop := range iface.Members.Properties {
			propmap[iface.Name+" "+prop.Name] = i
		}
	}

	// add in the documentation if it's present
	for _, iface := range a.Interfaces {
		name := iface.Name

		docidx, isset := docmap[name]
		if !isset {
			continue
		}
		doc := d.Interfaces[docidx]
		if doc == nil {
			continue
		}

		for _, m := range iface.Methods {
			if i, isset := methodmap[name+" "+m.Name]; isset {
				comment := strings.Trim(doc.Members.Methods[i].Comment, "/* \r\n\t/")
				m.Comment = strings.Join(strings.Split(comment, "\n"), "\n//")
			}
		}

		for _, m := range iface.Properties {
			if i, isset := propmap[name+" "+m.Name]; isset {
				comment := strings.Trim(doc.Members.Properties[i].Comment, "/* \r\n\t/")
				m.Comment = strings.Join(strings.Split(comment, "\n"), "\n//")
			}
		}
	}

	idx := index.Index{}

	// add to the index
	for _, d := range a.CallbackFunctions {
		def := defs.NewCallback(idx, d)
		if idx[def.ID()] != nil {
			return nil, fmt.Errorf("def already exists: %s", def.ID())
		}
		idx[def.ID()] = def
	}

	for _, d := range a.CallbackInterfaces {
		def := defs.NewCallbackInterface(idx, d)
		if idx[def.ID()] != nil {
			return nil, fmt.Errorf("def already exists: %s", def.ID())
		}
		idx[def.ID()] = def
	}

	for _, d := range a.Dictionaries {
		def := defs.NewDictionary(idx, d)
		if idx[def.ID()] != nil {
			return nil, fmt.Errorf("def already exists: %s", def.ID())
		}
		idx[def.ID()] = def
	}

	for _, d := range a.Enums {
		def := defs.NewEnum(idx, d)
		if idx[def.ID()] != nil {
			return nil, fmt.Errorf("def already exists: %s", def.ID())
		}
		idx[def.ID()] = def
	}

	for _, d := range a.TypeDefs {
		def := defs.NewTypeDef(idx, d)
		if idx[def.ID()] != nil {
			return nil, fmt.Errorf("def already exists: %s", def.ID())
		}
	}

	for _, d := range a.Interfaces {
		def := defs.NewInterface(idx, d)
		if idx[def.ID()] != nil {
			return nil, fmt.Errorf("def already exists: %s", def.ID())
		}
		idx[def.ID()] = def
	}

	for _, d := range a.MixinInterfaces {
		def := defs.NewInterface(idx, d)
		if idx[def.ID()] != nil {
			return nil, fmt.Errorf("def already exists: %s", def.ID())
		}
		idx[def.ID()] = def
	}

	return idx, nil
}

func sortDefs(m map[string]def.Definition) (sorted []def.Definition) {
	var order []string
	for k := range m {
		order = append(order, k)
	}

	sort.Strings(order)

	for _, o := range order {
		sorted = append(sorted, m[o])
	}

	return sorted
}
