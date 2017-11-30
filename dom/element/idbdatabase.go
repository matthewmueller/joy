package element



		var _ EventTarget = (*IDBDatabase)(nil)
		

		// IDBDatabase struct
		// js:"IDBDatabase,omit"
		type IDBDatabase struct {
		}

		
			// Close fn 
			// js:"close"
			func (*IDBDatabase) Close() {
				js.Rewrite("$_.close()", )
			}
		
		// CreateObjectStore fn 
		// js:"createObjectStore"
		func (*IDBDatabase) CreateObjectStore(name string, optionalParameters *idbobjectstoreparameters.IDBObjectStoreParameters) (i *IDBObjectStore) {
			js.Rewrite("$_.createObjectStore($1, $2)", name, optionalParameters)
			return i
		}
	
			// DeleteObjectStore fn 
			// js:"deleteObjectStore"
			func (*IDBDatabase) DeleteObjectStore(name string) {
				js.Rewrite("$_.deleteObjectStore($1)", name)
			}
		
		// Transaction fn 
		// js:"transaction"
		func (*IDBDatabase) Transaction(storeNames interface{}, mode *idbtransactionmode.IDBTransactionMode) (i *IDBTransaction) {
			js.Rewrite("$_.transaction($1, $2)", storeNames, mode)
			return i
		}
	
			// AddEventListener fn 
			// js:"addEventListener"
			func (*IDBDatabase) AddEventListener(kind string, listener func (evt Event), useCapture bool) {
				js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
			}
		
		// DispatchEvent fn 
		// js:"dispatchEvent"
		func (*IDBDatabase) DispatchEvent(evt Event) (b bool) {
			js.Rewrite("$_.dispatchEvent($1)", evt)
			return b
		}
	
			// RemoveEventListener fn 
			// js:"removeEventListener"
			func (*IDBDatabase) RemoveEventListener(kind string, listener func (evt Event), useCapture bool) {
				js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
			}
		

		
		// Name prop 
		// js:"name"
		func (*IDBDatabase) Name() (name string) {
			js.Rewrite("$_.name")
			return name
		}
		
		// ObjectStoreNames prop 
		// js:"objectStoreNames"
		func (*IDBDatabase) ObjectStoreNames() (objectStoreNames *domstringlist.DOMStringList) {
			js.Rewrite("$_.objectStoreNames")
			return objectStoreNames
		}
		
		// Onabort prop 
		// js:"onabort"
		func (*IDBDatabase) Onabort() (onabort func(Event)) {
			js.Rewrite("$_.onabort")
			return onabort
		}
		

		// SetOnabort prop 
		// js:"onabort"
		func (*IDBDatabase) SetOnabort (onabort func(Event)) {
			js.Rewrite("$_.onabort = $1", onabort)
		}
		
		// Onerror prop 
		// js:"onerror"
		func (*IDBDatabase) Onerror() (onerror func(Event)) {
			js.Rewrite("$_.onerror")
			return onerror
		}
		

		// SetOnerror prop 
		// js:"onerror"
		func (*IDBDatabase) SetOnerror (onerror func(Event)) {
			js.Rewrite("$_.onerror = $1", onerror)
		}
		
		// Version prop 
		// js:"version"
		func (*IDBDatabase) Version() (version uint64) {
			js.Rewrite("$_.version")
			return version
		}
		
	