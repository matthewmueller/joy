import {
  assign,
  toLower,
  splice,
  findWhere,
  createAttributeFilter
} from './util'

/*
const NODE_TYPES = {
	ELEMENT_NODE: 1,
	ATTRIBUTE_NODE: 2,
	TEXT_NODE: 3,
	CDATA_SECTION_NODE: 4,
	ENTITY_REFERENCE_NODE: 5,
	COMMENT_NODE: 6,
	PROCESSING_INSTRUCTION_NODE: 7,
	DOCUMENT_NODE: 9
};
*/

/** Create a minimally viable DOM Document
 *	@returns {Document} document
 */
export default function undom() {
  function isElement(node) {
    return node.nodeType === 1
  }

  class Node {
    constructor(nodeType, nodeName) {
      this.nodeType = nodeType
      this.nodeName = nodeName
      this.childNodes = []
    }
    get nextSibling() {
      let p = this.parentNode
      if (p) return p.childNodes[findWhere(p.childNodes, this, true) + 1]
    }
    get previousSibling() {
      let p = this.parentNode
      if (p) return p.childNodes[findWhere(p.childNodes, this, true) - 1]
    }
    get firstChild() {
      return this.childNodes[0]
    }
    get lastChild() {
      return this.childNodes[this.childNodes.length - 1]
    }
    appendChild(child) {
      this.insertBefore(child)
    }
    insertBefore(child, ref) {
      child.remove()
      child.parentNode = this
      if (!ref) this.childNodes.push(child)
      else splice(this.childNodes, ref, child)
    }
    replaceChild(child, ref) {
      if (ref.parentNode === this) {
        this.insertBefore(child, ref)
        ref.remove()
      }
    }
    removeChild(child) {
      splice(this.childNodes, child)
    }
    remove() {
      if (this.parentNode) this.parentNode.removeChild(this)
    }
  }

  class Text extends Node {
    constructor(text) {
      super(3, '#text') // TEXT_NODE
      this.nodeValue = text
    }
    set textContent(text) {
      this.nodeValue = text
    }
    get textContent() {
      return this.nodeValue
    }
  }

  class Element extends Node {
    constructor(nodeType, nodeName) {
      super(nodeType || 1, nodeName) // ELEMENT_NODE
      this.attributes = []
      this.__handlers = {}
      this.style = {}
      Object.defineProperty(this, 'className', {
        set: val => {
          this.setAttribute('class', val)
        },
        get: () => this.getAttribute('class')
      })
      Object.defineProperty(this.style, 'cssText', {
        set: val => {
          this.setAttribute('style', val)
        },
        get: () => this.getAttribute('style')
      })
    }

    get children() {
      return this.childNodes.filter(isElement)
    }

    setAttribute(key, value) {
      this.setAttributeNS(null, key, value)
    }
    getAttribute(key) {
      return this.getAttributeNS(null, key)
    }
    removeAttribute(key) {
      this.removeAttributeNS(null, key)
    }

    setAttributeNS(ns, name, value) {
      let attr = findWhere(this.attributes, createAttributeFilter(ns, name))
      if (!attr) this.attributes.push((attr = { ns, name }))
      attr.value = String(value)
    }
    getAttributeNS(ns, name) {
      let attr = findWhere(this.attributes, createAttributeFilter(ns, name))
      return attr && attr.value
    }
    removeAttributeNS(ns, name) {
      splice(this.attributes, createAttributeFilter(ns, name))
    }

    addEventListener(type, handler) {
      ;(this.__handlers[toLower(type)] || (this.__handlers[toLower(type)] = []))
        .push(handler)
    }
    removeEventListener(type, handler) {
      splice(this.__handlers[toLower(type)], handler, 0, true)
    }
    dispatchEvent(event) {
      let t = (event.currentTarget = this),
        c = event.cancelable,
        l,
        i
      do {
        l = t.__handlers && t.__handlers[toLower(event.type)]
        if (l)
          for (i = l.length; i--; ) {
            if ((l[i].call(t, event) === false || event._end) && c) break
          }
      } while (
        event.bubbles &&
        !(c && event._stop) &&
        (event.target = t = t.parentNode)
      )
      return !event.defaultPrevented
    }
  }

  class Document extends Element {
    constructor() {
      super(9, '#document') // DOCUMENT_NODE
    }
  }

  class Event {
    constructor(type, opts) {
      this.type = type
      this.bubbles = !!opts.bubbles
      this.cancelable = !!opts.cancelable
    }
    stopPropagation() {
      this._stop = true
    }
    stopImmediatePropagation() {
      this._end = this._stop = true
    }
    preventDefault() {
      this.defaultPrevented = true
    }
  }

  function createElement(type) {
    return new Element(null, String(type).toUpperCase())
  }

  function createElementNS(ns, type) {
    let element = createElement(type)
    element.namespace = ns
    return element
  }

  function createTextNode(text) {
    return new Text(text)
  }

  function createDocument() {
    let document = new Document()
    assign(
      document,
      (document.defaultView = {
        document,
        Document,
        Node,
        Text,
        Element,
        SVGElement: Element,
        Event
      })
    )
    assign(document, { createElement, createElementNS, createTextNode })
    document.appendChild((document.documentElement = createElement('html')))
    document.documentElement.appendChild(
      (document.head = createElement('head'))
    )
    document.documentElement.appendChild(
      (document.body = createElement('body'))
    )
    return document
  }

  return createDocument()
}
