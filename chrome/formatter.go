package chrome

import (
	"encoding/json"
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
		return string(value[1 : len(value)-1])
	default:
		return string(value)
	}
}

func formatProperty(prop runtime.PropertyPreview) string {
	return prop.String()
}
