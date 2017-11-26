package chrome

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/mafredri/cdp/protocol/runtime"
)

func formatValue(obj runtime.RemoteObject) string {
	preview := obj.Preview
	value := obj.Value

	switch obj.Type {
	case "object":
		if obj.Subtype == nil {
			return preview.String()
		}
		switch *obj.Subtype {
		case "array":
			if len(preview.Properties) == 0 {
				return "[]"
			}

			var arr []string
			for _, prop := range preview.Properties {
				arr = append(arr, formatProperty(prop))
			}
			return "[ " + strings.Join(arr, ", ") + " ]"
		case "error":
			var arr []string
			for _, prop := range preview.Properties {
				arr = append(arr, formatProperty(prop))
			}
			return strings.Join(arr, "\n")
		default:
			bytes, _ := json.Marshal(preview)
			return string(bytes)
		}
	case "string":
		v, e := strconv.Unquote(string(value))
		// ignore error if there is one
		if e != nil {
			return string(value)
		}
		return v
	default:
		return string(value)
	}
}

func formatProperty(prop runtime.PropertyPreview) string {
	if prop.Value == nil {
		return prop.String()
	}

	value := *prop.Value
	switch prop.Type {
	case "string":
		return `'` + value + `'`
	default:
		return prop.String()
	}
}
