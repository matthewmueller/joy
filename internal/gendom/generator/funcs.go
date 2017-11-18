package generator

import (
	"regexp"
	"strings"

	"github.com/matthewmueller/golly/internal/gen"
)

var resequence = regexp.MustCompile(`sequence<([\w<>]+)>`)
var repromise = regexp.MustCompile(`Promise<([\w<>]+)>`)

// coerce transforms the XML type to a Go type
func coerce(idx *index, kind string) (string, error) {
	kind = strings.TrimSpace(kind)
	isSlice := false

	// TODO: handle unions better
	if strings.Contains(kind, " or ") {
		return "interface{}", nil
	}

	matches := repromise.FindStringSubmatch(kind)
	if len(matches) > 1 {
		kind = matches[1]
	}

	matches = resequence.FindStringSubmatch(kind)
	if len(matches) > 1 {
		isSlice = true
		kind = matches[1]
	}

	if overrides[kind] == "" {
		kind = gen.Pointer(kind)
	} else {
		kind = overrides[kind]
	}

	if isSlice {
		return "[]" + kind, nil
	}
	return kind, nil
}

// check if the function is async
func isAsync(kind string) bool {
	return strings.Contains(kind, "Promise<")
}

// find the event, traversing up if necessary
func findEvent(idx *index, i *Interface, name string) *Event {
	for _, event := range i.Events {
		if event.Name == name {
			return event
		}
	}
	if i.Extends != "" && i.Extends != "Object" && idx.Interfaces[i.Extends] != nil {
		return findEvent(idx, idx.Interfaces[i.Extends], name)
	}
	return nil
}
