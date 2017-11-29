
		package window
		
		// js:"Event,omit"
		type Event interface {
		
			
			// PreventDefault 
			// js:"preventDefault,rewrite=preventdefault"
			PreventDefault()
		
			// StopImmediatePropagation 
			// js:"stopImmediatePropagation,rewrite=stopimmediatepropagation"
			StopImmediatePropagation()
		
			// StopPropagation 
			// js:"stopPropagation,rewrite=stoppropagation"
			StopPropagation()
		
	
			
		// Type prop 
		// js:"type,rewrite=kind"
		Type() (kind string)
		
		}

		
			// preventdefault fn 
			func preventdefault() {
				js.Rewrite("$_.preventDefault()", )
			}
		
			// stopimmediatepropagation fn 
			func stopimmediatepropagation() {
				js.Rewrite("$_.stopImmediatePropagation()", )
			}
		
			// stoppropagation fn 
			func stoppropagation() {
				js.Rewrite("$_.stopPropagation()", )
			}
		
		// kind prop 
		func kind() (kind string) {
			js.Rewrite("$_.type")
			return kind
		}
		
	package window

import "github.com/matthewmueller/golly/js"

// js:"Event,omit"
type Event interface {

	// PreventDefault
	// js:"preventDefault,rewrite=preventdefault"
	PreventDefault()

	// StopImmediatePropagation
	// js:"stopImmediatePropagation,rewrite=stopimmediatepropagation"
	StopImmediatePropagation()

	// StopPropagation
	// js:"stopPropagation,rewrite=stoppropagation"
	StopPropagation()

	// Type prop
	// js:"type,rewrite=kind"
	Type() (kind string)
}

// preventdefault fn
func preventdefault() {
	js.Rewrite("$_.preventDefault()")
}

// stopimmediatepropagation fn
func stopimmediatepropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// stoppropagation fn
func stoppropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// kind prop
func kind() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
