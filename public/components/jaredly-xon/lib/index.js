// people

var consts = require('./consts')
  , image = require('./image')

var helpers = {
  ObjectId: function (len) {
    var id = ''
      , chars = 'abcdef0123456789'
    len = len || 32
    for (var i=0; i<len; i++) {
      id += helpers.choice(chars)
    }
    return id
  },
  image: image,
  choice: function (items) {
    return items[helpers.randInt(items.length)]
  },
  fullName: function (/*maxlen*/) {
    // maxlen = maxlen || 23
    return helpers.choice(consts.name.first) + ' ' + helpers.choice(consts.name.last)
  },
  randInt: function (min, max) {
    if (arguments.length === 1) {
      max = min
      min = 0
    }
    return parseInt(Math.random() * (max - min)) + min
  },
  city: function () {
    return helpers.choice(consts.cities)
  },
  /*
  lipsum: function (min, max) {
    return consts.lipsum.split(' ').slice(0, helpers.randInt(min, max)).join(' ')
  },
  image: function (width, height) {
  },
  */
  some: function (min, max, fix) {
    var num, results = []
    if (arguments.length === 2) {
      fix = max
      num = min
    } else {
      num = helpers.randInt(min, max + 1)
    }
    for (var i=0; i<num; i++) {
      results.push(resolve(fix))
    }
    return results
  }
}

function binder(fn) {
  return function () {
    var args = arguments
      , self = this
      , res = function () {
          return fn.apply(self, args)
        }
    res._bound_fixture = true
    return res
  }
}

var bound = {}
for (var name in helpers) {
  bound[name] = binder(helpers[name])
}

function resolve(fix) {
  if ('function' === typeof fix && fix._bound_fixture) {
    return fix()
  }
  if ('object' !== typeof fix) return fix
  if (Array.isArray(fix)) {
    return fix.map(resolve)
  }
  var res = {}
  for (var name in fix) {
    res[name] = resolve(fix[name])
  }
  return res
}

module.exports = {
  helpers: helpers,
  resolve: resolve,
  binder: binder,
  bound: bound
}

