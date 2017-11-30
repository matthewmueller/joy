package element



		var _ EventTarget = (*MSInputMethodContext)(nil)
		

		// MSInputMethodContext struct
		// js:"MSInputMethodContext,omit"
		type MSInputMethodContext struct {
		}

		
		// GetCandidateWindowClientRect fn 
		// js:"getCandidateWindowClientRect"
		func (*MSInputMethodContext) GetCandidateWindowClientRect() (c *clientrect.ClientRect) {
			js.Rewrite("$_.getCandidateWindowClientRect()", )
			return c
		}
	
		// GetCompositionAlternatives fn 
		// js:"getCompositionAlternatives"
		func (*MSInputMethodContext) GetCompositionAlternatives() (s []string) {
			js.Rewrite("$_.getCompositionAlternatives()", )
			return s
		}
	
		// HasComposition fn 
		// js:"hasComposition"
		func (*MSInputMethodContext) HasComposition() (b bool) {
			js.Rewrite("$_.hasComposition()", )
			return b
		}
	
		// IsCandidateWindowVisible fn 
		// js:"isCandidateWindowVisible"
		func (*MSInputMethodContext) IsCandidateWindowVisible() (b bool) {
			js.Rewrite("$_.isCandidateWindowVisible()", )
			return b
		}
	
			// AddEventListener fn 
			// js:"addEventListener"
			func (*MSInputMethodContext) AddEventListener(kind string, listener func (evt Event), useCapture bool) {
				js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
			}
		
		// DispatchEvent fn 
		// js:"dispatchEvent"
		func (*MSInputMethodContext) DispatchEvent(evt Event) (b bool) {
			js.Rewrite("$_.dispatchEvent($1)", evt)
			return b
		}
	
			// RemoveEventListener fn 
			// js:"removeEventListener"
			func (*MSInputMethodContext) RemoveEventListener(kind string, listener func (evt Event), useCapture bool) {
				js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
			}
		

		
		// CompositionEndOffset prop 
		// js:"compositionEndOffset"
		func (*MSInputMethodContext) CompositionEndOffset() (compositionEndOffset uint) {
			js.Rewrite("$_.compositionEndOffset")
			return compositionEndOffset
		}
		
		// CompositionStartOffset prop 
		// js:"compositionStartOffset"
		func (*MSInputMethodContext) CompositionStartOffset() (compositionStartOffset uint) {
			js.Rewrite("$_.compositionStartOffset")
			return compositionStartOffset
		}
		
		// Oncandidatewindowhide prop 
		// js:"oncandidatewindowhide"
		func (*MSInputMethodContext) Oncandidatewindowhide() (oncandidatewindowhide func(Event)) {
			js.Rewrite("$_.oncandidatewindowhide")
			return oncandidatewindowhide
		}
		

		// SetOncandidatewindowhide prop 
		// js:"oncandidatewindowhide"
		func (*MSInputMethodContext) SetOncandidatewindowhide (oncandidatewindowhide func(Event)) {
			js.Rewrite("$_.oncandidatewindowhide = $1", oncandidatewindowhide)
		}
		
		// Oncandidatewindowshow prop 
		// js:"oncandidatewindowshow"
		func (*MSInputMethodContext) Oncandidatewindowshow() (oncandidatewindowshow func(Event)) {
			js.Rewrite("$_.oncandidatewindowshow")
			return oncandidatewindowshow
		}
		

		// SetOncandidatewindowshow prop 
		// js:"oncandidatewindowshow"
		func (*MSInputMethodContext) SetOncandidatewindowshow (oncandidatewindowshow func(Event)) {
			js.Rewrite("$_.oncandidatewindowshow = $1", oncandidatewindowshow)
		}
		
		// Oncandidatewindowupdate prop 
		// js:"oncandidatewindowupdate"
		func (*MSInputMethodContext) Oncandidatewindowupdate() (oncandidatewindowupdate func(Event)) {
			js.Rewrite("$_.oncandidatewindowupdate")
			return oncandidatewindowupdate
		}
		

		// SetOncandidatewindowupdate prop 
		// js:"oncandidatewindowupdate"
		func (*MSInputMethodContext) SetOncandidatewindowupdate (oncandidatewindowupdate func(Event)) {
			js.Rewrite("$_.oncandidatewindowupdate = $1", oncandidatewindowupdate)
		}
		
		// Target prop 
		// js:"target"
		func (*MSInputMethodContext) Target() (target HTMLElement) {
			js.Rewrite("$_.target")
			return target
		}
		
	