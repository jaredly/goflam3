
var lib = require('./lib')

module.exports = lib.resolve

for (var name in lib.bound) {
  module.exports[name] = lib.bound[name]
}

module.exports.register = function (name, fn) {
  if (arguments.length === 1) {
    fn = name
    name = fn.name
  }
  module.exports[name] = lib.binder(fn)
}
