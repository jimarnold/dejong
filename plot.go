package main

type Vector struct {
  x, y, z float64
}

type Plot struct {
  pixels map[Vector]int
  v Vector
}

func NewPlot() *Plot {
  pixels := make(map[Vector]int)
  return &Plot{pixels:pixels}
}

func (plot *Plot) Clear() {
  plot.pixels = make(map[Vector]int)
}
