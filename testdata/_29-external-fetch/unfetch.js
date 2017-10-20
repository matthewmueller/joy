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
