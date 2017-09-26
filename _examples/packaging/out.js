;(function() {
  var dep2 = (function() {
    // AnotherGet
    function AnotherGet() {
      return 'another world'
    }

    return {
      AnotherGet: AnotherGet
    }
  })()

  var dep = (function() {
    // Get
    function Get() {
      return 'world'
    }

    return {
      Get: Get
    }
  })()

  function main() {
    console.log('hi ' + dep.Get() + ' ' + Side() + ' ' + dep2.AnotherGet())
  }

  function Side() {
    return 'Side'
  }

  main()
})()
