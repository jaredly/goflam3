var x = require('xon')
  , request = require('superagent')

angular.module('MyApp', [])
  .factory('highDef', function () {
    return function (funcs, cb) {
      if (!funcs) funcs = []
      var arg = funcs.map(function (f) {return f.Num + ''}).join(':')
      if (window.localStorage['flame-hd-' + arg]) {
        return cb(window.localStorage['flame-hd-' + arg], true)
      }
      request.get('/high-def?funcs=' + arg)
        .end(function (req) {
          try {
            window.localStorage['flame-hd-' + arg] = req.text
          } catch (e) {}
          cb(req.text)
        })
    }
  })
  .factory('getData', function () {
    return function (funcs, cb) {
      if (!funcs) funcs = []
      var arg = funcs.map(function (f) {return f.Num + ''}).join(':')
      if (funcs.length && window.localStorage['flame-' + arg]) {
        try {
          return cb(JSON.parse(window.localStorage['flame-' + arg]), true)
        } catch (e) {}
      }
      request.get('/render?funcs=' + arg)
        .end(function (req) {
          try {
            window.localStorage['flame-' + arg] = JSON.stringify(req.body)
          } catch (e) {}
          cb(req.body)
        })
    }
  })
  .controller('MainController', ['$scope', 'getData', 'highDef', function ($scope, getData, highDef) {
    function gotData(data, cached) {
      for (var name in data) {
        if (!name.match(/^[a-zA-Z0-9_-]+$/)) continue;
        $scope[name] = data[name]
      }
      if (!cached) $scope.$digest()
    }
    getData([], gotData)

    $scope.useChild = function (fractal) {
      getData(fractal.Formulas, gotData)
    }

    $scope.showHD = function () {
      $scope.showingHD = true
      highDef($scope.MainFormulas, function (data, cached) {
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
