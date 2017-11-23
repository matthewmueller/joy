package extensionscriptapis

import "github.com/matthewmueller/golly/js"

// ExtensionScriptApis struct
// js:"ExtensionScriptApis,omit"
type ExtensionScriptApis struct {
}

// ExtensionIDToShortID fn
func (*ExtensionScriptApis) ExtensionIDToShortID(extensionId string) (i int) {
	js.Rewrite("$<.extensionIdToShortId($1)", extensionId)
	return i
}

// FireExtensionAPITelemetry fn
func (*ExtensionScriptApis) FireExtensionAPITelemetry(functionName string, isSucceeded bool, isSupported bool) {
	js.Rewrite("$<.fireExtensionApiTelemetry($1, $2, $3)", functionName, isSucceeded, isSupported)
}

// GenericFunction fn
func (*ExtensionScriptApis) GenericFunction(routerAddress interface{}, parameters *string, callbackId *int) {
	js.Rewrite("$<.genericFunction($1, $2, $3)", routerAddress, parameters, callbackId)
}

// GenericSynchronousFunction fn
func (*ExtensionScriptApis) GenericSynchronousFunction(functionId int, parameters *string) (s string) {
	js.Rewrite("$<.genericSynchronousFunction($1, $2)", functionId, parameters)
	return s
}

// GetExtensionID fn
func (*ExtensionScriptApis) GetExtensionID() (s string) {
	js.Rewrite("$<.getExtensionId()")
	return s
}

// RegisterGenericFunctionCallbackHandler fn
func (*ExtensionScriptApis) RegisterGenericFunctionCallbackHandler(callbackHandler func()) {
	js.Rewrite("$<.registerGenericFunctionCallbackHandler($1)", callbackHandler)
}

// RegisterGenericPersistentCallbackHandler fn
func (*ExtensionScriptApis) RegisterGenericPersistentCallbackHandler(callbackHandler func()) {
	js.Rewrite("$<.registerGenericPersistentCallbackHandler($1)", callbackHandler)
}
