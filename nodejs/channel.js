function Deferred() {
  var p = new Promise(function(resolve, reject) {
    this.resolve = resolve
    this.reject = reject
  })

  this.then = p.then.bind(p)
  this.catch = p.catch.bind(p)
}

/**
 * Initialize channel with the given buffer `capacity`. By default
 * the channel is unbuffered. A channel is basically a FIFO queue
 * for use with async/await or co().
 */
function Channel(capacity) {
  this.capacity = capacity || 0
  this.values = []
  this.sends = []
  this.recvs = []
  this.closed = false
}

/**
 * Send value, blocking unless there is room in the buffer.
 *
 * Calls to send() on a closed buffer will error.
 */
Channel.prototype.send = function(value) {
  if (this.closed) {
    return Promise.reject(new Error('send on closed channel'))
  }

  // recv pending
  if (this.recvs.length) {
    this.recvs.shift().resolve(value)
    return Promise.resolve()
  }

  // room in buffer
  if (this.values.length < this.capacity) {
    this.values.push(value)
    return Promise.resolve()
  }

  // no recv pending, block
  var promise = new Deferred()
  this.sends.push({ value: value, promise: promise })
  return promise
}

/**
 * Receive returns a value or blocks until one is present.
 *
 * A recv() on a closed channel will return undefined.
 */
Channel.prototype.recv = function() {
  // values in buffer
  if (this.values.length) {
    return Promise.resolve(this.values.shift())
  }

  // unblock pending sends
  if (this.sends.length) {
    var send = this.sends.shift()

    if (this.closed) {
      send.promise.reject(new Error('send on closed channel'))
      return Promise.resolve()
    }

    send.promise.resolve()
    return Promise.resolve(send.value)
  }

  // closed
  if (this.closed) {
    return Promise.resolve()
  }

  // no values, block
  var promise = new Deferred()
  this.recvs.push(promise)
  return promise
}

/**
 * Close the channel. Any pending recv() calls will be unblocked.
 *
 * Subsequent close() calls will throw.
 */
Channel.prototype.close = function() {
  if (this.closed) throw new Error('channel already closed')
  this.closed = true
  var recvs = this.recvs
  this.recvs = []
  recvs.forEach(function(p) {
    return p.resolve()
  })
}
