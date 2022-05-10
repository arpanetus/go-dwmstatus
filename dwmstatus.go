package main

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
import "C"

import (
	"fmt"
)


func main() {
	var dpy = C.XOpenDisplay(nil)

	fmt.Println(dpy)
}
