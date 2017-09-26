;(function() {
  var pkg = {}

  pkg[
    'github.com/matthewmueller/golly/_examples/packaging/another/dep'
  ] = (function() {
    console.log(pkg)
    var dep2 =
      pkg['github.com/matthewmueller/golly/_examples/packaging/another/more']
    function AnotherGet() {
      return 'another world ' + dep2.More()
    }
    return {
      AnotherGet: AnotherGet
    }
  })()

  pkg['github.com/matthewmueller/golly/_examples/packaging/dep'] = (function() {
    function Get() {
      return 'world'
    }
    return {
      Get: Get
    }
  })()

  pkg[
    'github.com/matthewmueller/golly/_examples/packaging/another/more'
  ] = (function() {
    function More() {
      return 'more'
    }
    return {
      More: More
    }
  })()

  pkg['github.com/matthewmueller/golly/_examples/packaging'] = (function() {
    var dep2 =
        pkg['github.com/matthewmueller/golly/_examples/packaging/another/dep'],
      dep = pkg['github.com/matthewmueller/golly/_examples/packaging/dep']
    function main() {
      console.log('hi ' + dep.Get() + ' ' + Side() + ' ' + dep2.AnotherGet())
    }
    function Test() {
      return 'another test'
    }
    function hi() {
      return 'hi'
    }
    var unexported = 'hi'
    function Side() {
      return 'side'
    }
    return {
      main: main,
      Test: Test,
      Side: Side
    }
  })()

  console.log(pkg)
  pkg['github.com/matthewmueller/golly/_examples/packaging'].main()
})()
