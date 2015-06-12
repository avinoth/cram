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

  imdb_id := tmdb(&Output)

  imdb(imdb_id, &Output)
  metacritic(&Output)

  json_out, _ := json.Marshal(&Output)
  fmt.Println(string(json_out))
}
