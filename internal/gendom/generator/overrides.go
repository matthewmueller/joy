package generator

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

	"void": "void",

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
