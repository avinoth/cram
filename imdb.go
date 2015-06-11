package main

import(
      "encoding/json"
      "log"
      )

const IMDB_URL = "http://www.omdbapi.com"

type IMDB_MOVIE struct {
  ImdbRating string
  TomatoRating string
  TomatoMeter string
}

func imdb(id string, final_resp *Final) {
  var m IMDB_MOVIE

  url := IMDB_URL + "/?i=" + id + "&tomatoes=true"

  out := call_api(url)

  err := json.NewDecoder(out.Body).Decode(&m)

  if err != nil {
    log.Fatal("IMDB - Somthing went wrong while unmarshalling the data: " + err.Error())
  }

  final_resp.Ratings["user"]["imdb"] = m.ImdbRating
  final_resp.Ratings["user"]["rotten_tomatoes"] = m.TomatoRating
  final_resp.Ratings["critic"]["rotten_tomatoes_meter"] = m.TomatoMeter
}
