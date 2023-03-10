package main

/*
 *    nws-scrape: National Weather Service GOES Image Scraper
 *
 *    Usage:  nws-scrape [OPTIONS]
 * 
 *    Options:
 *    -S, --sat [#]:            Show images available for a satellite and their
 *                              corresponding numbers. No number shows a list of
 *                              available satellites and their corresponding numbers.
 *
 *    -I, --img [#]:            Download image for a satellite to the current 
 *                              directory using the default filename. No number shows
 *                              a list of available images and their corresponding
 *                              numbers.
 *
 *    -O, --output <filename>:  Change the filename (and location) for the image being
 *                              saved. Note: this will not change the image format if
 *                              you change the file extension.
 *
 *    -L, --link-only:          Instead of downloading the image, just return the
 *                              images URL to stdout and exit.
 *
 *    -H, -h, ?, --help:        Show this help info
 *
 *    // Stephanie Sunshine // 2023 // Released under the MIT License //
 *    // https://github.com/StephanieSunshine/nws-scrape //
 */

import (
  "fmt"
)

// default flag default value work around
const MaxUint uint = uint((1 << 64) - 1)

// sat list
type sat struct {
  num uint
  loc string
  url string
}

// image list
type image struct {
  num uint
  description string
  url string
}

// cmdline input options
type Input struct {
  Satellite uint
  Image uint
  Output_file string
  Link_only bool
}

func main() {

  parent_url := "https://www.star.nesdis.noaa.gov/goes/index.php"

  options := Input{MaxUint, MaxUint, "", false}
  parseFlags(&options)

  switch {
    case options.Output_file != "":
      Fatal("Coming Soon")
    
    // default for options Satellite is maxUint as a flag that nothing was set
    case options.Satellite == MaxUint:
      listSatellites(parent_url)

      // default for options Image is maxUint as a flag that nothing was set
    case options.Image == MaxUint:
      sat := getSatUrls(parent_url)[options.Satellite]
      listImages(sat.url, sat.loc)

    case options.Link_only:
      sat := getSatUrls(parent_url)[options.Satellite]
      image := getImageUrls(sat.url)[options.Image]
      fmt.Println(image.url)

    default:
      sat := getSatUrls(parent_url)[options.Satellite]
      image := getImageUrls(sat.url)[options.Image]
      getImage(image.url)
  }
}
