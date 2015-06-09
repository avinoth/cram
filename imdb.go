package main

import(
      "encoding/json"
      "log"
      )

const IMDB_URL = "http://www.omdbapi.com"

type IMDB_MOVIE struct {
  TomatoRating string
  ImdbRating string
  TomatoMeter string
}

func imdb(id string, fin_out *Fin_JSON) {
  var m IMDB_MOVIE

  url := IMDB_URL + "/?i=" + id + "&tomatoes=true"

  out := call_api(url)

  err := json.NewDecoder(out.Body).Decode(&m)

  if err != nil {
    log.Fatal("Somthing went wrong while unmarshalling the data - IMDB")
  }

  fin_out.Ratings["rotten_tomatoes_rating"] = m.TomatoRating
  fin_out.Ratings["imdb"] = m.ImdbRating
  fin_out.Ratings["rotten_tomatoes_meter"] = m.TomatoMeter
}
