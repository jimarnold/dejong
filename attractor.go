package main

import (
  "math"
  "math/rand"
)

type Attractor struct {
  a float64
  b float64
  c float64
  d float64
  width float64
  height float64
  sensitivity float64
}

func NewAttractor(width, height int, sensitivity float64) *Attractor {
  a := &Attractor{width: float64(width), height: float64(height), sensitivity: sensitivity}
  a.seed()
  return a
}

func (a *Attractor) seed() {
  a.a = rand.Float64() * a.sensitivity
  a.b = rand.Float64() * a.sensitivity
  a.c = -rand.Float64() * a.sensitivity
  a.d = -rand.Float64() * a.sensitivity
}

func (a *Attractor) iterate(p *Plot, iterations int) {
  deJong := func(x, y float64) (float64, float64) {
    return math.Sin(a.a * y) - math.Cos(a.b * x), math.Sin(a.c * x) - math.Cos(a.d * y)
  }

  transform := func(x, y float64) (float64, float64) {
    //this is a linear transformation from the domain of sin/cos (-2 -> +2) into
    //the coordinate space of the plot (width/height)
    return ((x + 2) * a.width) / 4, ((y + 2) * a.height) / 4
  }

  for i := 0; i < iterations; i++ {
    x,y := transform(deJong(p.x,p.y))
    p.pixels[int(x)][int(y)] += brightnessStep
    p.x = x
    p.y = y
  }
}
