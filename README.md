An implementation of Peter de Jong's strange attractor algorithm, written as an exercise for learning the Go language.

It depends on the Go bindings for OpenGL and glfw, which may require you to do some wrangling of the native libraries to install correctly, depending on your environment.

This is my first go program. It was initially ported from a CoffeeScript implementation from Jeremy Ashkenas [@jashkenas](http://twitter.com/jashkenas) and heavily refactored (not that there was anything wrong with the original, but refactoring helps me to understand code).

More on the maths behind it:

http://www.cc.gatech.edu/~phlosoft/attractors/
http://www.complexification.net/gallery/machines/peterdejong/

Prerequisites for building on OSX:

brew install glew
brew install glfw
go get github.com/go-gl/gl
go get github.com/go-gl/glfw

To run:

go build
./dejong

Command line flags:

  -w  window width
  -h  window height
  -i  iterations per frame
  -f  frames per 'attraction'

