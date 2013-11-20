## goflam3

I've got this running on heroku here, so you can play around with the web interface.

### Usage:

`goflam3 serve` serve the webapp. That's what you see on heroku
`goflam3 render` render an image. Hit `goflame3 render -h` for more information.

## Background

Flame fractals in golang

Basing my work off of the following texts:

- http://web.mit.edu/fustflum/documents/flame.pdf
- http://flam3.com/flame_draves.pdf
- http://illusions.hu/Softwares/IFSIllusions/Enhancement%20of%20the%20Fractal%20Flame%20Algorithm.pdf
- http://en.wikipedia.org/wiki/Fractal_flame

## Roadmap

- Add colorization, color palettes
- Add more functions
- Indicate progress
- Perhaps jump to importing .flam3 files? Then it would be interesting to benchmark against flam3-render
- Use a pool of goroutines to make awesome faster
- enableable log equalization
- enableable gamma equalization
