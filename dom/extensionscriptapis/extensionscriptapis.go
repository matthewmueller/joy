package extensionscriptapis

import "github.com/matthewmueller/joy/macro"

// ExtensionScriptApis struct
// js:"ExtensionScriptApis,omit"
type ExtensionScriptApis struct {
}

// ExtensionIDToShortID fn
// js:"extensionIdToShortId"
func (*ExtensionScriptApis) ExtensionIDToShortID(extensionId string) (i int) {
	macro.Rewrite("$_.extensionIdToShortId($1)", extensionId)
	return i
}

// FireExtensionAPITelemetry fn
// js:"fireExtensionApiTelemetry"
func (*ExtensionScriptApis) FireExtensionAPITelemetry(functionName string, isSucceeded bool, isSupported bool) {
	macro.Rewrite("$_.fireExtensionApiTelemetry($1, $2, $3)", functionName, isSucceeded, isSupported)
}

// GenericFunction fn
// js:"genericFunction"
func (*ExtensionScriptApis) GenericFunction(routerAddress interface{}, parameters *string, callbackId *int) {
	macro.Rewrite("$_.genericFunction($1, $2, $3)", routerAddress, parameters, callbackId)
}

// GenericSynchronousFunction fn
// js:"genericSynchronousFunction"
func (*ExtensionScriptApis) GenericSynchronousFunction(functionId int, parameters *string) (s string) {
	macro.Rewrite("$_.genericSynchronousFunction($1, $2)", functionId, parameters)
	return s
}

// GetExtensionID fn
// js:"getExtensionId"
func (*ExtensionScriptApis) GetExtensionID() (s string) {
	macro.Rewrite("$_.getExtensionId()")
	return s
}

// RegisterGenericFunctionCallbackHandler fn
// js:"registerGenericFunctionCallbackHandler"
func (*ExtensionScriptApis) RegisterGenericFunctionCallbackHandler(callbackHandler func()) {
	macro.Rewrite("$_.registerGenericFunctionCallbackHandler($1)", callbackHandler)
}

// RegisterGenericPersistentCallbackHandler fn
// js:"registerGenericPersistentCallbackHandler"
func (*ExtensionScriptApis) RegisterGenericPersistentCallbackHandler(callbackHandler func()) {
	macro.Rewrite("$_.registerGenericPersistentCallbackHandler($1)", callbackHandler)
}
