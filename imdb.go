package main

import(
      "encoding/json"
      "log"
      )

const IMDB_URL = "http://www.omdbapi.com"

type IMDB_MOVIE struct {
  Rating float64 `json:imdbRating`
}

func imdb(id string, fin_out *Fin_JSON) {
  var m IMDB_MOVIE

  url := IMDB_URL + "/?i=" + id
  out := call_api(url)

  err := json.NewDecoder(out.Body).Decode(&m)

  if err != nil {
    log.Fatal("Somthing went wrong while unmarshalling the data - IMDB")
  }
  fin_out.Ratings["imdb"] = m.Rating
}
