function h(nodeName, attributes, children) {
  console.log('nodeName', nodeName.name || nodeName)
  console.log('attributes', attributes)
  console.log('children', children)
  return {
    nodeName: nodeName,
    attributes: attributes,
    children: children
  }
}

function render(vdom, dom) {
  console.log('vdom=', vdom)
  console.log('dom=', dom)
}

return {
  h: h,
  render: render
}
