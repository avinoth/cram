package main


import (
        "fmt"
        "log"
        "encoding/json"
        )

type Fin_JSON struct {
  Title string
  Ratings map[string]float64
}

func main() {
  Output := Fin_JSON{
    Title: "",
    Ratings: make(map[string]float64),
  }

  imdb_id, err := tmdb(&Output)
  if err != nil {
    log.Fatal("Something wrong with TMDB.")
  }

  imdb(imdb_id, &Output)

  json_out, _ := json.Marshal(&Output)
  fmt.Println(imdb_id)
  fmt.Println(string(json_out))
}
