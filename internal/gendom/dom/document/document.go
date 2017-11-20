package document

import (
	"github.com/apex/log/handlers/text"
	"github.com/matthewmueller/golly/internal/gendom/dom/attr"
	"github.com/matthewmueller/golly/internal/gendom/dom/cdatasection"
	"github.com/matthewmueller/golly/internal/gendom/dom/comment"
	"github.com/matthewmueller/golly/internal/gendom/dom/documentevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/documentfragment"
	"github.com/matthewmueller/golly/internal/gendom/dom/documenttype"
	"github.com/matthewmueller/golly/internal/gendom/dom/domimplementation"
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/globaleventhandlers"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlallcollection"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlcollection"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlheadelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/location"
	"github.com/matthewmueller/golly/internal/gendom/dom/msmanipulationevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/mssitemodeevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/internal/gendom/dom/nodefilter"
	"github.com/matthewmueller/golly/internal/gendom/dom/nodeiterator"
	"github.com/matthewmueller/golly/internal/gendom/dom/nodelist"
	"github.com/matthewmueller/golly/internal/gendom/dom/nodeselector"
	"github.com/matthewmueller/golly/internal/gendom/dom/processinginstruction"
	"github.com/matthewmueller/golly/internal/gendom/dom/rng"
	"github.com/matthewmueller/golly/internal/gendom/dom/selection"
	"github.com/matthewmueller/golly/internal/gendom/dom/stylesheetlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgsvgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/touch"
	"github.com/matthewmueller/golly/internal/gendom/dom/touchlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/treewalker"
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/internal/gendom/dom/visibilitystate"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/internal/gendom/dom/xpathexpression"
	"github.com/matthewmueller/golly/internal/gendom/dom/xpathnsresolver"
	"github.com/matthewmueller/golly/internal/gendom/dom/xpathresult"
	"github.com/matthewmueller/golly/js"
)

type Document struct {
	*node.Node
	*globaleventhandlers.GlobalEventHandlers
	*nodeselector.NodeSelector
	*documentevent.DocumentEvent
}

func (*Document) AdoptNode(source *node.Node) (n *node.Node) {
	js.Rewrite("$<.adoptNode($1)", source)
	return n
}

func (*Document) CaptureEvents() {
	js.Rewrite("$<.captureEvents()")
}

func (*Document) CaretRangeFromPoint(x float32, y float32) (r *rng.Range) {
	js.Rewrite("$<.caretRangeFromPoint($1, $2)", x, y)
	return r
}

func (*Document) Clear() {
	js.Rewrite("$<.clear()")
}

func (*Document) Close() {
	js.Rewrite("$<.close()")
}

func (*Document) CreateAttribute(name string) (a *attr.Attr) {
	js.Rewrite("$<.createAttribute($1)", name)
	return a
}

func (*Document) CreateAttributeNS(namespaceURI string, qualifiedName string) (a *attr.Attr) {
	js.Rewrite("$<.createAttributeNS($1, $2)", namespaceURI, qualifiedName)
	return a
}

func (*Document) CreateCDATASection(data string) (c *cdatasection.CDATASection) {
	js.Rewrite("$<.createCDATASection($1)", data)
	return c
}

func (*Document) CreateComment(data string) (c *comment.Comment) {
	js.Rewrite("$<.createComment($1)", data)
	return c
}

func (*Document) CreateDocumentFragment() (d *documentfragment.DocumentFragment) {
	js.Rewrite("$<.createDocumentFragment()")
	return d
}

func (*Document) CreateElement(tagName string) (e *element.Element) {
	js.Rewrite("$<.createElement($1)", tagName)
	return e
}

func (*Document) CreateElementNS(namespaceURI string, qualifiedName string) (e *element.Element) {
	js.Rewrite("$<.createElementNS($1, $2)", namespaceURI, qualifiedName)
	return e
}

func (*Document) CreateExpression(expression string, resolver *xpathnsresolver.XPathNSResolver) (x *xpathexpression.XPathExpression) {
	js.Rewrite("$<.createExpression($1, $2)", expression, resolver)
	return x
}

func (*Document) CreateNodeIterator(root *node.Node, whatToShow uint, filter *nodefilter.NodeFilter, entityReferenceExpansion bool) (n *nodeiterator.NodeIterator) {
	js.Rewrite("$<.createNodeIterator($1, $2, $3, $4)", root, whatToShow, filter, entityReferenceExpansion)
	return n
}

func (*Document) CreateNSResolver(nodeResolver *node.Node) (x *xpathnsresolver.XPathNSResolver) {
	js.Rewrite("$<.createNSResolver($1)", nodeResolver)
	return x
}

func (*Document) CreateProcessingInstruction(target string, data string) (p *processinginstruction.ProcessingInstruction) {
	js.Rewrite("$<.createProcessingInstruction($1, $2)", target, data)
	return p
}

func (*Document) CreateRange() (r *rng.Range) {
	js.Rewrite("$<.createRange()")
	return r
}

func (*Document) CreateTextNode(data string) (t *text.Text) {
	js.Rewrite("$<.createTextNode($1)", data)
	return t
}

func (*Document) CreateTouch(view *window.Window, target *eventtarget.EventTarget, identifier int, pageX int, pageY int, screenX int, screenY int) (t *touch.Touch) {
	js.Rewrite("$<.createTouch($1, $2, $3, $4, $5, $6, $7)", view, target, identifier, pageX, pageY, screenX, screenY)
	return t
}

func (*Document) CreateTouchList(touches *touch.Touch) (t *touchlist.TouchList) {
	js.Rewrite("$<.createTouchList($1)", touches)
	return t
}

func (*Document) CreateTreeWalker(root *node.Node, whatToShow uint, filter *nodefilter.NodeFilter, entityReferenceExpansion bool) (t *treewalker.TreeWalker) {
	js.Rewrite("$<.createTreeWalker($1, $2, $3, $4)", root, whatToShow, filter, entityReferenceExpansion)
	return t
}

func (*Document) ElementFromPoint(x int, y int) (e *element.Element) {
	js.Rewrite("$<.elementFromPoint($1, $2)", x, y)
	return e
}

func (*Document) Evaluate(expression string, contextNode *node.Node, resolver *xpathnsresolver.XPathNSResolver, kind uint8, result *xpathresult.XPathResult) (x *xpathresult.XPathResult) {
	js.Rewrite("$<.evaluate($1, $2, $3, $4, $5)", expression, contextNode, resolver, kind, result)
	return x
}

func (*Document) ExecCommand(commandId string, showUI bool, value interface{}) (b bool) {
	js.Rewrite("$<.execCommand($1, $2, $3)", commandId, showUI, value)
	return b
}

func (*Document) ExecCommandShowHelp(commandId string) (b bool) {
	js.Rewrite("$<.execCommandShowHelp($1)", commandId)
	return b
}

func (*Document) ExitFullscreen() {
	js.Rewrite("$<.exitFullscreen()")
}

func (*Document) ExitPointerLock() {
	js.Rewrite("$<.exitPointerLock()")
}

func (*Document) Focus() {
	js.Rewrite("$<.focus()")
}

func (*Document) GetElementByID(elementId string) (e *element.Element) {
	js.Rewrite("$<.getElementById($1)", elementId)
	return e
}

func (*Document) GetElementsByClassName(classNames string) (n *nodelist.NodeList) {
	js.Rewrite("$<.getElementsByClassName($1)", classNames)
	return n
}

func (*Document) GetElementsByName(elementName string) (n *nodelist.NodeList) {
	js.Rewrite("$<.getElementsByName($1)", elementName)
	return n
}

func (*Document) GetElementsByTagName(tagname string) (n *nodelist.NodeList) {
	js.Rewrite("$<.getElementsByTagName($1)", tagname)
	return n
}

func (*Document) GetElementsByTagNameNS(namespaceURI string, localName string) (n *nodelist.NodeList) {
	js.Rewrite("$<.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return n
}

func (*Document) GetSelection() (s *selection.Selection) {
	js.Rewrite("$<.getSelection()")
	return s
}

func (*Document) HasFocus() (b bool) {
	js.Rewrite("$<.hasFocus()")
	return b
}

func (*Document) ImportNode(importedNode *node.Node, deep bool) (n *node.Node) {
	js.Rewrite("$<.importNode($1, $2)", importedNode, deep)
	return n
}

func (*Document) MsElementsFromPoint(x float32, y float32) (n *nodelist.NodeList) {
	js.Rewrite("$<.msElementsFromPoint($1, $2)", x, y)
	return n
}

func (*Document) MsElementsFromRect(left float32, top float32, width float32, height float32) (n *nodelist.NodeList) {
	js.Rewrite("$<.msElementsFromRect($1, $2, $3, $4)", left, top, width, height)
	return n
}

func (*Document) Open(url string, name string, features string, replace bool) (i interface{}) {
	js.Rewrite("$<.open($1, $2, $3, $4)", url, name, features, replace)
	return i
}

func (*Document) QueryCommandEnabled(commandId string) (b bool) {
	js.Rewrite("$<.queryCommandEnabled($1)", commandId)
	return b
}

func (*Document) QueryCommandIndeterm(commandId string) (b bool) {
	js.Rewrite("$<.queryCommandIndeterm($1)", commandId)
	return b
}

func (*Document) QueryCommandState(commandId string) (b bool) {
	js.Rewrite("$<.queryCommandState($1)", commandId)
	return b
}

func (*Document) QueryCommandSupported(commandId string) (b bool) {
	js.Rewrite("$<.queryCommandSupported($1)", commandId)
	return b
}

func (*Document) QueryCommandText(commandId string) (s string) {
	js.Rewrite("$<.queryCommandText($1)", commandId)
	return s
}

func (*Document) QueryCommandValue(commandId string) (s string) {
	js.Rewrite("$<.queryCommandValue($1)", commandId)
	return s
}

func (*Document) ReleaseEvents() {
	js.Rewrite("$<.releaseEvents()")
}

func (*Document) UpdateSettings() {
	js.Rewrite("$<.updateSettings()")
}

func (*Document) WebkitCancelFullScreen() {
	js.Rewrite("$<.webkitCancelFullScreen()")
}

func (*Document) WebkitExitFullscreen() {
	js.Rewrite("$<.webkitExitFullscreen()")
}

func (*Document) Write(content string) {
	js.Rewrite("$<.write($1)", content)
}

func (*Document) Writeln(content string) {
	js.Rewrite("$<.writeln($1)", content)
}

func (*Document) GetActiveElement() (activeElement *element.Element) {
	js.Rewrite("$<.activeElement")
	return activeElement
}

func (*Document) GetAlinkColor() (alinkColor string) {
	js.Rewrite("$<.alinkColor")
	return alinkColor
}

func (*Document) SetAlinkColor(alinkColor string) {
	js.Rewrite("$<.alinkColor = $1", alinkColor)
}

func (*Document) GetAll() (all *htmlallcollection.HTMLAllCollection) {
	js.Rewrite("$<.all")
	return all
}

func (*Document) GetAnchors() (anchors *htmlcollection.HTMLCollection) {
	js.Rewrite("$<.anchors")
	return anchors
}

func (*Document) GetApplets() (applets *htmlcollection.HTMLCollection) {
	js.Rewrite("$<.applets")
	return applets
}

func (*Document) GetBgColor() (bgColor string) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

func (*Document) SetBgColor(bgColor string) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

func (*Document) GetBody() (body *htmlelement.HTMLElement) {
	js.Rewrite("$<.body")
	return body
}

func (*Document) SetBody(body *htmlelement.HTMLElement) {
	js.Rewrite("$<.body = $1", body)
}

func (*Document) GetCharacterSet() (characterSet string) {
	js.Rewrite("$<.characterSet")
	return characterSet
}

func (*Document) GetCharset() (charset string) {
	js.Rewrite("$<.charset")
	return charset
}

func (*Document) SetCharset(charset string) {
	js.Rewrite("$<.charset = $1", charset)
}

func (*Document) GetCompatMode() (compatMode string) {
	js.Rewrite("$<.compatMode")
	return compatMode
}

func (*Document) GetCookie() (cookie string) {
	js.Rewrite("$<.cookie")
	return cookie
}

func (*Document) SetCookie(cookie string) {
	js.Rewrite("$<.cookie = $1", cookie)
}

func (*Document) GetCurrentScript() (currentScript interface{}) {
	js.Rewrite("$<.currentScript")
	return currentScript
}

func (*Document) GetDefaultView() (defaultView *window.Window) {
	js.Rewrite("$<.defaultView")
	return defaultView
}

func (*Document) GetDesignMode() (designMode string) {
	js.Rewrite("$<.designMode")
	return designMode
}

func (*Document) SetDesignMode(designMode string) {
	js.Rewrite("$<.designMode = $1", designMode)
}

func (*Document) GetDir() (dir string) {
	js.Rewrite("$<.dir")
	return dir
}

func (*Document) SetDir(dir string) {
	js.Rewrite("$<.dir = $1", dir)
}

func (*Document) GetDoctype() (doctype *documenttype.DocumentType) {
	js.Rewrite("$<.doctype")
	return doctype
}

func (*Document) GetDocumentElement() (documentElement *element.Element) {
	js.Rewrite("$<.documentElement")
	return documentElement
}

func (*Document) GetDomain() (domain string) {
	js.Rewrite("$<.domain")
	return domain
}

func (*Document) SetDomain(domain string) {
	js.Rewrite("$<.domain = $1", domain)
}

func (*Document) GetEmbeds() (embeds *htmlcollection.HTMLCollection) {
	js.Rewrite("$<.embeds")
	return embeds
}

func (*Document) GetFgColor() (fgColor string) {
	js.Rewrite("$<.fgColor")
	return fgColor
}

func (*Document) SetFgColor(fgColor string) {
	js.Rewrite("$<.fgColor = $1", fgColor)
}

func (*Document) GetForms() (forms *htmlcollection.HTMLCollection) {
	js.Rewrite("$<.forms")
	return forms
}

func (*Document) GetFullscreenElement() (fullscreenElement *element.Element) {
	js.Rewrite("$<.fullscreenElement")
	return fullscreenElement
}

func (*Document) GetFullscreenEnabled() (fullscreenEnabled bool) {
	js.Rewrite("$<.fullscreenEnabled")
	return fullscreenEnabled
}

func (*Document) GetHead() (head *htmlheadelement.HTMLHeadElement) {
	js.Rewrite("$<.head")
	return head
}

func (*Document) GetHidden() (hidden bool) {
	js.Rewrite("$<.hidden")
	return hidden
}

func (*Document) GetImages() (images *htmlcollection.HTMLCollection) {
	js.Rewrite("$<.images")
	return images
}

func (*Document) GetImplementation() (implementation *domimplementation.DOMImplementation) {
	js.Rewrite("$<.implementation")
	return implementation
}

func (*Document) GetInputEncoding() (inputEncoding string) {
	js.Rewrite("$<.inputEncoding")
	return inputEncoding
}

func (*Document) GetLastModified() (lastModified string) {
	js.Rewrite("$<.lastModified")
	return lastModified
}

func (*Document) GetLinkColor() (linkColor string) {
	js.Rewrite("$<.linkColor")
	return linkColor
}

func (*Document) SetLinkColor(linkColor string) {
	js.Rewrite("$<.linkColor = $1", linkColor)
}

func (*Document) GetLinks() (links *htmlcollection.HTMLCollection) {
	js.Rewrite("$<.links")
	return links
}

func (*Document) GetLocation() (location *location.Location) {
	js.Rewrite("$<.location")
	return location
}

func (*Document) GetMsCapsLockWarningOff() (msCapsLockWarningOff bool) {
	js.Rewrite("$<.msCapsLockWarningOff")
	return msCapsLockWarningOff
}

func (*Document) SetMsCapsLockWarningOff(msCapsLockWarningOff bool) {
	js.Rewrite("$<.msCapsLockWarningOff = $1", msCapsLockWarningOff)
}

func (*Document) GetMsCSSOMElementFloatMetrics() (msCSSOMElementFloatMetrics bool) {
	js.Rewrite("$<.msCSSOMElementFloatMetrics")
	return msCSSOMElementFloatMetrics
}

func (*Document) SetMsCSSOMElementFloatMetrics(msCSSOMElementFloatMetrics bool) {
	js.Rewrite("$<.msCSSOMElementFloatMetrics = $1", msCSSOMElementFloatMetrics)
}

func (*Document) GetOnabort() (e *event.Event) {
	js.Rewrite("$<.onabort")
	return e
}

func (*Document) SetOnabort(e *event.Event) {
	js.Rewrite("$<.onabort = $1", e)
}

func (*Document) GetOnactivate() (e *event.Event) {
	js.Rewrite("$<.onactivate")
	return e
}

func (*Document) SetOnactivate(e *event.Event) {
	js.Rewrite("$<.onactivate = $1", e)
}

func (*Document) GetOnbeforeactivate() (e *event.Event) {
	js.Rewrite("$<.onbeforeactivate")
	return e
}

func (*Document) SetOnbeforeactivate(e *event.Event) {
	js.Rewrite("$<.onbeforeactivate = $1", e)
}

func (*Document) GetOnbeforedeactivate() (e *event.Event) {
	js.Rewrite("$<.onbeforedeactivate")
	return e
}

func (*Document) SetOnbeforedeactivate(e *event.Event) {
	js.Rewrite("$<.onbeforedeactivate = $1", e)
}

func (*Document) GetOnblur() (e *event.Event) {
	js.Rewrite("$<.onblur")
	return e
}

func (*Document) SetOnblur(e *event.Event) {
	js.Rewrite("$<.onblur = $1", e)
}

func (*Document) GetOncanplay() (e *event.Event) {
	js.Rewrite("$<.oncanplay")
	return e
}

func (*Document) SetOncanplay(e *event.Event) {
	js.Rewrite("$<.oncanplay = $1", e)
}

func (*Document) GetOncanplaythrough() (e *event.Event) {
	js.Rewrite("$<.oncanplaythrough")
	return e
}

func (*Document) SetOncanplaythrough(e *event.Event) {
	js.Rewrite("$<.oncanplaythrough = $1", e)
}

func (*Document) GetOnchange() (e *event.Event) {
	js.Rewrite("$<.onchange")
	return e
}

func (*Document) SetOnchange(e *event.Event) {
	js.Rewrite("$<.onchange = $1", e)
}

func (*Document) GetOnclick() (e *event.Event) {
	js.Rewrite("$<.onclick")
	return e
}

func (*Document) SetOnclick(e *event.Event) {
	js.Rewrite("$<.onclick = $1", e)
}

func (*Document) GetOncontextmenu() (e *event.Event) {
	js.Rewrite("$<.oncontextmenu")
	return e
}

func (*Document) SetOncontextmenu(e *event.Event) {
	js.Rewrite("$<.oncontextmenu = $1", e)
}

func (*Document) GetOndblclick() (e *event.Event) {
	js.Rewrite("$<.ondblclick")
	return e
}

func (*Document) SetOndblclick(e *event.Event) {
	js.Rewrite("$<.ondblclick = $1", e)
}

func (*Document) GetOndeactivate() (e *event.Event) {
	js.Rewrite("$<.ondeactivate")
	return e
}

func (*Document) SetOndeactivate(e *event.Event) {
	js.Rewrite("$<.ondeactivate = $1", e)
}

func (*Document) GetOndrag() (e *event.Event) {
	js.Rewrite("$<.ondrag")
	return e
}

func (*Document) SetOndrag(e *event.Event) {
	js.Rewrite("$<.ondrag = $1", e)
}

func (*Document) GetOndragend() (e *event.Event) {
	js.Rewrite("$<.ondragend")
	return e
}

func (*Document) SetOndragend(e *event.Event) {
	js.Rewrite("$<.ondragend = $1", e)
}

func (*Document) GetOndragenter() (e *event.Event) {
	js.Rewrite("$<.ondragenter")
	return e
}

func (*Document) SetOndragenter(e *event.Event) {
	js.Rewrite("$<.ondragenter = $1", e)
}

func (*Document) GetOndragleave() (e *event.Event) {
	js.Rewrite("$<.ondragleave")
	return e
}

func (*Document) SetOndragleave(e *event.Event) {
	js.Rewrite("$<.ondragleave = $1", e)
}

func (*Document) GetOndragover() (e *event.Event) {
	js.Rewrite("$<.ondragover")
	return e
}

func (*Document) SetOndragover(e *event.Event) {
	js.Rewrite("$<.ondragover = $1", e)
}

func (*Document) GetOndragstart() (e *event.Event) {
	js.Rewrite("$<.ondragstart")
	return e
}

func (*Document) SetOndragstart(e *event.Event) {
	js.Rewrite("$<.ondragstart = $1", e)
}

func (*Document) GetOndrop() (e *event.Event) {
	js.Rewrite("$<.ondrop")
	return e
}

func (*Document) SetOndrop(e *event.Event) {
	js.Rewrite("$<.ondrop = $1", e)
}

func (*Document) GetOndurationchange() (e *event.Event) {
	js.Rewrite("$<.ondurationchange")
	return e
}

func (*Document) SetOndurationchange(e *event.Event) {
	js.Rewrite("$<.ondurationchange = $1", e)
}

func (*Document) GetOnemptied() (e *event.Event) {
	js.Rewrite("$<.onemptied")
	return e
}

func (*Document) SetOnemptied(e *event.Event) {
	js.Rewrite("$<.onemptied = $1", e)
}

func (*Document) GetOnended() (e *event.Event) {
	js.Rewrite("$<.onended")
	return e
}

func (*Document) SetOnended(e *event.Event) {
	js.Rewrite("$<.onended = $1", e)
}

func (*Document) GetOnerror() (e *event.Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*Document) SetOnerror(e *event.Event) {
	js.Rewrite("$<.onerror = $1", e)
}

func (*Document) GetOnfocus() (e *event.Event) {
	js.Rewrite("$<.onfocus")
	return e
}

func (*Document) SetOnfocus(e *event.Event) {
	js.Rewrite("$<.onfocus = $1", e)
}

func (*Document) GetOnfullscreenchange() (fullscreenchange *event.Event) {
	js.Rewrite("$<.onfullscreenchange")
	return fullscreenchange
}

func (*Document) SetOnfullscreenchange(fullscreenchange *event.Event) {
	js.Rewrite("$<.onfullscreenchange = $1", fullscreenchange)
}

func (*Document) GetOnfullscreenerror() (fullscreenerror *event.Event) {
	js.Rewrite("$<.onfullscreenerror")
	return fullscreenerror
}

func (*Document) SetOnfullscreenerror(fullscreenerror *event.Event) {
	js.Rewrite("$<.onfullscreenerror = $1", fullscreenerror)
}

func (*Document) GetOninput() (e *event.Event) {
	js.Rewrite("$<.oninput")
	return e
}

func (*Document) SetOninput(e *event.Event) {
	js.Rewrite("$<.oninput = $1", e)
}

func (*Document) GetOninvalid() (e *event.Event) {
	js.Rewrite("$<.oninvalid")
	return e
}

func (*Document) SetOninvalid(e *event.Event) {
	js.Rewrite("$<.oninvalid = $1", e)
}

func (*Document) GetOnkeydown() (e *event.Event) {
	js.Rewrite("$<.onkeydown")
	return e
}

func (*Document) SetOnkeydown(e *event.Event) {
	js.Rewrite("$<.onkeydown = $1", e)
}

func (*Document) GetOnkeypress() (e *event.Event) {
	js.Rewrite("$<.onkeypress")
	return e
}

func (*Document) SetOnkeypress(e *event.Event) {
	js.Rewrite("$<.onkeypress = $1", e)
}

func (*Document) GetOnkeyup() (e *event.Event) {
	js.Rewrite("$<.onkeyup")
	return e
}

func (*Document) SetOnkeyup(e *event.Event) {
	js.Rewrite("$<.onkeyup = $1", e)
}

func (*Document) GetOnload() (e *event.Event) {
	js.Rewrite("$<.onload")
	return e
}

func (*Document) SetOnload(e *event.Event) {
	js.Rewrite("$<.onload = $1", e)
}

func (*Document) GetOnloadeddata() (e *event.Event) {
	js.Rewrite("$<.onloadeddata")
	return e
}

func (*Document) SetOnloadeddata(e *event.Event) {
	js.Rewrite("$<.onloadeddata = $1", e)
}

func (*Document) GetOnloadedmetadata() (e *event.Event) {
	js.Rewrite("$<.onloadedmetadata")
	return e
}

func (*Document) SetOnloadedmetadata(e *event.Event) {
	js.Rewrite("$<.onloadedmetadata = $1", e)
}

func (*Document) GetOnloadstart() (e *event.Event) {
	js.Rewrite("$<.onloadstart")
	return e
}

func (*Document) SetOnloadstart(e *event.Event) {
	js.Rewrite("$<.onloadstart = $1", e)
}

func (*Document) GetOnmousedown() (e *event.Event) {
	js.Rewrite("$<.onmousedown")
	return e
}

func (*Document) SetOnmousedown(e *event.Event) {
	js.Rewrite("$<.onmousedown = $1", e)
}

func (*Document) GetOnmousemove() (e *event.Event) {
	js.Rewrite("$<.onmousemove")
	return e
}

func (*Document) SetOnmousemove(e *event.Event) {
	js.Rewrite("$<.onmousemove = $1", e)
}

func (*Document) GetOnmouseout() (e *event.Event) {
	js.Rewrite("$<.onmouseout")
	return e
}

func (*Document) SetOnmouseout(e *event.Event) {
	js.Rewrite("$<.onmouseout = $1", e)
}

func (*Document) GetOnmouseover() (e *event.Event) {
	js.Rewrite("$<.onmouseover")
	return e
}

func (*Document) SetOnmouseover(e *event.Event) {
	js.Rewrite("$<.onmouseover = $1", e)
}

func (*Document) GetOnmouseup() (e *event.Event) {
	js.Rewrite("$<.onmouseup")
	return e
}

func (*Document) SetOnmouseup(e *event.Event) {
	js.Rewrite("$<.onmouseup = $1", e)
}

func (*Document) GetOnmousewheel() (e *event.Event) {
	js.Rewrite("$<.onmousewheel")
	return e
}

func (*Document) SetOnmousewheel(e *event.Event) {
	js.Rewrite("$<.onmousewheel = $1", e)
}

func (*Document) GetOnmscontentzoom() (MSContentZoom *uievent.UIEvent) {
	js.Rewrite("$<.onmscontentzoom")
	return MSContentZoom
}

func (*Document) SetOnmscontentzoom(MSContentZoom *uievent.UIEvent) {
	js.Rewrite("$<.onmscontentzoom = $1", MSContentZoom)
}

func (*Document) GetOnmsgesturechange() (e *event.Event) {
	js.Rewrite("$<.onmsgesturechange")
	return e
}

func (*Document) SetOnmsgesturechange(e *event.Event) {
	js.Rewrite("$<.onmsgesturechange = $1", e)
}

func (*Document) GetOnmsgesturedoubletap() (e *event.Event) {
	js.Rewrite("$<.onmsgesturedoubletap")
	return e
}

func (*Document) SetOnmsgesturedoubletap(e *event.Event) {
	js.Rewrite("$<.onmsgesturedoubletap = $1", e)
}

func (*Document) GetOnmsgestureend() (e *event.Event) {
	js.Rewrite("$<.onmsgestureend")
	return e
}

func (*Document) SetOnmsgestureend(e *event.Event) {
	js.Rewrite("$<.onmsgestureend = $1", e)
}

func (*Document) GetOnmsgesturehold() (e *event.Event) {
	js.Rewrite("$<.onmsgesturehold")
	return e
}

func (*Document) SetOnmsgesturehold(e *event.Event) {
	js.Rewrite("$<.onmsgesturehold = $1", e)
}

func (*Document) GetOnmsgesturestart() (e *event.Event) {
	js.Rewrite("$<.onmsgesturestart")
	return e
}

func (*Document) SetOnmsgesturestart(e *event.Event) {
	js.Rewrite("$<.onmsgesturestart = $1", e)
}

func (*Document) GetOnmsgesturetap() (e *event.Event) {
	js.Rewrite("$<.onmsgesturetap")
	return e
}

func (*Document) SetOnmsgesturetap(e *event.Event) {
	js.Rewrite("$<.onmsgesturetap = $1", e)
}

func (*Document) GetOnmsinertiastart() (e *event.Event) {
	js.Rewrite("$<.onmsinertiastart")
	return e
}

func (*Document) SetOnmsinertiastart(e *event.Event) {
	js.Rewrite("$<.onmsinertiastart = $1", e)
}

func (*Document) GetOnmsmanipulationstatechanged() (MSManipulationStateChanged *msmanipulationevent.MSManipulationEvent) {
	js.Rewrite("$<.onmsmanipulationstatechanged")
	return MSManipulationStateChanged
}

func (*Document) SetOnmsmanipulationstatechanged(MSManipulationStateChanged *msmanipulationevent.MSManipulationEvent) {
	js.Rewrite("$<.onmsmanipulationstatechanged = $1", MSManipulationStateChanged)
}

func (*Document) GetOnmspointercancel() (e *event.Event) {
	js.Rewrite("$<.onmspointercancel")
	return e
}

func (*Document) SetOnmspointercancel(e *event.Event) {
	js.Rewrite("$<.onmspointercancel = $1", e)
}

func (*Document) GetOnmspointerdown() (e *event.Event) {
	js.Rewrite("$<.onmspointerdown")
	return e
}

func (*Document) SetOnmspointerdown(e *event.Event) {
	js.Rewrite("$<.onmspointerdown = $1", e)
}

func (*Document) GetOnmspointerenter() (e *event.Event) {
	js.Rewrite("$<.onmspointerenter")
	return e
}

func (*Document) SetOnmspointerenter(e *event.Event) {
	js.Rewrite("$<.onmspointerenter = $1", e)
}

func (*Document) GetOnmspointerleave() (e *event.Event) {
	js.Rewrite("$<.onmspointerleave")
	return e
}

func (*Document) SetOnmspointerleave(e *event.Event) {
	js.Rewrite("$<.onmspointerleave = $1", e)
}

func (*Document) GetOnmspointermove() (e *event.Event) {
	js.Rewrite("$<.onmspointermove")
	return e
}

func (*Document) SetOnmspointermove(e *event.Event) {
	js.Rewrite("$<.onmspointermove = $1", e)
}

func (*Document) GetOnmspointerout() (e *event.Event) {
	js.Rewrite("$<.onmspointerout")
	return e
}

func (*Document) SetOnmspointerout(e *event.Event) {
	js.Rewrite("$<.onmspointerout = $1", e)
}

func (*Document) GetOnmspointerover() (e *event.Event) {
	js.Rewrite("$<.onmspointerover")
	return e
}

func (*Document) SetOnmspointerover(e *event.Event) {
	js.Rewrite("$<.onmspointerover = $1", e)
}

func (*Document) GetOnmspointerup() (e *event.Event) {
	js.Rewrite("$<.onmspointerup")
	return e
}

func (*Document) SetOnmspointerup(e *event.Event) {
	js.Rewrite("$<.onmspointerup = $1", e)
}

func (*Document) GetOnmssitemodejumplistitemremoved() (mssitemodejumplistitemremoved *mssitemodeevent.MSSiteModeEvent) {
	js.Rewrite("$<.onmssitemodejumplistitemremoved")
	return mssitemodejumplistitemremoved
}

func (*Document) SetOnmssitemodejumplistitemremoved(mssitemodejumplistitemremoved *mssitemodeevent.MSSiteModeEvent) {
	js.Rewrite("$<.onmssitemodejumplistitemremoved = $1", mssitemodejumplistitemremoved)
}

func (*Document) GetOnmsthumbnailclick() (msthumbnailclick *mssitemodeevent.MSSiteModeEvent) {
	js.Rewrite("$<.onmsthumbnailclick")
	return msthumbnailclick
}

func (*Document) SetOnmsthumbnailclick(msthumbnailclick *mssitemodeevent.MSSiteModeEvent) {
	js.Rewrite("$<.onmsthumbnailclick = $1", msthumbnailclick)
}

func (*Document) GetOnpause() (e *event.Event) {
	js.Rewrite("$<.onpause")
	return e
}

func (*Document) SetOnpause(e *event.Event) {
	js.Rewrite("$<.onpause = $1", e)
}

func (*Document) GetOnplay() (e *event.Event) {
	js.Rewrite("$<.onplay")
	return e
}

func (*Document) SetOnplay(e *event.Event) {
	js.Rewrite("$<.onplay = $1", e)
}

func (*Document) GetOnplaying() (e *event.Event) {
	js.Rewrite("$<.onplaying")
	return e
}

func (*Document) SetOnplaying(e *event.Event) {
	js.Rewrite("$<.onplaying = $1", e)
}

func (*Document) GetOnpointerlockchange() (e *event.Event) {
	js.Rewrite("$<.onpointerlockchange")
	return e
}

func (*Document) SetOnpointerlockchange(e *event.Event) {
	js.Rewrite("$<.onpointerlockchange = $1", e)
}

func (*Document) GetOnpointerlockerror() (e *event.Event) {
	js.Rewrite("$<.onpointerlockerror")
	return e
}

func (*Document) SetOnpointerlockerror(e *event.Event) {
	js.Rewrite("$<.onpointerlockerror = $1", e)
}

func (*Document) GetOnprogress() (e *event.Event) {
	js.Rewrite("$<.onprogress")
	return e
}

func (*Document) SetOnprogress(e *event.Event) {
	js.Rewrite("$<.onprogress = $1", e)
}

func (*Document) GetOnratechange() (e *event.Event) {
	js.Rewrite("$<.onratechange")
	return e
}

func (*Document) SetOnratechange(e *event.Event) {
	js.Rewrite("$<.onratechange = $1", e)
}

func (*Document) GetOnreadystatechange() (readystatechange *event.Event) {
	js.Rewrite("$<.onreadystatechange")
	return readystatechange
}

func (*Document) SetOnreadystatechange(readystatechange *event.Event) {
	js.Rewrite("$<.onreadystatechange = $1", readystatechange)
}

func (*Document) GetOnreset() (e *event.Event) {
	js.Rewrite("$<.onreset")
	return e
}

func (*Document) SetOnreset(e *event.Event) {
	js.Rewrite("$<.onreset = $1", e)
}

func (*Document) GetOnscroll() (e *event.Event) {
	js.Rewrite("$<.onscroll")
	return e
}

func (*Document) SetOnscroll(e *event.Event) {
	js.Rewrite("$<.onscroll = $1", e)
}

func (*Document) GetOnseeked() (e *event.Event) {
	js.Rewrite("$<.onseeked")
	return e
}

func (*Document) SetOnseeked(e *event.Event) {
	js.Rewrite("$<.onseeked = $1", e)
}

func (*Document) GetOnseeking() (e *event.Event) {
	js.Rewrite("$<.onseeking")
	return e
}

func (*Document) SetOnseeking(e *event.Event) {
	js.Rewrite("$<.onseeking = $1", e)
}

func (*Document) GetOnselect() (e *event.Event) {
	js.Rewrite("$<.onselect")
	return e
}

func (*Document) SetOnselect(e *event.Event) {
	js.Rewrite("$<.onselect = $1", e)
}

func (*Document) GetOnselectionchange() (selectionchange *event.Event) {
	js.Rewrite("$<.onselectionchange")
	return selectionchange
}

func (*Document) SetOnselectionchange(selectionchange *event.Event) {
	js.Rewrite("$<.onselectionchange = $1", selectionchange)
}

func (*Document) GetOnselectstart() (e *event.Event) {
	js.Rewrite("$<.onselectstart")
	return e
}

func (*Document) SetOnselectstart(e *event.Event) {
	js.Rewrite("$<.onselectstart = $1", e)
}

func (*Document) GetOnstalled() (e *event.Event) {
	js.Rewrite("$<.onstalled")
	return e
}

func (*Document) SetOnstalled(e *event.Event) {
	js.Rewrite("$<.onstalled = $1", e)
}

func (*Document) GetOnstop() (stop *event.Event) {
	js.Rewrite("$<.onstop")
	return stop
}

func (*Document) SetOnstop(stop *event.Event) {
	js.Rewrite("$<.onstop = $1", stop)
}

func (*Document) GetOnsubmit() (e *event.Event) {
	js.Rewrite("$<.onsubmit")
	return e
}

func (*Document) SetOnsubmit(e *event.Event) {
	js.Rewrite("$<.onsubmit = $1", e)
}

func (*Document) GetOnsuspend() (e *event.Event) {
	js.Rewrite("$<.onsuspend")
	return e
}

func (*Document) SetOnsuspend(e *event.Event) {
	js.Rewrite("$<.onsuspend = $1", e)
}

func (*Document) GetOntimeupdate() (e *event.Event) {
	js.Rewrite("$<.ontimeupdate")
	return e
}

func (*Document) SetOntimeupdate(e *event.Event) {
	js.Rewrite("$<.ontimeupdate = $1", e)
}

func (*Document) GetOntouchcancel() (e *event.Event) {
	js.Rewrite("$<.ontouchcancel")
	return e
}

func (*Document) SetOntouchcancel(e *event.Event) {
	js.Rewrite("$<.ontouchcancel = $1", e)
}

func (*Document) GetOntouchend() (e *event.Event) {
	js.Rewrite("$<.ontouchend")
	return e
}

func (*Document) SetOntouchend(e *event.Event) {
	js.Rewrite("$<.ontouchend = $1", e)
}

func (*Document) GetOntouchmove() (e *event.Event) {
	js.Rewrite("$<.ontouchmove")
	return e
}

func (*Document) SetOntouchmove(e *event.Event) {
	js.Rewrite("$<.ontouchmove = $1", e)
}

func (*Document) GetOntouchstart() (e *event.Event) {
	js.Rewrite("$<.ontouchstart")
	return e
}

func (*Document) SetOntouchstart(e *event.Event) {
	js.Rewrite("$<.ontouchstart = $1", e)
}

func (*Document) GetOnvolumechange() (e *event.Event) {
	js.Rewrite("$<.onvolumechange")
	return e
}

func (*Document) SetOnvolumechange(e *event.Event) {
	js.Rewrite("$<.onvolumechange = $1", e)
}

func (*Document) GetOnwaiting() (e *event.Event) {
	js.Rewrite("$<.onwaiting")
	return e
}

func (*Document) SetOnwaiting(e *event.Event) {
	js.Rewrite("$<.onwaiting = $1", e)
}

func (*Document) GetOnwebkitfullscreenchange() (e *event.Event) {
	js.Rewrite("$<.onwebkitfullscreenchange")
	return e
}

func (*Document) SetOnwebkitfullscreenchange(e *event.Event) {
	js.Rewrite("$<.onwebkitfullscreenchange = $1", e)
}

func (*Document) GetOnwebkitfullscreenerror() (e *event.Event) {
	js.Rewrite("$<.onwebkitfullscreenerror")
	return e
}

func (*Document) SetOnwebkitfullscreenerror(e *event.Event) {
	js.Rewrite("$<.onwebkitfullscreenerror = $1", e)
}

func (*Document) GetPlugins() (plugins *htmlcollection.HTMLCollection) {
	js.Rewrite("$<.plugins")
	return plugins
}

func (*Document) GetPointerLockElement() (pointerLockElement *element.Element) {
	js.Rewrite("$<.pointerLockElement")
	return pointerLockElement
}

func (*Document) GetReadyState() (readyState string) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*Document) GetReferrer() (referrer string) {
	js.Rewrite("$<.referrer")
	return referrer
}

func (*Document) GetRootElement() (rootElement *svgsvgelement.SVGSVGElement) {
	js.Rewrite("$<.rootElement")
	return rootElement
}

func (*Document) GetScripts() (scripts *htmlcollection.HTMLCollection) {
	js.Rewrite("$<.scripts")
	return scripts
}

func (*Document) GetScrollingElement() (scrollingElement *element.Element) {
	js.Rewrite("$<.scrollingElement")
	return scrollingElement
}

func (*Document) GetStyleSheets() (styleSheets *stylesheetlist.StyleSheetList) {
	js.Rewrite("$<.styleSheets")
	return styleSheets
}

func (*Document) GetTitle() (title string) {
	js.Rewrite("$<.title")
	return title
}

func (*Document) SetTitle(title string) {
	js.Rewrite("$<.title = $1", title)
}

func (*Document) GetURL() (URL string) {
	js.Rewrite("$<.URL")
	return URL
}

func (*Document) GetURLUnencoded() (URLUnencoded string) {
	js.Rewrite("$<.URLUnencoded")
	return URLUnencoded
}

func (*Document) GetVisibilityState() (visibilityState *visibilitystate.VisibilityState) {
	js.Rewrite("$<.visibilityState")
	return visibilityState
}

func (*Document) GetVlinkColor() (vlinkColor string) {
	js.Rewrite("$<.vlinkColor")
	return vlinkColor
}

func (*Document) SetVlinkColor(vlinkColor string) {
	js.Rewrite("$<.vlinkColor = $1", vlinkColor)
}

func (*Document) GetWebkitCurrentFullScreenElement() (webkitCurrentFullScreenElement *element.Element) {
	js.Rewrite("$<.webkitCurrentFullScreenElement")
	return webkitCurrentFullScreenElement
}

func (*Document) GetWebkitFullscreenElement() (webkitFullscreenElement *element.Element) {
	js.Rewrite("$<.webkitFullscreenElement")
	return webkitFullscreenElement
}

func (*Document) GetWebkitFullscreenEnabled() (webkitFullscreenEnabled bool) {
	js.Rewrite("$<.webkitFullscreenEnabled")
	return webkitFullscreenEnabled
}

func (*Document) GetWebkitIsFullScreen() (webkitIsFullScreen bool) {
	js.Rewrite("$<.webkitIsFullScreen")
	return webkitIsFullScreen
}

func (*Document) GetXMLEncoding() (xmlEncoding string) {
	js.Rewrite("$<.xmlEncoding")
	return xmlEncoding
}

func (*Document) GetXMLStandalone() (xmlStandalone bool) {
	js.Rewrite("$<.xmlStandalone")
	return xmlStandalone
}

func (*Document) SetXMLStandalone(xmlStandalone bool) {
	js.Rewrite("$<.xmlStandalone = $1", xmlStandalone)
}

func (*Document) GetXMLVersion() (xmlVersion string) {
	js.Rewrite("$<.xmlVersion")
	return xmlVersion
}

func (*Document) SetXMLVersion(xmlVersion string) {
	js.Rewrite("$<.xmlVersion = $1", xmlVersion)
}
