package main

type RGB struct {
  r uint8
  g uint8
  b uint8
}

type Palette struct {
  colors []RGB
}

func NewPalette() *Palette {
  palette := &Palette{}
  palette.colors = make([]RGB, 256)
  var r uint8 = 0
  var g uint8 = 0
  var b uint8 = 31
  for i:=0;i<=255;i++ {
    if r < 255 {
      r++
    }
    if g < 255 {
      g++
    }
    if b < 255 {
      b++
    }
    palette.colors[i] = RGB{r,g,b}
  }
  return palette
}
