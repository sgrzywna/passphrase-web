var passwords = function () {
  var ready = function (fn) {
    if (document.attachEvent ? document.readyState === 'complete' : document.readyState !== 'loading') {
      fn()
    } else {
      document.addEventListener('DOMContentLoaded', fn)
    }
  }

  var onReady = function () {
    var elem = document.getElementById('submit-btn')
    elem.addEventListener('click', onSubmit)
    getPasswords()
  }

  var getPasswords = function () {
    // dictionary
    var dictElem = document.getElementById('dictionary')
    var dict = dictElem.options[dictElem.selectedIndex].value
    // number of words
    var wordsElem = document.getElementById('words-num')
    var words = wordsElem.options[wordsElem.selectedIndex].value
    // number of passwords
    var passElem = document.getElementById('passwords-num')
    var pass = passElem.options[passElem.selectedIndex].value

    var req = new XMLHttpRequest()
    var params = 'd=' + dict + '&w=' + words + '&p=' + pass
    req.open('GET', '/passwords.json?' + params)

    req.onload = function () {
      if (req.status >= 200 && req.status < 400) {
        var data = JSON.parse(req.responseText)
        processResponse(data)
      } else {
        responseError()
      }
    }

    req.onerror = function () {
      responseError()
    }

    req.send()
  }

  var processResponse = function (data) {
    var elem = document.getElementById('passwords')
    // clean
    var divs = elem.querySelectorAll('div')
    divs.forEach(function (div) {
      elem.removeChild(div)
    })
    // render
    data.forEach(function (pass) {
      var div = document.createElement('div')
      div.setAttribute('class', 'siimple-list-item')
      div.innerHTML = pass
      elem.appendChild(div)
    })
  }

  var responseError = function () {
    console.log('response error')
  }

  var onSubmit = function (event) {
    getPasswords()
  }

  ready(onReady)
}

passwords()
