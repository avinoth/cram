package main


import (
        "fmt"
        "log"
        "encoding/json"
        )

type Final struct {
  Title string
  Ratings map[string]map[string]string
}

func main() {
  Output := Final{
    Title: "",
    Ratings: map[string]map[string]string{
      "user": make(map[string]string),
      "critic": make(map[string]string),
    },
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
