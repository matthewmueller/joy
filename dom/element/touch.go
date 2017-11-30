package element



		

		// Touch struct
		// js:"Touch,omit"
		type Touch struct {
		}

		

		
		// ClientX prop 
		// js:"clientX"
		func (*Touch) ClientX() (clientX int) {
			js.Rewrite("$_.clientX")
			return clientX
		}
		
		// ClientY prop 
		// js:"clientY"
		func (*Touch) ClientY() (clientY int) {
			js.Rewrite("$_.clientY")
			return clientY
		}
		
		// Identifier prop 
		// js:"identifier"
		func (*Touch) Identifier() (identifier int) {
			js.Rewrite("$_.identifier")
			return identifier
		}
		
		// PageX prop 
		// js:"pageX"
		func (*Touch) PageX() (pageX int) {
			js.Rewrite("$_.pageX")
			return pageX
		}
		
		// PageY prop 
		// js:"pageY"
		func (*Touch) PageY() (pageY int) {
			js.Rewrite("$_.pageY")
			return pageY
		}
		
		// ScreenX prop 
		// js:"screenX"
		func (*Touch) ScreenX() (screenX int) {
			js.Rewrite("$_.screenX")
			return screenX
		}
		
		// ScreenY prop 
		// js:"screenY"
		func (*Touch) ScreenY() (screenY int) {
			js.Rewrite("$_.screenY")
			return screenY
		}
		
		// Target prop 
		// js:"target"
		func (*Touch) Target() (target EventTarget) {
			js.Rewrite("$_.target")
			return target
		}
		
	