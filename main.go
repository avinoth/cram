package main


import (
        "fmt"
        "log"
        "encoding/json"
        )

type Final struct {
  Title string
  Ratings map[string]string
}

func main() {
  Output := Final{
    Title: "",
    Ratings: make(map[string]string),
  }

  imdb_id, err := tmdb(&Output)
  if err != nil {
    log.Fatal("Something wrong with TMDB.")
  }

  imdb(imdb_id, &Output)
  metacritic(&Output)

  json_out, _ := json.Marshal(&Output)
  fmt.Println(string(json_out))
}
