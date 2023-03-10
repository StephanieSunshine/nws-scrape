package main

import (
  "os"
  "fmt"
  "github.com/cavaliercoder/grab"
)


func getImage(url string, path string){
  resp, err := grab.Get(path, url)
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
  }
  fmt.Println("Download saved to:", resp.Filename)
}
