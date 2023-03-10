nws-scrape
==========

National Weather Service GOES Image Scraper

# Usage nws-scrape [OPTIONS]

Usage
=====

```
  Flags            Description
                                                                                                         
  -I, --img        Download image for a satelite to the current directory using the default filename. No 
                   flag shows a list of available images and their corresponding numbers.              
                                                                                                         
  -L, --link-only  Instead of downloading the image, just return the images URL to stdout and exit.      
                                                                                                         
  -O, --output     Change the filename (and location) for the image being saved. Note: this will not     
                   change the image format if you change the file extension.                             
                                                                                                         
  -S, --sat        Show images available for a satellite and their corresponding numbers. No flag shows
                   a list of available satellites and their corresponding numbers.
```
