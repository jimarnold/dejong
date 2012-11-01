package main

type Plot struct {
  pixels [][]int
  x float64
  y float64
}

func NewPlot(width, height int) *Plot {
  pixels := make([][]int, width)
  for x := range(pixels) {
    pixels[x] = make([]int, height)
  }
  return &Plot{pixels:pixels}
}

func (plot *Plot) Clear() {
  for x := range(plot.pixels) {
    for y := range(plot.pixels[x]) {
      plot.pixels[x][y] = 0
    }
  }
}
