package main
type RGB struct {
  r float64
  g float64
  b float64
}

type Palette []RGB

func lerp(a, b , t float64) float64 {
  return a + (b - a) * t
}

func NewPalette() Palette {
  palette := make(Palette, 256)
  c1 := RGB{0,0,0}
  c2 := RGB{0.9,0.75,1}

  for i:=0;i<=255;i++ {
    t := (float64(i)/255.0)
    r := lerp(c1.r, c2.r, t)
    g := lerp(c1.g, c2.g, t)
    b := lerp(c1.b, c2.b, t)
    palette[i] = RGB{r,g,b}
  }
  return palette
}
