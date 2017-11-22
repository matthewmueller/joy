package index

import (
	"regexp"
	"strings"

	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/gen"
)

// Index struct
type Index map[string]def.Definition

var overrides = map[string]string{
	"DOMString":        "string",
	"USVString":        "string",
	"IDBKeyPath":       "string",
	"AAGUID":           "string",
	"EndOfStreamError": "string",
	"ReadyState":       "string",

	"boolean":                "bool",
	"bool":                   "bool",
	"Boolean":                "bool",
	"MSAttestationStatement": "bool",

	"unsigned long":      "uint",
	"unsigned long long": "uint64",
	"unsigned short":     "uint8",
	"Uint8Array":         "[]uint8",
	"Uint8ClampedArray":  "[]uint8",

	"short":               "int8",
	"long":                "int",
	"long long":           "int64",
	"DOMTimeStamp":        "int",
	"DOMHighResTimeStamp": "int",
	"Int32Array":          "[]int32",

	"float":               "float32",
	"double":              "float32",
	"unrestricted double": "float32",
	"UnrestrictedDouble":  "float32",
	"Float32Array":        "[]float32",

	"void": "",

	"object":       "interface{}",
	"any":          "interface{}",
	"Dictionary":   "interface{}",
	"payloadType":  "interface{}",
	"Entry":        "interface{}",
	"Transferable": "interface{}",
	"BodyInit":     "interface{}",

	"ArrayBuffer":     "[]byte",
	"ArrayBufferView": "[]byte",
	"BufferSource":    "[]byte",
	"octet":           "byte",

	"EventHandler": "EventHandler",

	"RequestInfo": "*Request",

	"function": "func()",
	"Function": "func()",

	"Date": "time.Time",

	"WebKitEntry[]": "[]*WebKitEntry",
}

var resequence = regexp.MustCompile(`sequence<([\w<>]+)>`)
var repromise = regexp.MustCompile(`Promise<([\w<>]+)>`)

// Coerce the type
func (i Index) Coerce(kind string) (string, error) {
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

	if _, isset := overrides[kind]; isset {
		kind = overrides[kind]
	} else {
		kind = gen.Pointer(kind)
	}

	if isSlice {
		return "[]" + kind, nil
	}
	return kind, nil
}

// Find a definition by it's type
func (i Index) Find(kind string) def.Definition {
	kind = strings.TrimSpace(kind)

	// TODO: handle unions better
	if strings.Contains(kind, " or ") {
		return nil
	}

	matches := repromise.FindStringSubmatch(kind)
	if len(matches) > 1 {
		kind = matches[1]
	}

	matches = resequence.FindStringSubmatch(kind)
	if len(matches) > 1 {
		kind = matches[1]
	}

	if overrides[kind] != "" {
		return nil
	}

	return i[kind]
}
