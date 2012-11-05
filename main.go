package main

import (
  "fmt"
  "time"
  "flag"
  "math/rand"
  "github.com/go-gl/gl"
  "github.com/go-gl/glfw"
)

const brightnessStep int = 3

func init() {
  rand.Seed( time.Now().UTC().UnixNano())
}

func main() {
  var iterations int
  var displayTime float64
  var width int
  var height int

  flag.IntVar(&iterations, "i", 50000, "iterations per frame")
  flag.Float64Var(&displayTime, "t", 10, "display time, in seconds, for each plot")
  flag.IntVar(&width, "w", 800, "width")
  flag.IntVar(&height, "h", 600, "height")
  flag.Parse()

  initGlfw(width, height)
  defer terminateGlfw()

  palette := NewPalette()
  plot := NewPlot(width, height)
  attractor := NewAttractor(width, height)
  startTime := time.Now()

  for glfw.WindowParam(glfw.Opened) == 1 && glfw.Key(glfw.KeyEsc) != glfw.KeyPress {
    if time.Since(startTime).Seconds() > displayTime {
      reseed(plot, attractor)
      startTime = time.Now()
    }
    attractor.iterate(plot, iterations)
    render(plot, palette)
    glfw.SwapBuffers()
  }
}

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

func reseed(plot *Plot, attractor *Attractor) {
  attractor.seed()
  plot.Clear()
}

func initGlfw(width, height int) {
  var err error
  if err = glfw.Init(); err != nil {
    fmt.Printf("%v\n", err)
    return
  }

  if err = glfw.OpenWindow(width, height, 0, 0, 0, 0, 0, 0, glfw.Windowed); err != nil {
    fmt.Printf("%logv\n", err)
    return
  }

  glfw.SetWindowSizeCallback(onResize)
  glfw.SetSwapInterval(1)
}

func terminateGlfw() {
  glfw.Terminate()
}

func onResize(w, h int) {
  gl.MatrixMode(gl.PROJECTION)
  gl.LoadIdentity()
  gl.Viewport(0, 0, w, h)
  gl.Ortho(0, float64(w), float64(h), 0, -1, 1)
  gl.ClearColor(0, 0, 0, 0)
  gl.Clear(gl.COLOR_BUFFER_BIT)
  gl.MatrixMode(gl.MODELVIEW)
  gl.LoadIdentity()
}

func render(plot *Plot, palette *Palette) {
  gl.ClearColor(0.0, 0.0, 0.05, 0)
  gl.Clear(gl.COLOR_BUFFER_BIT)
  gl.Enable(gl.POINT_SMOOTH)
  gl.Enable( gl.BLEND )
  gl.BlendFunc( gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA )
  gl.PointSize(1)
  gl.Begin(gl.POINTS)

  for x := range plot.pixels {
    for y := range plot.pixels[x] {
      level := plot.pixels[x][y]
      if level <= 0 {
        continue
      }
      if level > 255 {
        level = 255
      }
      rgb := palette.colors[level]
      gl.Color3ub(rgb.r, rgb.g, rgb.b)
      gl.Vertex3i(x, y, 0)
    }
  }
  gl.End()
}
