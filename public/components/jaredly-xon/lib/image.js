
var _cache = {}

function text_size(width, height, template) {
  height = parseInt(height, 10);
  width = parseInt(width, 10);
  var bigSide = Math.max(height, width)
  var smallSide = Math.min(height, width)
  var scale = 1 / 12;
  var newHeight = Math.min(smallSide * 0.75, 0.75 * bigSide * scale);
  return {
    height: Math.round(Math.max(template.size, newHeight))
  }
}

var canvas = document.createElement('canvas');
var ctx = canvas.getContext("2d");

module.exports = function (width, height, text) {
  var args = {
    dimensions: {
      width: width,
      height: height
    },
    template: {
      background: "#aaa",
      foreground: '#fff',
      text: text
    },
    ratio: 1
  }
  var key = JSON.stringify(args)
  if (_cache[key]) return _cache[key]
  var dimensions = args.dimensions;
  var template = args.template;
  var ratio = args.ratio;

  var ts = text_size(dimensions.width, dimensions.height, template);
  var text_height = ts.height;
  var width = dimensions.width * ratio,
  height = dimensions.height * ratio;
  var font = template.font ? template.font : "Arial,Helvetica,sans-serif";
  canvas.width = width;
  canvas.height = height;
  ctx.textAlign = "center";
  ctx.textBaseline = "middle";
  ctx.fillStyle = template.background;
  ctx.fillRect(0, 0, width, height);
  ctx.fillStyle = template.foreground;
  ctx.font = "bold " + text_height + "px " + font;
  var text = template.text ? template.text : (Math.floor(dimensions.width) + "x" + Math.floor(dimensions.height));
  var text_width = ctx.measureText(text).width;
  if (text_width / width >= 0.75) {
    text_height = Math.floor(text_height * 0.75 * (width / text_width));
  }
  //Resetting font size if necessary
  ctx.font = "bold " + (text_height * ratio) + "px " + font;
  ctx.fillText(text, (width / 2), (height / 2), width);
  var dataUrl =  canvas.toDataURL("image/png");
  _cache[key] = dataUrl;
  return dataUrl
}
