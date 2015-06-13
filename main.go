package main


import (
        "log"
        "encoding/json"
        "net/http"

         "github.com/gorilla/mux"
        )

type Final struct {
  Title string
  Ratings map[string]map[string]string
}

func main() {
  router := mux.NewRouter()

  router.HandleFunc("/api/movies/{movieName}", GetRatings)

  log.Fatal(http.ListenAndServe(":8080", router))
}

func GetRatings(w http.ResponseWriter, r *http.Request) {
  Output := Final{
    Title: "",
    Ratings: map[string]map[string]string{
      "user": make(map[string]string),
      "critic": make(map[string]string),
    },
  }

  vars := mux.Vars(r)
  name := vars["movieName"]

  imdb_id := tmdb(name, &Output)

  imdb(imdb_id, &Output)
  metacritic(&Output)

  json.NewEncoder(w).Encode(&Output)
}
