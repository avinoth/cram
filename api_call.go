package main

import (
        "net/http"
        "log"
        )

func call_api(url string) *http.Response {
  client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)

  if err != nil {
    log.Fatal("Something wrong with the URL: " + url + " - " + err.Error())
  }

  req.Header.Set("Accept", "application/json")

  resp, err := client.Do(req)

  if err != nil {
    log.Fatal("Something went wrong while fetching the data: " + err.Error())
  }

  return resp
}
