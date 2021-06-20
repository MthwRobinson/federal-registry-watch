package main

import (
  "fmt"
  "golang.org/x/net/html"
  "github.com/asmcos/requests"
  "strings"
)

func main (){
        url := "https://www.federalregister.gov/documents/current"
        // data := requests.Datas{
        //   "name":"requests_post_test",
        // }
        resp,_ := requests.Get(url)
        fmt.Println(resp.Text())

        doc, _ := html.Parse(strings.NewReader(resp.Text()))
        fmt.Println(doc)
}
