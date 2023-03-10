package main

import (
  "fmt"
  "os"
)

func Fatal(msg string) {
  fmt.Fprintln(os.Stderr, msg)
  os.Exit(255)
}
