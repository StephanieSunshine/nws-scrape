package main

import (
  flag "github.com/spf13/pflag"
  "github.com/StephanieSunshine/columnize"
  "github.com/TwiN/go-color"
  "fmt"
  "os"
)

func newFlagUsage() {
    // f := flag.CommandLine
    output := []string{"Flags | Description", "|"}

    _, _ = fmt.Fprintf(os.Stderr, "\n    Usage: %s [Flags]\n\n", os.Args[0])
    flag.VisitAll(func(flag_ *flag.Flag) {
        if flag_.Usage == "" {
            return
        }
        _, usage := flag.UnquoteUsage(flag_)
        output = append(output, fmt.Sprintf("-%s, --%s | %s", flag_.Shorthand, flag_.Name, usage), "|") // Two spaces before -; see next two comments.
      })
    
    config := columnize.Config{
      Prefix: "  ",
      OutputWidth: columnize.AUTO,
      HeaderColors: &columnize.ColorList{color.InYellow, color.InPurple},
      BodyColors: &columnize.ColorList{color.InWhite, color.InCyan},
    }

    result := columnize.Format(output, &config)
    fmt.Fprintln(os.Stderr, result)
    fmt.Fprintln(os.Stderr)
}

