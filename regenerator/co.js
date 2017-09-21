var co = require('co')
console.log(co)

co(
  function*() {
    yield sleep(5000)
    return 'hi'
  },
  function(err, msg) {
    console.log('all done!', msg)
  }
)

function sleep(ms) {
  return function(fn) {
    setTimeout(fn, ms)
  }
}
