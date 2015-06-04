package main

import(
      "fmt"
      "net/http"
      "encoding/json"
      "errors"
      )

type Result struct {
  Id int `json:"id"`
}

type Search struct {
  Results []Result `json:"results"`
  TotalResults int `json:"total_results"`
}

func tmdb(){
  base_url := "https://api.themoviedb.org/3/"
  search_movie(base_url)
}

func search_movie(base_url string) {

  var r Search

  search_url := base_url + "search/movie?api_key=" + TMDB_KEY + "&query=fight"

  out, err := call_api(r, search_url)

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(out.Id)
  }
}


func call_api(r Search, url string) (out Result, err error) {
  var emp_struct Result
  client := &http.Client{}

  req, err := http.NewRequest("GET", url, nil)

  if err != nil {
    fmt.Println("Something wrong with the URL" + url)
    return emp_struct, err
  }
  req.Header.Set("Accept", "application/json")

  resp, err := client.Do(req)
  if err != nil {
    fmt.Println("Something wrong while fetching the data")
    return emp_struct, err
  }

  defer resp.Body.Close()

  err = json.NewDecoder(resp.Body).Decode(&r)

  if err != nil {
    fmt.Println("Somthing went wrong while unmarshalling the data")
    return emp_struct, err
  }

  if r.TotalResults < 1 {
    err := errors.New("ERROR: Movie Name not found")
    return emp_struct, err
  } else {
    return r.Results[0], nil
  }
}
