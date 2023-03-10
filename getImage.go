package main

import (
  "os"
  "fmt"
  "github.com/cavaliercoder/grab"
)


func getImage(url string){
  resp, err := grab.Get(".", url)
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
  }
  fmt.Println("Download saved to:", resp.Filename)
}
