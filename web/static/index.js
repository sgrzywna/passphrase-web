function ready (fn) {
  if (document.attachEvent ? document.readyState === 'complete' : document.readyState !== 'loading') {
    fn()
  } else {
    document.addEventListener('DOMContentLoaded', fn)
  }
}

function onReady () {
  var elem = document.getElementById('submit-btn')
  elem.addEventListener('click', onSubmit)
  getPasswords()
}

function getPasswords () {
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
    if (req.status === 200) {
      var data = JSON.parse(req.responseText)
      processResponse(data)
    } else if (req.status === 429) {
      showTooMany()
    } else {
      responseError()
    }
  }

  req.onerror = function () {
    responseError()
  }

  req.send()
}

function processResponse (data) {
  hideError()
  hideTooMany()
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

function responseError () {
  showError()
}

function hideError () {
  var elem = document.getElementById('error')
  addClass(elem, 'siimple--display-none')
}

function showError () {
  var elem = document.getElementById('error')
  removeClass(elem, 'siimple--display-none')
}

function hideTooMany () {
  var elem = document.getElementById('toomany')
  addClass(elem, 'siimple--display-none')
}

function showTooMany () {
  var elem = document.getElementById('toomany')
  removeClass(elem, 'siimple--display-none')
}

function addClass (elem, cls) {
  var arr = elem.className.split(' ')
  if (arr.indexOf(cls) === -1) {
    elem.className += ' ' + cls
  }
}

function removeClass (elem, cls) {
  var arr = elem.className.split(' ')
  var ndx = arr.indexOf(cls)
  if (ndx !== -1) {
    arr.splice(ndx, 1)
    elem.className = arr.join(' ')
  }
}

function onSubmit (event) {
  getPasswords()
}

ready(onReady)
