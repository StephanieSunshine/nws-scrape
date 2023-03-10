package main

import (                                                                                                 
  // flag "github.com/spf13/pflag"                                                                       
  "fmt"                                                                                                  
  "github.com/PuerkitoBio/goquery"                                                                       
  "net/http"                                                                                             
  "os"                                                                                                   
  _ "github.com/davecgh/go-spew/spew"                                                                    
)

// get a list of satellites 
func getSatUrls(url string) []sat {
  results := []sat{}

  prefix := url[:len(url)-9]

  res, err := http.Get(url)
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
  }

  defer res.Body.Close()
  if res.StatusCode != 200 {
    fmt.Fprintf(os.Stderr, "status code error: %d %s", res.StatusCode, res.Status)
  }

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
  }

  // Find the review items
  doc.Find(".list-group-item").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the title
    title := s.Text()

    // and the link
    link, _ := s.Attr("href")
    
    results = append(results, sat{uint(i), title, prefix+link})
  })

  return results
}
