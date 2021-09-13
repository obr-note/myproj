package myproj

import (
  "flag"
  "fmt"
)

var version = "1.0.0"

func Run() {
  var showVersion bool
  flag.BoolVar(&showVersion, "v", false, "show version")
  flag.BoolVar(&showVersion, "version", false, "show version")
  flag.Parse()
  if showVersion {
    fmt.Println("version:", version)
    return
  }
}
