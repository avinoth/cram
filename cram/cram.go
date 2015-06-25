package cram


import (
        "encoding/json"
        "net/http"
        "fmt"
        )

type Final struct {
  Title string
  Ratings map[string]map[string]string
}

func init() {
  http.HandleFunc("/", GetRatings)
}

func GetRatings(w http.ResponseWriter, r *http.Request) {
  Output := Final{
    Title: "",
    Ratings: map[string]map[string]string{
      "user": make(map[string]string),
      "critic": make(map[string]string),
    },
  }

  name := r.URL.Path[1:]

  imdb_id := tmdb(name, &Output)

  imdb(imdb_id, &Output)
  metacritic(&Output)

  w.Header().Set("Application-Type", "Application/json")
  res, err := json.Marshal(Output)
  fmt.Println(res)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  fmt.Fprint(w, string(res))
}
