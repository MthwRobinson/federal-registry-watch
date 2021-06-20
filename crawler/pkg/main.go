package main

import "github.com/asmcos/requests"


func main (){
        url := "https://www.federalregister.gov/documents/current"
        // data := requests.Datas{
        //   "name":"requests_post_test",
        // }
        resp,_ := requests.Get(url)
        println(resp.Text())
}
