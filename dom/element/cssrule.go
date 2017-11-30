package element


		// CSSRule interface
		// js:"CSSRule"
		type CSSRule interface {
		
			
	
			
		// CSSText prop 
		// js:"cssText"
		CSSText() (cssText string)
		

		// SetCSSText prop 
		// js:"cssText"
		SetCSSText (cssText string)
		
		// ParentRule prop 
		// js:"parentRule"
		ParentRule() (parentRule CSSRule)
		
		// ParentStyleSheet prop 
		// js:"parentStyleSheet"
		ParentStyleSheet() (parentStyleSheet *CSSStyleSheet)
		
		// Type prop 
		// js:"type"
		Type() (kind uint8)
		
		}
	