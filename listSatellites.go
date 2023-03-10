package main

import (
  "github.com/StephanieSunshine/columnize"
  "github.com/TwiN/go-color"
  "fmt"
  "os"
)

func listSatellites(url string) {
  output := []string{"Sat | Location"}

  for _, v := range getSatUrls(url) {
    output = append(output, fmt.Sprintf("%d | %s", v.num, v.loc))
  }

  config := columnize.Config{
    Prefix: "  ",
    OutputWidth: columnize.AUTO,
    HeaderColors: &columnize.ColorList{color.InYellow, color.InPurple},
    BodyColors: &columnize.ColorList{color.InWhite, color.InCyan},
  }

  fmt.Fprintln(os.Stderr, "\nNo satellite selected.  Pick One:\n")
  fmt.Fprintln(os.Stderr, columnize.Format(output, &config))
  fmt.Fprintln(os.Stderr) 
}
