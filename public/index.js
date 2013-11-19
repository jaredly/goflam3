var x = require('xon')
  , request = require('superagent')

function makeKey(funcs) {
  return funcs.map(function (f) {return f.Num + ''}).join(':')
}

angular.module('MyApp', [])
  .factory('highDef', function () {
    return function (key, cb) {
      if (window.localStorage['flame-hd-' + key]) {
        return cb(window.localStorage['flame-hd-' + key], true)
      }
      request.get('/high-def?funcs=' + key)
        .end(function (req) {
          try {
            window.localStorage['flame-hd-' + key] = req.text
          } catch (e) {}
          cb(req.text)
        })
    }
  })
  .factory('getData', function () {
    return function (key, cb) {
      if (window.localStorage['flame-' + key]) {
        try {
          return cb(JSON.parse(window.localStorage['flame-' + key]), true)
        } catch (e) {}
      }
      request.get('/render?funcs=' + key)
        .end(function (req) {
          try {
            window.localStorage['flame-' + key] = JSON.stringify(req.body)
          } catch (e) {}
          cb(req.body)
        })
    }
  })
  .controller('MainController', ['$scope', 'getData', 'highDef', function ($scope, getData, highDef) {
    
    window.addEventListener("hashchange", function () {
      getData(window.location.hash.slice(1), function (data) {
        gotData(data)
      })
    })

    function gotData(data, cached) {
      for (var name in data) {
        if (!name.match(/^[a-zA-Z0-9_-]+$/)) continue;
        $scope[name] = data[name]
      }
      if (!cached) $scope.$digest()
    }
    // getData([], gotData)
    if (!window.location.hash || window.location.hash == '#') {
      window.location.hash = '5:7'
    } else {
      getData(window.location.hash.slice(1), gotData)
    }
    
    $scope.logEqualize = false;
    $scope.$watch('logEqualize', function (value, old) {
      if ('undefined' === typeof old) return
      var hash = window.location.hash.slice(1)
        , parts = hash.split('&')
        , current = parts.length === 2
      if (value === current) return
      if (value) {
        hash += '&log'
      } else {
        hash = parts[0]
      }
      window.location.hash = hash
    })

    $scope.useChild = function (fractal) {
      window.location.hash = makeKey(fractal.Formulas)
    }

    $scope.showHD = function () {
      if (!$scope.MainFormulas.length) return
      $scope.showingHD = true
      highDef(makeKey($scope.MainFormulas), function (data, cached) {
        $scope.HighDef = data
        if (!cached) $scope.$digest()
      })
    }

    $scope.hideHD = function () {
      $scope.showingHD = false
      $scope.HighDef = false
    }
  }])

module.exports = function (document) {
  var el = document.getElementById('main')
  angular.bootstrap(el, ['MyApp'])
}
