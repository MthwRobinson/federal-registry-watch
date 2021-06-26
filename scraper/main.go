package main

import (
	"fmt"
	"net/http"
	"os"
  "io/ioutil"
)


func fetchUrl(url string) string {
  // Makes an http call to the specified url and returns a string representation of the
  // HTML response.
  resp, _ := http.Get(url)
  bytes, _ := ioutil.ReadAll(resp.Body)
  return string(bytes)
}


func main() {
  var url, html string
  url = os.Args[1]
  html = fetchUrl(url)
  fmt.Println(html)
}
