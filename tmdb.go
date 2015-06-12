package main

import(
      "encoding/json"
      "errors"
      "log"
      "strconv"
      )

const TMDB_KEY = "YOUR_TMDB_API_KEY_HERE"
const TMDB_URL = "https://api.themoviedb.org/3"

type Result struct {
  Id int `json:"id"`
}

type Search struct {
  Results []Result `json:"results"`
  TotalResults int `json:"total_results"`
}

type Movie struct {
  IMDB_ID string `json:"imdb_id"`
  Rating float64 `json:"vote_average"`
  Name string `json:"original_title"`
}

func tmdb(final_resp *Final) (string){
  movie_id, err := search_movie()

  if err != nil {
    log.Fatal("TMDB - Error Searching for movie: " + err.Error())
  }

  movie := get_movie(movie_id)

  if err != nil {
    log.Fatal("TMDB - Error Fetching the Movie: " + err.Error())
  }

  final_resp.Title = movie.Name
  final_resp.Ratings["user"]["tmdb"] = strconv.FormatFloat(movie.Rating, 'f', 2, 64)
  return movie.IMDB_ID
}

func search_movie() (int, error){
  var r Search
  url :=  TMDB_URL + "/search/movie?api_key=" + TMDB_KEY + "&query=fight+club"

  out := call_api(url)
  err := json.NewDecoder(out.Body).Decode(&r)

  if err != nil {
    log.Fatal("TMDB - Something went wrong while unmarshalling the data: " + err.Error())
  }

  if r.TotalResults < 1 {
    err := errors.New("ERROR: Movie Name not found")
    return 0, err
  } else {
    return r.Results[0].Id, nil
  }
}

func get_movie(id int) (Movie) {
  var m Movie

  url := TMDB_URL + "/movie/" + strconv.Itoa(id) + "?api_key=" + TMDB_KEY
  out := call_api(url)

  err := json.NewDecoder(out.Body).Decode(&m)

  if err != nil {
    log.Fatal("TMDB - Something went wrong while unmarshalling the data: " + err.Error())
  }

  return m

}
