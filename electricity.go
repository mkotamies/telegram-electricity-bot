package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
  resp, err := http.Get("https://api.porssisahko.net/v1/latest-prices.json")
  if(err != nil) {
    fmt.Println(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if(err != nil) {
    fmt.Println(err)
  }
  prices := string(body)
  fmt.Println(prices)
}
