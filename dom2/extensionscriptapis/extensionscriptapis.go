package extensionscriptapis

import "github.com/matthewmueller/golly/js"

// ExtensionScriptApis struct
// js:"ExtensionScriptApis,omit"
type ExtensionScriptApis struct {
}

// ExtensionIDToShortID
func (*ExtensionScriptApis) ExtensionIDToShortID(extensionId string) (i int) {
	js.Rewrite("$<.ExtensionIDToShortID($1)", extensionId)
	return i
}

// FireExtensionAPITelemetry
func (*ExtensionScriptApis) FireExtensionAPITelemetry(functionName string, isSucceeded bool, isSupported bool) {
	js.Rewrite("$<.FireExtensionAPITelemetry($1, $2, $3)", functionName, isSucceeded, isSupported)
}

// GenericFunction
func (*ExtensionScriptApis) GenericFunction(routerAddress interface{}, parameters *string, callbackId *int) {
	js.Rewrite("$<.GenericFunction($1, $2, $3)", routerAddress, parameters, callbackId)
}

// GenericSynchronousFunction
func (*ExtensionScriptApis) GenericSynchronousFunction(functionId int, parameters *string) (s string) {
	js.Rewrite("$<.GenericSynchronousFunction($1, $2)", functionId, parameters)
	return s
}

// GetExtensionID
func (*ExtensionScriptApis) GetExtensionID() (s string) {
	js.Rewrite("$<.GetExtensionID()")
	return s
}

// RegisterGenericFunctionCallbackHandler
func (*ExtensionScriptApis) RegisterGenericFunctionCallbackHandler(callbackHandler func()) {
	js.Rewrite("$<.RegisterGenericFunctionCallbackHandler($1)", callbackHandler)
}

// RegisterGenericPersistentCallbackHandler
func (*ExtensionScriptApis) RegisterGenericPersistentCallbackHandler(callbackHandler func()) {
	js.Rewrite("$<.RegisterGenericPersistentCallbackHandler($1)", callbackHandler)
}
