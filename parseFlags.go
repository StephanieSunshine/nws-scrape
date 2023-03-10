package main

import (
  flag "github.com/spf13/pflag"
)

func parseFlags(options *Input) {
  flag.Usage = newFlagUsage
  flag.UintVarP(&options.Satellite, "sat", "S", MaxUint, "Show images available for a satellite and their corresponding numbers. No flag shows a list of available satellites and their corresponding numbers.")
  flag.UintVarP(&options.Image, "img", "I", MaxUint, "Download image for a satelite to the current directory using the default filename. No flag shows a list of available images and their corresponding numbers.")
  flag.StringVarP(&options.Output_file, "output", "O", "", "Change the filename (and location) for the image being saved. Note: this will not change the image format if you change the file extension.")
  flag.BoolVarP(&options.Link_only, "link-only", "L", false, "Instead of downloading the image, just return the images URL to stdout and exit.")                                                                                                       
  flag.Parse()
}
