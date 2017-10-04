;(function() {
  var pkg = {}
  pkg[
    'github.com/matthewmueller/golly/testdata/18-goroutine-basic'
  ] = (function() {
    async function main() {
      var ch = new Channel()
      ;(async function(msg) {
        await ch.Send(msg)
      })('hi')
      console.log(await ch.Recv())
    }
    return {
      main: main
    }
  })()
  pkg['github.com/matthewmueller/golly/testdata/18-goroutine-basic'].main()
})()
