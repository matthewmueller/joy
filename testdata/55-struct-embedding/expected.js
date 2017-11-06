function inherit(ctor, superCtor) {
  ctor.super_ = superCtor
  ctor.prototype = Object.create(superCtor.prototype, {
    constructor: {
      value: ctor,
      enumerable: false,
      writable: true,
      configurable: true
    }
  })
}

function thing() {}

thing.prototype.kind = function() {
  return 1
}

function animal(o) {
  o = o || {}
  this.hello = o.hello || ''
  for (var k in thing.prototype) {
    this[k] = thing.prototype[k]
  }
}

animal.prototype.name = function() {
  return 'animal'
}

function dog(o) {
  o = o || {}
  this.animal = o.animal || new animal()

  for (var $k in this.animal || animal.prototype) {
    this[$k] = (this.animal || animal.prototype)[$k]
  }
}

dog.prototype.legs = function() {
  return 4
}

var d = new dog({
  animal: new animal({
    hello: 'hi'
  })
})
console.log(d.name(), d.legs(), d.kind(), d.hello)
