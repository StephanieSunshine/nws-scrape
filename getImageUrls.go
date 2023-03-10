package main

import (                                                                                                 
  // flag "github.com/spf13/pflag"                                                                       
  "fmt"                                                                                                  
  "github.com/PuerkitoBio/goquery"                                                                       
  "net/http"                                                                                             
  "os"                                                                                                   
  _ "github.com/davecgh/go-spew/spew"                                                                    
)

// get a list of satellites ( title, images )
func getImageUrls(url string) []image {
  results := []image{}
  // prefix := url[:len(url)-9]

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
  doc.Find("a").Each(func(i int, s *goquery.Selection) {
    // resolution
    res := s.Text()
    
    // For each item found, get the title
    title, _ := s.Attr("title")
    
    // and the link
    link, _ := s.Attr("href")

    if len(link) > 3 && link[len(link)-3:] == "jpg" { 
      results = append(results, image{uint(i), "["+res+"] "+ title, link})
    }
  })

  return results
}
