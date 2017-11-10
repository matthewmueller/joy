;(function() {
  var pkg = {}
  pkg[
    'github.com/matthewmueller/golly/testdata/56-hn-preact/preact/preact.js'
  ] = (function() {
    !(function() {
      'use strict'

      function VNode() {}
      function h(nodeName, attributes) {
        var lastSimple,
          child,
          simple,
          i,
          children = EMPTY_CHILDREN
        for (i = arguments.length; i-- > 2; ) {
          stack.push(arguments[i])
        }
        if (attributes && null != attributes.children) {
          if (!stack.length) stack.push(attributes.children)
          delete attributes.children
        }
        while (stack.length) {
          if ((child = stack.pop()) && void 0 !== child.pop)
            for (i = child.length; i--; ) {
              stack.push(child[i])
            }
          else {
            if ('boolean' == typeof child) child = null
            if ((simple = 'function' != typeof nodeName))
              if (null == child) child = ''
              else if ('number' == typeof child) child = String(child)
              else if ('string' != typeof child) simple = !1
            if (simple && lastSimple) children[children.length - 1] += child
            else if (children === EMPTY_CHILDREN) children = [child]
            else children.push(child)
            lastSimple = simple
          }
        }
        var p = new VNode()
        p.nodeName = nodeName
        p.children = children
        p.attributes = null == attributes ? void 0 : attributes
        p.key = null == attributes ? void 0 : attributes.key
        if (void 0 !== options.vnode) options.vnode(p)
        return p
      }
      function extend(obj, props) {
        for (var i in props) {
          obj[i] = props[i]
        }
        return obj
      }
      function cloneElement(vnode, props) {
        return h(
          vnode.nodeName,
          extend(extend({}, vnode.attributes), props),
          arguments.length > 2 ? [].slice.call(arguments, 2) : vnode.children
        )
      }
      function enqueueRender(component) {
        if (
          !component.__d &&
          (component.__d = !0) &&
          1 == items.push(component)
        )
          (options.debounceRendering || defer)(rerender)
      }
      function rerender() {
        var p,
          list = items
        items = []
        while ((p = list.pop())) {
          if (p.__d) renderComponent(p)
        }
      }
      function isSameNodeType(node, vnode, hydrating) {
        if ('string' == typeof vnode || 'number' == typeof vnode)
          return void 0 !== node.splitText
        if ('string' == typeof vnode.nodeName)
          return (
            !node._componentConstructor && isNamedNode(node, vnode.nodeName)
          )
        else return hydrating || node._componentConstructor === vnode.nodeName
      }
      function isNamedNode(node, nodeName) {
        return (
          node.__n === nodeName ||
          node.nodeName.toLowerCase() === nodeName.toLowerCase()
        )
      }
      function getNodeProps(vnode) {
        var props = extend({}, vnode.attributes)
        props.children = vnode.children
        var defaultProps = vnode.nodeName.defaultProps
        if (void 0 !== defaultProps)
          for (var i in defaultProps) {
            if (void 0 === props[i]) props[i] = defaultProps[i]
          }
        return props
      }
      function createNode(nodeName, isSvg) {
        var node = isSvg
          ? document.createElementNS('http://www.w3.org/2000/svg', nodeName)
          : document.createElement(nodeName)
        node.__n = nodeName
        return node
      }
      function removeNode(node) {
        var parentNode = node.parentNode
        if (parentNode) parentNode.removeChild(node)
      }
      function setAccessor(node, name, old, value, isSvg) {
        if ('className' === name) name = 'class'
        if ('key' === name);
        else if ('ref' === name) {
          if (old) old(null)
          if (value) value(node)
        } else if ('class' === name && !isSvg) node.className = value || ''
        else if ('style' === name) {
          if (!value || 'string' == typeof value || 'string' == typeof old)
            node.style.cssText = value || ''
          if (value && 'object' == typeof value) {
            if ('string' != typeof old)
              for (var i in old) {
                if (!(i in value)) node.style[i] = ''
              }
            for (var i in value) {
              node.style[i] =
                'number' == typeof value[i] && !1 === IS_NON_DIMENSIONAL.test(i)
                  ? value[i] + 'px'
                  : value[i]
            }
          }
        } else if ('dangerouslySetInnerHTML' === name) {
          if (value) node.innerHTML = value.__html || ''
        } else if ('o' == name[0] && 'n' == name[1]) {
          var useCapture = name !== (name = name.replace(/Capture$/, ''))
          name = name.toLowerCase().substring(2)
          if (value) {
            if (!old) node.addEventListener(name, eventProxy, useCapture)
          } else node.removeEventListener(name, eventProxy, useCapture)
          ;(node.__l || (node.__l = {}))[name] = value
        } else if (
          'list' !== name &&
          'type' !== name &&
          !isSvg &&
          name in node
        ) {
          setProperty(node, name, null == value ? '' : value)
          if (null == value || !1 === value) node.removeAttribute(name)
        } else {
          var ns = isSvg && name !== (name = name.replace(/^xlink\:?/, ''))
          if (null == value || !1 === value) {
            if (ns)
              node.removeAttributeNS(
                'http://www.w3.org/1999/xlink',
                name.toLowerCase()
              )
            else node.removeAttribute(name)
          } else if ('function' != typeof value)
            if (ns)
              node.setAttributeNS(
                'http://www.w3.org/1999/xlink',
                name.toLowerCase(),
                value
              )
            else node.setAttribute(name, value)
        }
      }
      function setProperty(node, name, value) {
        try {
          node[name] = value
        } catch (e) {}
      }
      function eventProxy(e) {
        return this.__l[e.type]((options.event && options.event(e)) || e)
      }
      function flushMounts() {
        var c
        while ((c = mounts.pop())) {
          if (options.afterMount) options.afterMount(c)
          if (c.componentDidMount) c.componentDidMount()
        }
      }
      function diff(dom, vnode, context, mountAll, parent, componentRoot) {
        if (!diffLevel++) {
          isSvgMode = null != parent && void 0 !== parent.ownerSVGElement
          hydrating = null != dom && !('__preactattr_' in dom)
        }
        var ret = idiff(dom, vnode, context, mountAll, componentRoot)
        if (parent && ret.parentNode !== parent) parent.appendChild(ret)
        if (!--diffLevel) {
          hydrating = !1
          if (!componentRoot) flushMounts()
        }
        return ret
      }
      function idiff(dom, vnode, context, mountAll, componentRoot) {
        var out = dom,
          prevSvgMode = isSvgMode
        if (null == vnode || 'boolean' == typeof vnode) vnode = ''
        if ('string' == typeof vnode || 'number' == typeof vnode) {
          if (
            dom &&
            void 0 !== dom.splitText &&
            dom.parentNode &&
            (!dom._component || componentRoot)
          ) {
            if (dom.nodeValue != vnode) dom.nodeValue = vnode
          } else {
            out = document.createTextNode(vnode)
            if (dom) {
              if (dom.parentNode) dom.parentNode.replaceChild(out, dom)
              recollectNodeTree(dom, !0)
            }
          }
          out.__preactattr_ = !0
          return out
        }
        var vnodeName = vnode.nodeName
        if ('function' == typeof vnodeName)
          return buildComponentFromVNode(dom, vnode, context, mountAll)
        isSvgMode =
          'svg' === vnodeName
            ? !0
            : 'foreignObject' === vnodeName ? !1 : isSvgMode
        vnodeName = String(vnodeName)
        if (!dom || !isNamedNode(dom, vnodeName)) {
          out = createNode(vnodeName, isSvgMode)
          if (dom) {
            while (dom.firstChild) {
              out.appendChild(dom.firstChild)
            }
            if (dom.parentNode) dom.parentNode.replaceChild(out, dom)
            recollectNodeTree(dom, !0)
          }
        }
        var fc = out.firstChild,
          props = out.__preactattr_,
          vchildren = vnode.children
        if (null == props) {
          props = out.__preactattr_ = {}
          for (var a = out.attributes, i = a.length; i--; ) {
            props[a[i].name] = a[i].value
          }
        }
        if (
          !hydrating &&
          vchildren &&
          1 === vchildren.length &&
          'string' == typeof vchildren[0] &&
          null != fc &&
          void 0 !== fc.splitText &&
          null == fc.nextSibling
        ) {
          if (fc.nodeValue != vchildren[0]) fc.nodeValue = vchildren[0]
        } else if ((vchildren && vchildren.length) || null != fc)
          innerDiffNode(
            out,
            vchildren,
            context,
            mountAll,
            hydrating || null != props.dangerouslySetInnerHTML
          )
        diffAttributes(out, vnode.attributes, props)
        isSvgMode = prevSvgMode
        return out
      }
      function innerDiffNode(dom, vchildren, context, mountAll, isHydrating) {
        var j,
          c,
          f,
          vchild,
          child,
          originalChildren = dom.childNodes,
          children = [],
          keyed = {},
          keyedLen = 0,
          min = 0,
          len = originalChildren.length,
          childrenLen = 0,
          vlen = vchildren ? vchildren.length : 0
        if (0 !== len)
          for (var i = 0; i < len; i++) {
            var _child = originalChildren[i],
              props = _child.__preactattr_,
              key =
                vlen && props
                  ? _child._component ? _child._component.__k : props.key
                  : null
            if (null != key) {
              keyedLen++
              keyed[key] = _child
            } else if (
              props ||
              (void 0 !== _child.splitText
                ? isHydrating ? _child.nodeValue.trim() : !0
                : isHydrating)
            )
              children[childrenLen++] = _child
          }
        if (0 !== vlen)
          for (var i = 0; i < vlen; i++) {
            vchild = vchildren[i]
            child = null
            var key = vchild.key
            if (null != key) {
              if (keyedLen && void 0 !== keyed[key]) {
                child = keyed[key]
                keyed[key] = void 0
                keyedLen--
              }
            } else if (!child && min < childrenLen)
              for (j = min; j < childrenLen; j++) {
                if (
                  void 0 !== children[j] &&
                  isSameNodeType((c = children[j]), vchild, isHydrating)
                ) {
                  child = c
                  children[j] = void 0
                  if (j === childrenLen - 1) childrenLen--
                  if (j === min) min++
                  break
                }
              }
            child = idiff(child, vchild, context, mountAll)
            f = originalChildren[i]
            if (child && child !== dom && child !== f)
              if (null == f) dom.appendChild(child)
              else if (child === f.nextSibling) removeNode(f)
              else dom.insertBefore(child, f)
          }
        if (keyedLen)
          for (var i in keyed) {
            if (void 0 !== keyed[i]) recollectNodeTree(keyed[i], !1)
          }
        while (min <= childrenLen) {
          if (void 0 !== (child = children[childrenLen--]))
            recollectNodeTree(child, !1)
        }
      }
      function recollectNodeTree(node, unmountOnly) {
        var component = node._component
        if (component) unmountComponent(component)
        else {
          if (null != node.__preactattr_ && node.__preactattr_.ref)
            node.__preactattr_.ref(null)
          if (!1 === unmountOnly || null == node.__preactattr_) removeNode(node)
          removeChildren(node)
        }
      }
      function removeChildren(node) {
        node = node.lastChild
        while (node) {
          var next = node.previousSibling
          recollectNodeTree(node, !0)
          node = next
        }
      }
      function diffAttributes(dom, attrs, old) {
        var name
        for (name in old) {
          if ((!attrs || null == attrs[name]) && null != old[name])
            setAccessor(dom, name, old[name], (old[name] = void 0), isSvgMode)
        }
        for (name in attrs) {
          if (
            !(
              'children' === name ||
              'innerHTML' === name ||
              (name in old &&
                attrs[name] ===
                  ('value' === name || 'checked' === name
                    ? dom[name]
                    : old[name]))
            )
          )
            setAccessor(
              dom,
              name,
              old[name],
              (old[name] = attrs[name]),
              isSvgMode
            )
        }
      }
      function collectComponent(component) {
        var name = component.constructor.name
        ;(components[name] || (components[name] = [])).push(component)
      }
      function createComponent(Ctor, props, context) {
        var inst,
          list = components[Ctor.name]
        if (Ctor.prototype && Ctor.prototype.render) {
          inst = new Ctor(props, context)
          Component.call(inst, props, context)
        } else {
          inst = new Component(props, context)
          inst.constructor = Ctor
          inst.render = doRender
        }
        if (list)
          for (var i = list.length; i--; ) {
            if (list[i].constructor === Ctor) {
              inst.__b = list[i].__b
              list.splice(i, 1)
              break
            }
          }
        return inst
      }
      function doRender(props, state, context) {
        return this.constructor(props, context)
      }
      function setComponentProps(component, props, opts, context, mountAll) {
        if (!component.__x) {
          component.__x = !0
          if ((component.__r = props.ref)) delete props.ref
          if ((component.__k = props.key)) delete props.key
          if (!component.base || mountAll) {
            if (component.componentWillMount) component.componentWillMount()
          } else if (component.componentWillReceiveProps)
            component.componentWillReceiveProps(props, context)
          if (context && context !== component.context) {
            if (!component.__c) component.__c = component.context
            component.context = context
          }
          if (!component.__p) component.__p = component.props
          component.props = props
          component.__x = !1
          if (0 !== opts)
            if (
              1 === opts ||
              !1 !== options.syncComponentUpdates ||
              !component.base
            )
              renderComponent(component, 1, mountAll)
            else enqueueRender(component)
          if (component.__r) component.__r(component)
        }
      }
      function renderComponent(component, opts, mountAll, isChild) {
        if (!component.__x) {
          var rendered,
            inst,
            cbase,
            props = component.props,
            state = component.state,
            context = component.context,
            previousProps = component.__p || props,
            previousState = component.__s || state,
            previousContext = component.__c || context,
            isUpdate = component.base,
            nextBase = component.__b,
            initialBase = isUpdate || nextBase,
            initialChildComponent = component._component,
            skip = !1
          if (isUpdate) {
            component.props = previousProps
            component.state = previousState
            component.context = previousContext
            if (
              2 !== opts &&
              component.shouldComponentUpdate &&
              !1 === component.shouldComponentUpdate(props, state, context)
            )
              skip = !0
            else if (component.componentWillUpdate)
              component.componentWillUpdate(props, state, context)
            component.props = props
            component.state = state
            component.context = context
          }
          component.__p = component.__s = component.__c = component.__b = null
          component.__d = !1
          if (!skip) {
            rendered = component.render(props, state, context)
            if (component.getChildContext)
              context = extend(extend({}, context), component.getChildContext())
            var toUnmount,
              base,
              childComponent = rendered && rendered.nodeName
            if ('function' == typeof childComponent) {
              var childProps = getNodeProps(rendered)
              inst = initialChildComponent
              if (
                inst &&
                inst.constructor === childComponent &&
                childProps.key == inst.__k
              )
                setComponentProps(inst, childProps, 1, context, !1)
              else {
                toUnmount = inst
                component._component = inst = createComponent(
                  childComponent,
                  childProps,
                  context
                )
                inst.__b = inst.__b || nextBase
                inst.__u = component
                setComponentProps(inst, childProps, 0, context, !1)
                renderComponent(inst, 1, mountAll, !0)
              }
              base = inst.base
            } else {
              cbase = initialBase
              toUnmount = initialChildComponent
              if (toUnmount) cbase = component._component = null
              if (initialBase || 1 === opts) {
                if (cbase) cbase._component = null
                base = diff(
                  cbase,
                  rendered,
                  context,
                  mountAll || !isUpdate,
                  initialBase && initialBase.parentNode,
                  !0
                )
              }
            }
            if (
              initialBase &&
              base !== initialBase &&
              inst !== initialChildComponent
            ) {
              var baseParent = initialBase.parentNode
              if (baseParent && base !== baseParent) {
                baseParent.replaceChild(base, initialBase)
                if (!toUnmount) {
                  initialBase._component = null
                  recollectNodeTree(initialBase, !1)
                }
              }
            }
            if (toUnmount) unmountComponent(toUnmount)
            component.base = base
            if (base && !isChild) {
              var componentRef = component,
                t = component
              while ((t = t.__u)) {
                ;(componentRef = t).base = base
              }
              base._component = componentRef
              base._componentConstructor = componentRef.constructor
            }
          }
          if (!isUpdate || mountAll) mounts.unshift(component)
          else if (!skip) {
            if (component.componentDidUpdate)
              component.componentDidUpdate(
                previousProps,
                previousState,
                previousContext
              )
            if (options.afterUpdate) options.afterUpdate(component)
          }
          if (null != component.__h)
            while (component.__h.length) {
              component.__h.pop().call(component)
            }
          if (!diffLevel && !isChild) flushMounts()
        }
      }
      function buildComponentFromVNode(dom, vnode, context, mountAll) {
        var c = dom && dom._component,
          originalComponent = c,
          oldDom = dom,
          isDirectOwner = c && dom._componentConstructor === vnode.nodeName,
          isOwner = isDirectOwner,
          props = getNodeProps(vnode)
        while (c && !isOwner && (c = c.__u)) {
          isOwner = c.constructor === vnode.nodeName
        }
        if (c && isOwner && (!mountAll || c._component)) {
          setComponentProps(c, props, 3, context, mountAll)
          dom = c.base
        } else {
          if (originalComponent && !isDirectOwner) {
            unmountComponent(originalComponent)
            dom = oldDom = null
          }
          c = createComponent(vnode.nodeName, props, context)
          if (dom && !c.__b) {
            c.__b = dom
            oldDom = null
          }
          setComponentProps(c, props, 1, context, mountAll)
          dom = c.base
          if (oldDom && dom !== oldDom) {
            oldDom._component = null
            recollectNodeTree(oldDom, !1)
          }
        }
        return dom
      }
      function unmountComponent(component) {
        if (options.beforeUnmount) options.beforeUnmount(component)
        var base = component.base
        component.__x = !0
        if (component.componentWillUnmount) component.componentWillUnmount()
        component.base = null
        var inner = component._component
        if (inner) unmountComponent(inner)
        else if (base) {
          if (base.__preactattr_ && base.__preactattr_.ref)
            base.__preactattr_.ref(null)
          component.__b = base
          removeNode(base)
          collectComponent(component)
          removeChildren(base)
        }
        if (component.__r) component.__r(null)
      }
      function Component(props, context) {
        this.__d = !0
        this.context = context
        this.props = props
        this.state = this.state || {}
      }
      function render(vnode, parent, merge) {
        return diff(merge, vnode, {}, !1, parent, !1)
      }
      var options = {}
      var stack = []
      var EMPTY_CHILDREN = []
      var defer =
        'function' == typeof Promise
          ? Promise.resolve().then.bind(Promise.resolve())
          : setTimeout
      var IS_NON_DIMENSIONAL = /acit|ex(?:s|g|n|p|$)|rph|ows|mnc|ntw|ine[ch]|zoo|^ord/i
      var items = []
      var mounts = []
      var diffLevel = 0
      var isSvgMode = !1
      var hydrating = !1
      var components = {}
      extend(Component.prototype, {
        setState: function(state, callback) {
          var s = this.state
          if (!this.__s) this.__s = extend({}, s)
          extend(s, 'function' == typeof state ? state(s, this.props) : state)
          if (callback) (this.__h = this.__h || []).push(callback)
          enqueueRender(this)
        },
        forceUpdate: function(callback) {
          if (callback) (this.__h = this.__h || []).push(callback)
          renderComponent(this, 2)
        },
        render: function() {}
      })
      var preact = {
        h: h,
        createElement: h,
        cloneElement: cloneElement,
        Component: Component,
        render: render,
        rerender: rerender,
        options: options
      }
      if ('undefined' != typeof module) module.exports = preact
      else self.preact = preact
    })()
  })()
  pkg[
    'github.com/matthewmueller/golly/testdata/56-hn-preact/fetch/unfetch.js'
  ] = (function() {
    return function fetch(url, options) {
      options = options || {}
      return new Promise(function(resolve, reject) {
        var request = new XMLHttpRequest()

        request.open(options.method || 'get', url)

        for (var i in options.headers) {
          request.setRequestHeader(i, options.headers[i])
        }

        request.withCredentials = options.credentials == 'include'

        request.onload = function() {
          resolve(response())
        }

        request.onerror = reject

        request.send(options.body)

        function response() {
          var keys = [],
            all = [],
            headers = {},
            header

          request
            .getAllResponseHeaders()
            .replace(/^(.*?):\s*([\s\S]*?)$/gm, function(m, key, value) {
              keys.push((key = key.toLowerCase()))
              all.push([key, value])
              header = headers[key]
              headers[key] = header ? header + ',' + value : value
            })

          return {
            ok: ((request.status / 200) | 0) == 1, // 200-299
            status: request.status,
            statusText: request.statusText,
            url: request.responseURL,
            clone: response,
            responseText: request.responseText,
            text: function() {
              return Promise.resolve(request.responseText)
            },
            json: function() {
              return Promise.resolve(request.responseText).then(JSON.parse)
            },
            blob: function() {
              return Promise.resolve(new Blob([request.response]))
            },
            headers: {
              keys: function() {
                return keys
              },
              entries: function() {
                return all
              },
              get: function(n) {
                return headers[n.toLowerCase()]
              },
              has: function(n) {
                return n.toLowerCase() in headers
              }
            }
          }
        }
      })
    }
  })()
  pkg[
    'github.com/matthewmueller/golly/testdata/56-hn-preact/item'
  ] = (function() {
    function Item(o) {
      o = o || {}
      this.ID = o.ID || ''
      this.Score = o.Score || 0
      this.Descendants = o.Descendants || 0
      this.Title = o.Title || ''
      this.URL = o.URL || ''
    }
    return {
      Item: Item
    }
  })()
  pkg[
    'github.com/matthewmueller/golly/testdata/56-hn-preact/preact'
  ] = (function() {
    var preact =
      pkg[
        'github.com/matthewmueller/golly/testdata/56-hn-preact/preact/preact.js'
      ]
    return {}
  })()
  pkg[
    'github.com/matthewmueller/golly/testdata/56-hn-preact/app'
  ] = (function() {
    var item = pkg['github.com/matthewmueller/golly/testdata/56-hn-preact/item']
    var api = 'https://hacker-news.firebaseio.com'
    app.prototype.loadFirst = function _callee5(nth) {
      var a, $res, res, err, itemIDs, items, _, id, it

      return regeneratorRuntime.async(
        function _callee5$(_context5) {
          while (1) {
            switch ((_context5.prev = _context5.next)) {
              case 0:
                a = this
                _context5.next = 3
                return regeneratorRuntime.awrap(
                  (function _callee(unfetch, url) {
                    var res
                    return regeneratorRuntime.async(
                      function _callee$(_context) {
                        while (1) {
                          switch ((_context.prev = _context.next)) {
                            case 0:
                              _context.prev = 0
                              _context.next = 3
                              return regeneratorRuntime.awrap(unfetch(url))

                            case 3:
                              res = _context.sent
                              return _context.abrupt('return', [res, null])

                            case 7:
                              _context.prev = 7
                              _context.t0 = _context['catch'](0)
                              return _context.abrupt('return', [
                                null,
                                _context.t0
                              ])

                            case 10:
                            case 'end':
                              return _context.stop()
                          }
                        }
                      },
                      null,
                      this,
                      [[0, 7]]
                    )
                  })(
                    pkg[
                      'github.com/matthewmueller/golly/testdata/56-hn-preact/fetch/unfetch.js'
                    ],
                    api + '/v0/topstories.json'
                  )
                )

              case 3:
                $res = _context5.sent
                res = $res[0]
                err = $res[1]

                if (!(err != null)) {
                  _context5.next = 8
                  break
                }

                return _context5.abrupt('return', [null, err])

              case 8:
                itemIDs = []
                _context5.next = 11
                return regeneratorRuntime.awrap(
                  (function _callee2($obj) {
                    var $o, $k
                    return regeneratorRuntime.async(
                      function _callee2$(_context2) {
                        while (1) {
                          switch ((_context2.prev = _context2.next)) {
                            case 0:
                              _context2.prev = 0
                              _context2.next = 3
                              return regeneratorRuntime.awrap(res.json())

                            case 3:
                              $o = _context2.sent

                              for ($k in $o) {
                                $obj[$k] = $o[$k]
                              }
                              return _context2.abrupt('return', null)

                            case 8:
                              _context2.prev = 8
                              _context2.t0 = _context2['catch'](0)
                              return _context2.abrupt('return', _context2.t0)

                            case 11:
                            case 'end':
                              return _context2.stop()
                          }
                        }
                      },
                      null,
                      this,
                      [[0, 8]]
                    )
                  })(itemIDs)
                )

              case 11:
                err = _context5.sent

                if (!(err != null)) {
                  _context5.next = 14
                  break
                }

                return _context5.abrupt('return', [null, err])

              case 14:
                items = []
                _ = 0

              case 16:
                if (!(_ < itemIDs.slice(0, nth).length)) {
                  _context5.next = 36
                  break
                }

                id = itemIDs.slice(0, nth)[_]
                _context5.next = 20
                return regeneratorRuntime.awrap(
                  (function _callee3(unfetch, url) {
                    var res
                    return regeneratorRuntime.async(
                      function _callee3$(_context3) {
                        while (1) {
                          switch ((_context3.prev = _context3.next)) {
                            case 0:
                              _context3.prev = 0
                              _context3.next = 3
                              return regeneratorRuntime.awrap(unfetch(url))

                            case 3:
                              res = _context3.sent
                              return _context3.abrupt('return', [res, null])

                            case 7:
                              _context3.prev = 7
                              _context3.t0 = _context3['catch'](0)
                              return _context3.abrupt('return', [
                                null,
                                _context3.t0
                              ])

                            case 10:
                            case 'end':
                              return _context3.stop()
                          }
                        }
                      },
                      null,
                      this,
                      [[0, 7]]
                    )
                  })(
                    pkg[
                      'github.com/matthewmueller/golly/testdata/56-hn-preact/fetch/unfetch.js'
                    ],
                    api + '/v0/item/' + String(id) + '.json'
                  )
                )

              case 20:
                $res = _context5.sent
                res = $res[0]
                err = $res[1]

                if (!(err != null)) {
                  _context5.next = 25
                  break
                }

                return _context5.abrupt('return', [null, err])

              case 25:
                it = new item.Item()
                _context5.next = 28
                return regeneratorRuntime.awrap(
                  (function _callee4($obj) {
                    var $o, $k
                    return regeneratorRuntime.async(
                      function _callee4$(_context4) {
                        while (1) {
                          switch ((_context4.prev = _context4.next)) {
                            case 0:
                              _context4.prev = 0
                              _context4.next = 3
                              return regeneratorRuntime.awrap(res.json())

                            case 3:
                              $o = _context4.sent

                              for ($k in $o) {
                                console.log($k, $o[$k])
                                $obj[$k] = $o[$k]
                              }
                              return _context4.abrupt('return', null)

                            case 8:
                              _context4.prev = 8
                              _context4.t0 = _context4['catch'](0)
                              return _context4.abrupt('return', _context4.t0)

                            case 11:
                            case 'end':
                              return _context4.stop()
                          }
                        }
                      },
                      null,
                      this,
                      [[0, 8]]
                    )
                  })(it)
                )

              case 28:
                err = _context5.sent

                console.log('TITLE', it.title, it.Title)

                if (!(err != null)) {
                  _context5.next = 32
                  break
                }

                return _context5.abrupt('return', [null, err])

              case 32:
                items = items.concat(it)

              case 33:
                _++
                _context5.next = 16
                break

              case 36:
                return _context5.abrupt('return', [items, null])

              case 37:
              case 'end':
                return _context5.stop()
            }
          }
        },
        null,
        this
      )
    }
    function state(o) {
      o = o || {}
      this.items = o.items || []
    }
    app.prototype.componentDidMount = function _callee6() {
      var a, $items, items, err
      return regeneratorRuntime.async(
        function _callee6$(_context6) {
          while (1) {
            switch ((_context6.prev = _context6.next)) {
              case 0:
                a = this

                console.log('loading first...')
                _context6.next = 4
                return regeneratorRuntime.awrap(a.loadFirst.bind(a)(2))

              case 4:
                $items = _context6.sent
                items = $items[0]
                err = $items[1]

                if (!(err != null)) {
                  _context6.next = 9
                  break
                }

                throw err

              case 9:
                console.log('setting state...')
                a.setState(
                  new state({
                    items: items
                  })
                )

              case 11:
              case 'end':
                return _context6.stop()
            }
          }
        },
        null,
        this
      )
    }
    app.prototype.render = function() {
      var a = this
      var items = []
      var children = ['Hi'].concat(items)
      return preact.h('div', null, [children])
    }
    function props(o) {
      o = o || {}
    }
    function app(o) {
      o = o || {}
      this.Component = o.Component || new preact.Component()
      for (var $k in this.Component || preact.Component.prototype) {
        this[$k] =
          this[$k] || (this.Component || preact.Component.prototype)[$k]
      }
      this.props = o.props || new props()
      this.state = o.state || new state()
    }
    function New() {
      var a
      return regeneratorRuntime.async(
        function New$(_context7) {
          while (1) {
            switch ((_context7.prev = _context7.next)) {
              case 0:
                a = preact.h(app, {})
                return _context7.abrupt('return', a)

              case 2:
              case 'end':
                return _context7.stop()
            }
          }
        },
        null,
        this
      )
    }
    return {
      New: New
    }
  })()
  pkg['github.com/matthewmueller/golly/testdata/56-hn-preact'] = (function() {
    var app = pkg['github.com/matthewmueller/golly/testdata/56-hn-preact/app']
    function main() {
      return regeneratorRuntime.async(
        function main$(_context8) {
          while (1) {
            switch ((_context8.prev = _context8.next)) {
              case 0:
                _context8.t0 = preact
                _context8.next = 3
                return regeneratorRuntime.awrap(app.New())

              case 3:
                _context8.t1 = _context8.sent
                _context8.t2 = document.body

                _context8.t0.render.call(
                  _context8.t0,
                  _context8.t1,
                  _context8.t2
                )

                _context8.next = 8
                return regeneratorRuntime.awrap(
                  new Promise(function(resolve, reject) {
                    setTimeout(resolve, 3 * 1000)
                  })
                )

              case 8:
                console.log('done')

              case 9:
              case 'end':
                return _context8.stop()
            }
          }
        },
        null,
        this
      )
    }
    return {
      main: main
    }
  })()
  return pkg['github.com/matthewmueller/golly/testdata/56-hn-preact'].main()
})()
