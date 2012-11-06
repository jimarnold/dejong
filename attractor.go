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
  e float64
  f float64
  width float64
  height float64
  depth float64
  sensitivity float64
}

func NewAttractor(width, height int, sensitivity float64) *Attractor {
  a := &Attractor{width: float64(width), height: float64(height), depth: float64(height), sensitivity: sensitivity}
  a.seed()
  return a
}

func (a *Attractor) seed() {
  a.a = rand.Float64() * a.sensitivity
  a.b = rand.Float64() * a.sensitivity
  a.c = -rand.Float64() * a.sensitivity
  a.d = -rand.Float64() * a.sensitivity
  a.e = rand.Float64() * a.sensitivity
  a.f = rand.Float64() * a.sensitivity
}

func (a *Attractor) iterate(p *Plot, iterations int) {
  deJong := func(v Vector) (Vector) {
    return Vector{math.Sin(a.a *v.y) - math.Cos(a.b * v.x), math.Sin(a.c * v.x) - math.Cos(a.d * v.z), math.Sin(a.e * v.z) - math.Cos(a.f * v.x)}
  }

  transform := func(v Vector) (Vector) {
    //this is a linear transformation from the domain of sin/cos (-2 -> +2) into
    //the coordinate space of the plot (width/height)
    return Vector{((v.x + 2) * a.width) / 4, ((v.y + 2) * a.height) / 4, ((v.z + 2) * a.depth) / 4}
  }

  for i := 0; i < iterations; i++ {
    v := transform(deJong(p.v))
    p.pixels[v] += brightnessStep
    p.v = v
  }
}
