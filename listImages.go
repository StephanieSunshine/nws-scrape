package main

import (
  "fmt"
  "os"
  "github.com/StephanieSunshine/columnize"
  "github.com/TwiN/go-color"
)

func listImages(url string, loc string) {
  images := getImageUrls(url)
  output := []string{"Image | Description"}

  for i, v := range images {
    output = append(output, fmt.Sprintf("%d | %s ", i, v.description))
  }

  config := columnize.Config{ 
    Prefix: "  ",
    OutputWidth: columnize.AUTO,
    HeaderColors: &columnize.ColorList{color.InYellow, color.InPurple},
    BodyColors: &columnize.ColorList{color.InWhite, color.InCyan},
  }

  fmt.Fprintln(os.Stderr, "\nNo image selected.  Pick One For: "+loc+"\n")
  fmt.Fprintln(os.Stderr, columnize.Format(output, &config)) 
  fmt.Fprintln(os.Stderr)
}
