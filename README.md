## goflam3

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

## Examples
So far they're pretty subdued because I haven't implemented good equalization yet. But they're super cool!

![one](https://github.com/jaredly/goflam3/blob/master/new.png?raw=true)
![two](https://github.com/jaredly/goflam3/blob/master/new1235.png?raw=true)
